package module

import (
	"github.com/gin-gonic/gin"
	"memoirs/api"
	"memoirs/middleware"
)

type MenuRouter struct{}

func (this *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	//userRouter := Router.Group("menu")
	privateRouter := Router.Group("menu").Use(middleware.GinJwt())
	menuApi := api.ApiGroupApp.MenuApi
	{
		privateRouter.POST("/add", menuApi.AddMenu)
		privateRouter.GET("/tree", menuApi.QueryMenuTree)
		privateRouter.GET("/list", menuApi.QueryMenuList)
		privateRouter.POST("/remove", menuApi.RemoveMenu)
	}
}
