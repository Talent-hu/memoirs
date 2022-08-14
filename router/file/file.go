package file

import (
	"github.com/gin-gonic/gin"
	"memoirs/api"
	"memoirs/middleware"
)

type FileRouter struct{}

func (router *FileRouter) InitFileRouter(Router *gin.RouterGroup) {
	privateRouter := Router.Group("file").Use(middleware.GinJwt())
	fileApi := api.ApiGroupApp.ImageApi
	{
		privateRouter.POST("/uploadImage", fileApi.ImageUpload)
	}
}
