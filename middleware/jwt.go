package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"memoirs/common/constant"
	"memoirs/common/jwts"
	"memoirs/common/response"
	"memoirs/global"
	"net/http"
	"time"
)

func GinJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get(constant.TOKEN_NAME)
		if token == "" {
			global.Log.Error("权限不足")
			response.FailWithMessage(ctx, "未登录或非法访问")
			ctx.Abort()
			return
		}
		jwt := jwts.NewJWT()
		// 解析token信息
		claims, err := jwt.ParseToken(token)
		if err != nil {
			if err == jwts.TokenExpired {
				response.FailWithMessage(ctx, "授权已过期")
				ctx.Abort()
				return
			}
			response.FailWithMessage(ctx, err.Error())
			ctx.Abort()
			return
		}
		val := global.Redis.Get(context.Background(), claims.UserInfo.Identity).Val()
		if val != token {
			response.FailWithDetail(ctx, http.StatusUnauthorized, "授权已过期")
			ctx.Abort()
			return
		}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + global.Config.ExpireTime
			newToken, _ := jwt.CreateTokenByOldToken(token, *claims)
			ctx.Header("new-Token", newToken)

			// redis记录token
			global.Redis.Set(context.Background(), claims.UserInfo.Identity, newToken, time.Second*time.Duration(global.Config.ExpireTime))
		}
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
