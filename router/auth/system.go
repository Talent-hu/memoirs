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
		privateRouter.GET("/queryDict", systemApi.QueryDictAll)
		privateRouter.POST("/addDict", systemApi.AddDict)
		privateRouter.POST("/updateDict", systemApi.UpdateDict)
		privateRouter.POST("/deleteDict", systemApi.DeleteDict)

		privateRouter.GET("/queryDictItem", systemApi.QueryDictItem)
		privateRouter.POST("/addDictItem", systemApi.AddDictItem)
		privateRouter.POST("/updateDictItem", systemApi.UpdateDictItem)
		privateRouter.POST("/deleteDictItem", systemApi.DeleteDictItem)
	}
}
