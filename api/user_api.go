package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"memoirs/common/jwts"
	"memoirs/common/response"
	"memoirs/global"
	"memoirs/model"
	"memoirs/model/vo"
	"memoirs/utils"
	"memoirs/validate"
	"strings"
	"time"
)

type UserApi struct{}

// @Tag UserApi
// @Summary 用户登录
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param data body vo.LoginRequest false "查询参数"
// @Success 200 {object} response.Response{data=vo.LoginResponse,message=string} "返回用户token"
// @Router /user/login [post]
func (this *UserApi) Login(ctx *gin.Context) {
	var loginReq vo.LoginRequest
	_ = ctx.ShouldBindJSON(&loginReq)
	if err := validate.Verify(loginReq, validate.LoginVerify); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	if len(loginReq.Account) > 32 {
		// RSA 私钥解密
		privateKey := global.Redis.Get(context.Background(), "rsa_private").Val()
		account, _ := utils.RsaDecrypt([]byte(loginReq.Account), []byte(privateKey))
		password, err := utils.RsaDecrypt([]byte(loginReq.Password), []byte(privateKey))
		if err != nil {
			response.FailWithMessage(ctx, "用户名或密码错误")
			return
		}
		loginReq.Account = string(account)
		loginReq.Password = string(password)
	}
	encPwd := utils.GenerateMD5(loginReq.Password)
	user, err := userService.Login(loginReq.Account)
	if err != nil {
		response.FailWithMessage(ctx, "用户名或密码错误")
		return
	}
	if user.Password != encPwd {
		response.FailWithMessage(ctx, "用户名或密码错误")
		return
	}
	this.TokenNext(ctx, user)
}

func (this *UserApi) TokenNext(ctx *gin.Context, user *model.User) {
	jwt := jwts.NewJWT()
	usrClaims := &jwts.UserClaims{
		UserId:   user.ID,
		UserName: user.Username,
		NickName: user.NickName,
		Identity: user.Identity,
	}
	claims := jwt.CreateClaims(usrClaims)
	token, err := jwt.CreateToken(claims)
	if err != nil {
		global.Log.Error("生成token失败！", zap.Error(err))
		response.FailWithMessage(ctx, "获取token失败！")
		return
	}
	// 将token存入redis缓存中
	global.Redis.Set(context.Background(), user.Identity, token, time.Second*time.Duration(global.Config.ExpireTime))
	resp := new(vo.LoginResponse)
	resp.Token = token
	response.OkWithDetail(ctx, "登录成功", resp)
}

// @Tag UserApi
// @Summary 获取RSA公钥
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {object} response.Response{data=vo.GetRsaKeyResponse,message=string} "返回用户token"
// @Router /user/publicKey [get]
func (this *UserApi) PublicKey(ctx *gin.Context) {
	resp := new(vo.GetRsaKeyResponse)
	publicKey := global.Redis.Get(context.Background(), "rsa_public").Val()
	if publicKey == "" {
		privateKey, publicKey, err := utils.GeneratorRSAKey()
		if err != nil {
			response.FailWithMessage(ctx, "生成RSA密钥对失败")
			return
		}
		global.Redis.Set(context.Background(), "rsa_public", string(publicKey), time.Hour*24)
		global.Redis.Set(context.Background(), "rsa_private", string(privateKey), time.Hour*24)

	}
	fmt.Println(publicKey)
	split := strings.Split(publicKey, "\n")
	publicKey = strings.Join(split[1:len(split)-2], "")
	resp.PublicKey = publicKey
	response.OkWithData(ctx, resp)
}

// @Tag UserApi
// @Summary 获取用户信息
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {object} response.Response{data=vo.GetUserInfoResponse,message=string} "返回用户token"
// @Router /user/info [get]
func (this *UserApi) GetUserInfo(ctx *gin.Context) {
	userId := utils.GetUserID(ctx)
	userInfo, err := userService.GetUserInfo(userId)
	if err != nil {
		response.FailWithMessage(ctx, "系统错误，未查询到对应数据。")
		return
	}
	if len(userInfo.Roles) <= 0 {
		response.FailWithMessage(ctx, "用户还没有赋予权限，请联系管理员授权")
		return
	}
	resp := new(vo.UserInfoResponse)
	_ = utils.CopyProperties(userInfo, resp)
	for _, role := range userInfo.Roles {
		roleModel := new(vo.RoleModel)
		roleModel.RoleName = role.RoleName
		roleModel.RoleCode = role.RoleCode
		resp.Roles = append(resp.Roles, *roleModel)
	}
	response.OkWithData(ctx, resp)
}

func (this *UserApi) QueryUserList(ctx *gin.Context) {
	var pageInfo vo.BasePage
	_ = ctx.ShouldBindJSON(&pageInfo)
	list, err := userService.QueryUserList(pageInfo.PageSize, pageInfo.Offset())
	if err != nil {
		response.FailWithMessage(ctx, "系统错误，未查询到对应数据。")
		return
	}
	var userInfoList []vo.UserInfoResponse
	for _, user := range list {
		var userInfo vo.UserInfoResponse
		_ = utils.CopyProperties(user, &userInfo)
		for _, role := range user.Roles {
			roleModel := new(vo.RoleModel)
			roleModel.RoleName = role.RoleName
			roleModel.RoleCode = role.RoleCode
			userInfo.Roles = append(userInfo.Roles, *roleModel)
		}
		userInfoList = append(userInfoList, userInfo)
	}
	response.OkWithData(ctx, userInfoList)
}
