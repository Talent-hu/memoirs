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
		privateRouter.GET("/listAll", roleApi.QueryRoleAllList)
		privateRouter.GET("/listUserRole", roleApi.QueryRoleAllList)
		privateRouter.POST("/addRoleMenu", roleApi.AddRoleAndMenu)
		privateRouter.POST("/add", roleApi.AddRole)
		privateRouter.POST("/update", roleApi.UpdateRole)
		privateRouter.POST("/delete", roleApi.DeleteRole)
	}
}
