package api

import (
	"github.com/gin-gonic/gin"
	"memoirs/common/response"
	"memoirs/model"
	"memoirs/model/vo"
	"memoirs/utils"
)

type RoleApi struct{}

func (this *RoleApi) QueryRoleAllList(ctx *gin.Context) {
	list, err := roleService.QueryList()
	if err != nil {
		response.FailWithMessage(ctx, "查询角色列表失败")
		return
	}
	response.OkWithData(ctx, list)
}

func (this *RoleApi) QueryUserRoleList(ctx *gin.Context) {
	userId := utils.GetUserID(ctx)
	list, err := roleService.QueryUserRoleList(userId)
	if err != nil {
		response.FailWithMessage(ctx, "查询角色列表失败")
		return
	}
	response.OkWithData(ctx, list)
}

func (this *RoleApi) AddRole(ctx *gin.Context) {
	var role vo.RoleRequest
	_ = ctx.ShouldBindJSON(&role)
	var roled model.Role
	_ = utils.CopyProperties(&role, &roled)
	err := roleService.AddRole(roled)
	if err != nil {
		response.FailWithMessage(ctx, "新增角色失败")
		return
	}
	response.Ok(ctx)
}

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

func (this *RoleApi) UpdateRole(ctx *gin.Context) {
	var roleReq vo.RoleUpdate
	_ = ctx.ShouldBindJSON(&roleReq)
	var role model.Role
	_ = utils.CopyProperties(&roleReq, &role)
	err := roleService.UpdateRole(role)
	if err != nil {
		response.FailWithMessage(ctx, "更新失败！")
		return
	}
	response.Ok(ctx)
}

// 删除角色
func (this *RoleApi) DeleteRole(ctx *gin.Context) {
	var roleReq vo.DeletedRole
	_ = ctx.ShouldBindJSON(&roleReq)
	err := roleService.DeleteRole(roleReq.RoleId)
	if err != nil {
		response.FailWithMessage(ctx, "删除角色失败")
		return
	}
	response.Ok(ctx)
}
