package auth

import (
	"github.com/gin-gonic/gin"
	"memoirs/api"
	"memoirs/middleware"
)

type SystemRouter struct{}

func (sys *SystemRouter) InitSystemRouter(Router *gin.RouterGroup) {
	systemRouter := Router.Group("system").Use(middleware.GinJwt())
	systemApi := api.ApiGroupApp.SystemApi
	{
		systemRouter.GET("/setting", systemApi.QuerySetting)
	}

}
