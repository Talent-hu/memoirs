package api

import (
	"github.com/gin-gonic/gin"
	"memoirs/common/constant"
	"memoirs/common/response"
	"memoirs/model"
	"memoirs/model/vo"
	"memoirs/utils"
)

type MenuApi struct{}

func (this *MenuApi) QueryMenuList(ctx *gin.Context) {
	userId := utils.GetUserID(ctx)
	menus, err := menuService.QueryMenuInfo(userId)
	if err != nil {
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
