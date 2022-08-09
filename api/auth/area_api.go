package auth

import (
	"github.com/gin-gonic/gin"
	"memoirs/pkg/response"
)

type AreaApi struct{}

func (api *AreaApi) QueryList(ctx *gin.Context) {
	resp, err := areaService.QueryList()
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, resp)
}
