package auth

import (
	"github.com/gin-gonic/gin"
	"memoirs/api"
	"memoirs/middleware"
)

type CasbinRouter struct{}

func (router *CasbinRouter) InitCasbinRouter(Router *gin.RouterGroup) {
	privateRouter := Router.Group("casbin").Use(middleware.GinJwt())
	casbinApi := api.ApiGroupApp.CasbinApi
	{
		privateRouter.POST("update", casbinApi.UpdateCasbin)
		privateRouter.GET("query", casbinApi.GetPolicyPathByRoleId)
	}
}
