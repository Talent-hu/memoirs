package auth

import (
	"github.com/gin-gonic/gin"
	"memoirs/api"
	"memoirs/middleware"
)

type UserRouter struct{}

func (this *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	privateRouter := Router.Group("user").Use(middleware.GinJwt())
	userApi := api.ApiGroupApp.UserApi
	{
		userRouter.POST("/login", userApi.Login)
		userRouter.GET("/publicKey", userApi.PublicKey)
	}
	{
		privateRouter.GET("/info", userApi.QueryUserInfo)
		privateRouter.GET("/listAll", userApi.QueryUserList)
	}
}
