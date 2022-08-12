package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"memoirs/middleware"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(middleware.GinLogger(),
		middleware.GinRecovery(true),
		middleware.Cors(),
		middleware.RecordOptions())
	publicGroup := router.Group("")
	{
		RouterGroupApp.UserRouter.InitUserRouter(publicGroup)
		RouterGroupApp.MenuRouter.InitMenuRouter(publicGroup)
		RouterGroupApp.RoleRouter.InitRoleRouter(publicGroup)
		RouterGroupApp.SystemRouter.InitSystemRouter(publicGroup)
		RouterGroupApp.AreaRouter.InitAreaRouter(publicGroup)
		RouterGroupApp.CasbinRouter.InitCasbinRouter(publicGroup)
	}
	{
		RouterGroupApp.SubjectRouter.InitSubjectRouter(publicGroup)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
