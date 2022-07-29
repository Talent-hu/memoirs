package api

import (
	"github.com/gin-gonic/gin"
	"memoirs/common/response"
	"memoirs/model/vo"
)

type RoleApi struct{}

func (this *RoleApi) AddRoleAndMenu(ctx *gin.Context) {
	var relation vo.RoleMenuRelation
	_ = ctx.ShouldBindJSON(&relation)
	err := roleService.AddRoleAndMenu(relation.RoleId, relation.MenuIds)
	if err != nil {
		response.FailWithMessage(ctx, "新增权限失败")
		return
	}
	response.Ok(ctx)
}
