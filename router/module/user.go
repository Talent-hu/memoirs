package module

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
		userRouter.GET("/listAll", userApi.QueryUserList)
	}
	{
		privateRouter.GET("/info", userApi.GetUserInfo)
	}
}
