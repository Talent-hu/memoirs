package auth

import (
	"context"
	"encoding/base64"
	"errors"
	"go.uber.org/zap"
	"memoirs/global"
	"memoirs/model/auth"
	"memoirs/model/vo"
	"memoirs/pkg/jwts"
	"memoirs/utils"
	"strings"
	"time"
)

type AuthService struct{}

func (auth *AuthService) Login(loginReq vo.LoginRequest) (any, error) {
	// 如果账号长度大于32，说明是加密账号，走加密逻辑
	if len(loginReq.Account) > 32 {
		// RSA 私钥解密
		account, err := base64.StdEncoding.DecodeString(loginReq.Account)
		password, err := base64.StdEncoding.DecodeString(loginReq.Password)
		if err != nil {
			global.Log.Error("登录信息解码失败")
			return nil, errors.New("登录信息解码失败")
		}
		privateKey := global.Redis.Get(context.Background(), "rsa_private").Val()
		account, _ = utils.RsaDecrypt(account, []byte(privateKey))
		password, err = utils.RsaDecrypt(password, []byte(privateKey))
		if err != nil {
			global.Log.Error("登录信息解码失败", zap.Error(err))
			return nil, errors.New("登录信息解码失败")
		}
		loginReq.Account = string(account)
		loginReq.Password = string(password)
	}
	encPwd := utils.GenerateMD5(loginReq.Password)
	user, err := userMapper.Login(loginReq.Account)
	if err != nil {
		global.Log.Error("数据库异常", zap.Error(err))
		return nil, errors.New("数据库异常")
	}
	if user.Password != encPwd {
		global.Log.Error("账号或密码错误")
		return nil, errors.New("账号或密码错误")
	}
	resp, err := auth.NextToken(user)
	return resp, err
}

func (auth *AuthService) NextToken(user *auth.User) (any, error) {
	// 查询用户的角色列表
	userInfo, _ := userMapper.QueryUserInfo(user.ID)
	roles := userInfo.Roles
	var roleCodes []string
	for _, item := range roles {
		roleCodes = append(roleCodes, item.RoleCode)
	}
	jwt := jwts.NewJWT()
	usrClaims := &jwts.UserClaims{
		UserId:    user.ID,
		UserName:  user.Username,
		NickName:  user.NickName,
		Identity:  user.Identity,
		RoleCodes: roleCodes,
	}
	claims := jwt.CreateClaims(usrClaims)
	token, err := jwt.CreateToken(claims)
	if err != nil {
		global.Log.Error("生成token失败！", zap.Error(err))
		return nil, errors.New("获取token失败")
	}
	// 将token存入redis缓存中
	global.Redis.Set(context.Background(), user.Identity, token, time.Second*time.Duration(int64(jwts.ExpireTime)))
	resp := new(vo.LoginReply)
	resp.Token = token
	return resp, nil
}

func (auth *AuthService) RsaPublicSecret() (any, error) {
	publicKey := global.Redis.Get(context.Background(), "rsa_public").Val()
	if publicKey == "" {
		privKey, pubKey, err := utils.GeneratorRSAKey()
		publicKey = string(pubKey)
		privateKey := string(privKey)
		if err != nil {
			global.Log.Error("生成RSA密钥对失败", zap.Error(err))
			return nil, errors.New("生成RSA密钥对失败")
		}
		global.Redis.Set(context.Background(), "rsa_public", publicKey, time.Hour*24)
		global.Redis.Set(context.Background(), "rsa_private", privateKey, time.Hour*24)

	}
	split := strings.Split(publicKey, "\n")
	publicKey = strings.Join(split[1:len(split)-2], "")
	resp := new(vo.RsaKeyReply)
	resp.PublicKey = publicKey
	return resp, nil
}
