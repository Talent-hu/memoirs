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
		privateRouter.GET("/queryRoleAll", roleApi.QueryRoleAllList)
		privateRouter.GET("/queryUserRole", roleApi.QueryRoleAllList)
		privateRouter.POST("/addRoleMenu", roleApi.AddRoleAndMenu)
		privateRouter.POST("/addRole", roleApi.AddRole)
		privateRouter.POST("/update", roleApi.UpdateRole)
		privateRouter.POST("/delete", roleApi.DeleteRole)
	}
}
