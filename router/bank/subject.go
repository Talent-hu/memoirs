package bank

import (
	"github.com/gin-gonic/gin"
	"memoirs/api"
	"memoirs/middleware"
)

type SubjectRouter struct{}

func (this *SubjectRouter) InitSubjectRouter(Router *gin.RouterGroup) {
	privateRoute := Router.Group("subject").Use(middleware.GinJwt())
	subjectApi := api.ApiGroupApp.SubjectApi
	{
		privateRoute.GET("/list", subjectApi.QueryAll)
		privateRoute.GET("/queryById", subjectApi.QueryById)
		privateRoute.POST("/add", subjectApi.Insert)
		privateRoute.POST("/update", subjectApi.Update)
		privateRoute.POST("/delete", subjectApi.Delete)
	}

}
