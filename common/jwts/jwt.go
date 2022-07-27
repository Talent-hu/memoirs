package jwts

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"memoirs/global"
	"time"
)

type JWT struct {
	SignKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		SignKey: []byte(global.Config.SignKey),
	}
}

func (this *JWT) CreateClaims(uClaim *UserClaims) UserStdClaims {
	stdClaims := jwt.StandardClaims{
		NotBefore: time.Now().Unix() - 1000,                     // 签名生效时间
		ExpiresAt: time.Now().Unix() + global.Config.ExpireTime, // 签名过期时间
		IssuedAt:  time.Now().Unix(),
		Issuer:    global.Config.AppIss, // 签名的发行者
	}
	return UserStdClaims{
		StandardClaims: stdClaims,
		BufferTime:     global.Config.BufferTime,
		UserInfo:       uClaim,
	}
}

func (this *JWT) CreateToken(claims UserStdClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(this.SignKey)
	return tokenStr, err
}

// 旧token 换新token 使用归并回源避免并发问题
func (this *JWT) CreateTokenByOldToken(oldToken string, claims UserStdClaims) (string, error) {
	v, err, _ := global.Concurrent_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
		return this.CreateToken(claims)
	})
	return v.(string), err
}

func (this *JWT) ParseToken(tokenStr string) (*UserStdClaims, error) {
	if tokenStr == "" {
		return nil, errors.New("token is empty")
	}
	token, err := jwt.ParseWithClaims(tokenStr, &UserStdClaims{}, func(token *jwt.Token) (interface{}, error) {
		return this.SignKey,nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*UserStdClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, TokenInvalid
}

