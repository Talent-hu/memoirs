package auth

import (
	"github.com/gin-gonic/gin"
	"memoirs/model/vo"
	"memoirs/pkg/response"
	"memoirs/validate"
	"strconv"
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

func (api *SystemApi) UpdateDict(ctx *gin.Context) {
	var dict vo.Dict
	_ = ctx.ShouldBindJSON(&dict)
	err := systemService.UpdateDict(dict)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

func (api *SystemApi) DeleteDict(ctx *gin.Context) {
	dictIdStr := ctx.Request.URL.Query().Get("id")
	if dictIdStr == "" {
		response.FailWithMessage(ctx, "id不能为空")
		return
	}
	dictId, _ := strconv.Atoi(dictIdStr)
	err := systemService.DeleteDict(uint(dictId))
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

func (api *SystemApi) QueryDictItem(ctx *gin.Context) {
	var pageItem vo.PageDictItem
	_ = ctx.ShouldBindJSON(&pageItem)
	list, total, err := systemService.QueryDictItemList(pageItem)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	resp := new(vo.PageQueryReply)
	resp.Total = total
	resp.List = list
	response.OkWithData(ctx, resp)
}

func (api *SystemApi) AddDictItem(ctx *gin.Context) {
	var item vo.DictItem
	_ = ctx.ShouldBindJSON(&item)
	err := systemService.InsertDictItem(item)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

func (api *SystemApi) UpdateDictItem(ctx *gin.Context) {
	var item vo.DictItem
	_ = ctx.ShouldBindJSON(&item)
	err := systemService.UpdateDictItem(item)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

func (api *SystemApi) DeleteDictItem(ctx *gin.Context) {
	itemId := ctx.Request.URL.Query().Get("id")
	if itemId == "" {
		response.FailWithMessage(ctx, "id不能为空")
		return
	}
	dictItemId, _ := strconv.Atoi(itemId)
	err := systemService.DeleteDictItem(uint(dictItemId))
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}
