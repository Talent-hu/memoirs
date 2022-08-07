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
	}
	{
		RouterGroupApp.SubjectRouter.InitSubjectRouter(publicGroup)
	}
	url := ginSwagger.URL("http://127.0.0.1:8888/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
