package middleware

import (
	"github.com/gin-gonic/gin"
	"memoirs/pkg/response"
	"memoirs/service"
	"memoirs/utils"
)

var casbinService = service.ServiceGroupApp.CasbinService

func CasbinHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userClaims, _ := utils.GetClaims(ctx)
		// 获取请求的path
		obj := ctx.Request.URL.Path
		// 获取请求的方法
		act := ctx.Request.Method
		// 获取用户角色
		roleCodes := userClaims.UserInfo.RoleCodes
		e := casbinService.Casbin()
		flag := false
		for _, sub := range roleCodes {
			flag, _ = e.Enforce(sub, obj, act)
			if flag {
				break
			}
		}
		if flag {
			ctx.Next()
		} else {
			response.FailWithMessage(ctx, "权限不足")
			ctx.Abort()
			return
		}
	}
}
