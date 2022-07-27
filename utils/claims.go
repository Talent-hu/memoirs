package utils

import (
	"github.com/gin-gonic/gin"
	"memoirs/common/constant"
	"memoirs/common/jwts"
	"memoirs/global"
)

func GetClaims(ctx *gin.Context) (*jwts.UserStdClaims,error) {
	token := ctx.Request.Header.Get(constant.TOKEN_NAME)
	jwt := jwts.NewJWT()
	claims, err := jwt.ParseToken(token)
	if err != nil {
		global.Log.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在Authorization且claims是否为规定结构")
	}
	return claims,err
}


// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(ctx *gin.Context) uint {
	if claims, exists := ctx.Get("claims"); !exists {
		if cl, err := GetClaims(ctx); err != nil {
			return 0
		} else {
			return cl.UserInfo.UserId
		}
	} else {
		waitUse := claims.(*jwts.UserStdClaims)
		return waitUse.UserInfo.UserId
	}
}