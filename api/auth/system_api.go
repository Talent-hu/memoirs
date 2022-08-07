package auth

import (
	"github.com/gin-gonic/gin"
	"memoirs/model/vo"
	"memoirs/pkg/response"
	"memoirs/validate"
)

type SystemApi struct{}

func (api *SystemApi) QueryDictAll(ctx *gin.Context) {
	var queryPage vo.ListQuery
	_ = ctx.ShouldBindJSON(&queryPage)
	if err := validate.Verify(queryPage, validate.QueryPageListVerify); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	list, total, err := systemService.QueryDict(queryPage)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	resp := new(vo.PageQueryReply)
	resp.List = list
	resp.Total = total
	response.OkWithData(ctx, resp)
}

func (api *SystemApi) AddDict(ctx *gin.Context) {
	var dict vo.Dict
	_ = ctx.ShouldBindJSON(&dict)
	err := systemService.AddDict(dict)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}
