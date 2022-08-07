package auth

import (
	"github.com/gin-gonic/gin"
	"memoirs/api"
	"memoirs/middleware"
)

type SystemRouter struct{}

func (router *SystemRouter) InitSystemRouter(Router *gin.RouterGroup) {
	privateRouter := Router.Group("system").Use(middleware.GinJwt())
	systemApi := api.ApiGroupApp.SystemApi
	{
		privateRouter.POST("/dict/list", systemApi.QueryDictAll)
		privateRouter.POST("/dict/add", systemApi.AddDict)
		privateRouter.POST("/dict/update", systemApi.UpdateDict)
		privateRouter.POST("/dict/delete", systemApi.DeleteDict)

		privateRouter.POST("/item/list", systemApi.QueryDictItem)
		privateRouter.POST("/item/add", systemApi.AddDictItem)
		privateRouter.POST("/item/update", systemApi.UpdateDictItem)
		privateRouter.POST("/item/delete", systemApi.DeleteDictItem)
	}
}
