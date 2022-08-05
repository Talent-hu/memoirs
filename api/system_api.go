package api

import (
	"github.com/gin-gonic/gin"
	"memoirs/pkg/response"
)

type SystemApi struct{}

func (sys *SystemApi) QuerySetting(ctx *gin.Context) {
	resp, err := systemService.QuerySetting()
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, resp)
}
