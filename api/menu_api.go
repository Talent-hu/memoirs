package api

import (
	"github.com/gin-gonic/gin"
	"memoirs/common/constant"
	"memoirs/common/response"
	"memoirs/global"
	"memoirs/model"
	"memoirs/model/vo"
	"memoirs/utils"
)

type MenuApi struct{}

// @Tag MenuApi
// @Summary 查询用户下菜单列表
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {object} response.Response{data=vo.MenuTree,message=string} "返回用户token"
// @Router /menu/list [post]
func (this *MenuApi) QueryMenuList(ctx *gin.Context) {
	userId := utils.GetUserID(ctx)
	menus, err := menuService.QueryMenuInfo(userId)
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(ctx, "系统错误，未查询到对应数据。")
		return
	}
	response.OkWithData(ctx, menus)

}

// @Tag MenuApi
// @Summary 新增菜单
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param data body vo.AddMenuRequest false "查询参数"
// @Success 200 {object} response.Response{data=string,message=string} "返回用户token"
// @Router /menu/add [post]
func (this *MenuApi) AddMenu(ctx *gin.Context) {
	var menuReq vo.AddMenuRequest
	var menu model.Menu
	ctx.ShouldBindJSON(&menuReq)
	_ = utils.CopyProperties(&menuReq, &menu)
	err := menuService.AddMenu(&menu)
	if err != nil {
		response.FailWithMessage(ctx, "新增菜单失败！")
		return
	}
	response.Ok(ctx)
}

// @Tag MenuApi
// @Summary 查询用户下菜单树
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {object} response.Response{data=vo.MenuTree,message=string} "返回用户token"
// @Router /menu/tree [post]
func (this *MenuApi) QueryMenuTree(ctx *gin.Context) {
	userId := utils.GetUserID(ctx)
	menus, err := menuService.QueryMenuInfo(userId)
	if err != nil {
		response.FailWithMessage(ctx, "系统错误，未查询到对应数据。")
		return
	}
	treeList := BuildTree(menus, constant.SUPER_PARENT_ID)
	response.OkWithData(ctx, treeList)
}

func BuildTree(menuList []vo.MenuTree, rootId uint) []vo.MenuTree {
	var result []vo.MenuTree
	for _, menu := range menuList {
		if menu.ParentId == rootId {
			result = append(result, findChild(menu, menuList))
		}
	}
	return result
}

func findChild(rootNode vo.MenuTree, menuList []vo.MenuTree) vo.MenuTree {
	for _, menuNode := range menuList {
		if rootNode.ID == menuNode.ParentId {
			if menuNode.Children == nil {
				menuNode.Children = []vo.MenuTree{}
			}
			rootNode.Children = append(rootNode.Children, findChild(menuNode, menuList))
		}
	}
	return rootNode
}

func (this *MenuApi) RemoveMenu(ctx *gin.Context) {
	var menuIds vo.RemoveMenu
	_ = ctx.ShouldBindJSON(&menuIds)
	_, err := menuService.DeleteMenu(menuIds.MenuIds)
	if err != nil {
		response.FailWithMessage(ctx, "删除菜单失败")
		return
	}
	response.Ok(ctx)
}

func (this *MenuApi) DelMenuAndRoleRel(ctx *gin.Context) {
	var delRel vo.RoleMenuRelation
	_ = ctx.ShouldBindJSON(&delRel)
	err := roleService.DelRelation(delRel.RoleId, delRel.MenuIds)
	if err != nil {
		response.FailWithMessage(ctx, "删除菜单权限失败！")
		return
	}
	response.Ok(ctx)
}
