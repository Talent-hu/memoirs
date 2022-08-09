package auth

import (
	"github.com/gin-gonic/gin"
	"memoirs/api"
)

type AreaRouter struct{}

func (router *AreaRouter) InitAreaRouter(Router *gin.RouterGroup) {
	publicRouter := Router.Group("/area")
	areaApi := api.ApiGroupApp.AreaApi
	{
		publicRouter.GET("/list", areaApi.QueryList)
	}
}
