package api

import (
	"github.com/gin-gonic/gin"
	"memoirs/model/vo"
	"memoirs/pkg/response"
	"memoirs/utils"
	"memoirs/validate"
)

type RoleApi struct{}

// 分页查询所有的角色列表
func (this *RoleApi) QueryRoleAllList(ctx *gin.Context) {
	var queryPage vo.ListQuery
	_ = ctx.ShouldBindJSON(&queryPage)
	if err := validate.Verify(queryPage, validate.QueryPageListVerify); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	resp := roleService.QueryAll(queryPage)
	response.OkWithData(ctx, resp)
}

// 查询用户拥有的角色信息
func (this *RoleApi) QueryUserRoleList(ctx *gin.Context) {
	userId := utils.GetUserID(ctx)
	resp, err := roleService.QueryUserRole(userId)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, resp)
}

// 新增角色
func (this *RoleApi) AddRole(ctx *gin.Context) {
	var role vo.RoleRequest
	_ = ctx.ShouldBindJSON(&role)
	if err := validate.Verify(role, validate.AddRoleVerify); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err := roleService.AddRole(role)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

// 新增角色与菜单关系
func (this *RoleApi) AddRoleAndMenu(ctx *gin.Context) {
	var relation vo.RoleMenuRelation
	_ = ctx.ShouldBindJSON(&relation)
	if err := validate.Verify(relation, validate.RoleAndMenuVerify); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err := roleService.AddRoleAndMenu(relation)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

// 更新角色信息
func (this *RoleApi) UpdateRole(ctx *gin.Context) {
	var roleReq vo.RoleInfo
	_ = ctx.ShouldBindJSON(&roleReq)
	if err := validate.Verify(roleReq, validate.UpdateRoleVerify); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err := roleService.UpdateRole(roleReq)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

// 删除角色
func (this *RoleApi) DeleteRole(ctx *gin.Context) {
	var roleReq vo.DeletedRole
	_ = ctx.ShouldBindJSON(&roleReq)
	if err := validate.Verify(roleReq, validate.DeleteRoleVerify); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err := roleService.DeleteRole(roleReq.RoleId)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

// 删除角色与菜单的关联关系
func (this *RoleApi) DelMenuAndRoleRel(ctx *gin.Context) {
	var relation vo.RoleMenuRelation
	_ = ctx.ShouldBindJSON(&relation)
	if err := validate.Verify(relation, validate.RoleAndMenuVerify); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err := roleService.DeleteRoleAndMenu(relation)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}
