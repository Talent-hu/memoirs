package bank

import (
	"github.com/gin-gonic/gin"
	"memoirs/model/vo"
	"memoirs/pkg/response"
	"memoirs/validate"
	"strconv"
)

type SubjectApi struct{}

func (this *SubjectApi) QueryAll(ctx *gin.Context) {
	resp := subjectService.QueryAll()
	response.OkWithData(ctx, resp)
}

func (this *SubjectApi) QueryById(ctx *gin.Context) {
	subIdStr := ctx.Request.URL.Query().Get("id")
	if subIdStr == "" {
		response.FailWithMessage(ctx, "id不能为空")
		return
	}
	subId, err := strconv.Atoi(subIdStr)
	if err != nil {
		response.FailWithMessage(ctx, "id必须是数字类型")
		return
	}
	resp := subjectService.QueryById(uint(subId))
	response.OkWithData(ctx, resp)
}

func (this *SubjectApi) Insert(ctx *gin.Context) {
	var subReq vo.Subject
	_ = ctx.ShouldBindJSON(&subReq)
	if err := validate.Verify(subReq, validate.AddSubjectVerify); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err := subjectService.Insert(subReq)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

func (this *SubjectApi) Update(ctx *gin.Context) {
	var subReq vo.Subject
	_ = ctx.ShouldBindJSON(&subReq)
	if err := validate.Verify(subReq, validate.UpdateSubjectVerify); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err := subjectService.Update(subReq)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

func (this *SubjectApi) Delete(ctx *gin.Context) {
	subIdStr := ctx.Request.URL.Query().Get("id")
	if subIdStr == "" {
		response.FailWithMessage(ctx, "id不能为空")
		return
	}
	subId, err := strconv.Atoi(subIdStr)
	if err != nil {
		response.FailWithMessage(ctx, "id必须是数字类型")
		return
	}
	err = subjectService.DeleteById(uint(subId))
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}
