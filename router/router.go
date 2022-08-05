package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"memoirs/middleware"
	"memoirs/router/module"
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
		module.RouterGroupApp.UserRouter.InitUserRouter(publicGroup)
		module.RouterGroupApp.MenuRouter.InitMenuRouter(publicGroup)
		module.RouterGroupApp.RoleRouter.InitRoleRouter(publicGroup)
		module.RouterGroupApp.SystemRouter.InitSystemRouter(publicGroup)
	}
	url := ginSwagger.URL("http://127.0.0.1:8888/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
