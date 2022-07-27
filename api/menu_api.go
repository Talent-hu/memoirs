package api

import (
	"github.com/gin-gonic/gin"
	"memoirs/common/constant"
	"memoirs/common/response"
	"memoirs/model"
	"memoirs/model/vo"
	"memoirs/utils"
)

type MenuApi struct {}


func (this *MenuApi) QueryMenuTree(ctx *gin.Context) {
	userId := utils.GetUserID(ctx)
	roleList, err := roleService.GetRoleInfo(userId)
	if err != nil {
		response.FailWithMessage(ctx,"系统错误，未查询到对应数据。")
		return
	}
	for _, role := range roleList {
		if constant.SUPER_PARENT_ID == role.ParentId {

		}
	}
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
		response.FailWithMessage(ctx,"新增菜单失败！")
		return
	}
	response.Ok(ctx)
}

func (this *MenuApi) BuildMenuList(ctx *gin.Context) {
	userId := utils.GetUserID(ctx)
	roleList, err := roleService.GetRoleInfo(userId)
	if err != nil {
		response.FailWithMessage(ctx,"系统错误，未查询到对应数据。")
		return
	}
	var menuList []vo.QueryMenuListResponse
	for _, role := range roleList {
		menus := role.Menus
		for _, menu := range menus {
			var menuInfo vo.QueryMenuListResponse
			utils.CopyProperties(&menu,&menuInfo)
			menuInfo.ID = menu.ID
			menuList = append(menuList,menuInfo)
		}
	}
	// 去重
	var menuLst []vo.QueryMenuListResponse
	temp := map[uint]struct{}{}
	for _, menu := range menuList {
		if _,ok := temp[menu.ID];!ok {
			temp[menu.ID] = struct{}{}
			menuLst = append(menuLst,menu)
		}
	}
	treeList := BuildTree(menuLst, constant.SUPER_PARENT_ID)
	response.OkWithData(ctx,treeList)
}

func BuildTree(menuList []vo.QueryMenuListResponse,rootId uint) []vo.QueryMenuListResponse {
	var result []vo.QueryMenuListResponse
	for _,menu := range menuList {
		if menu.ParentId == rootId {
			result = append(result, findChild(menu,menuList))
		}
	}
	return result
}

func findChild(rootNode vo.QueryMenuListResponse,menuList []vo.QueryMenuListResponse) vo.QueryMenuListResponse {
	for _, menuNode := range menuList {
		if rootNode.ID == menuNode.ParentId {
			if menuNode.Children == nil {
				menuNode.Children = nil
			}
			rootNode.Children = append(rootNode.Children, findChild(menuNode,menuList))
		}
	}
	return rootNode
}