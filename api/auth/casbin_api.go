package auth

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"memoirs/global"
	"memoirs/model/vo"
	"memoirs/pkg/response"
	"memoirs/validate"
)

type CasbinApi struct{}

func (cas *CasbinApi) UpdateCasbin(ctx *gin.Context) {
	var cr vo.CasbinInReceive
	_ = ctx.ShouldBindJSON(&cr)
	if err := validate.Verify(cr, validate.CasbinVerify); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	if err := casbinService.UpdateCasbin(cr.RoleCode, cr.CasbinInfos); err != nil {
		global.Log.Error("权限数据更新失败", zap.Error(err))
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

func (cas *CasbinApi) GetPolicyPathByRoleId(ctx *gin.Context) {
	roleCode := ctx.Request.URL.Query().Get("roleCode")
	if roleCode == "" {
		response.FailWithMessage(ctx, "roleCode不能为空")
		return
	}
	resp := casbinService.GetPolicyPathByRoleId(roleCode)
	response.OkWithData(ctx, resp)
}
