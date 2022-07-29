package module

import (
	"github.com/gin-gonic/gin"
	"memoirs/api"
	"memoirs/middleware"
)

type RoleRouter struct{}

func (this *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	privateRouter := Router.Group("role").Use(middleware.GinJwt())
	roleApi := api.ApiGroupApp.RoleApi
	{
		privateRouter.POST("/add", roleApi.AddRoleAndMenu)
	}
}
