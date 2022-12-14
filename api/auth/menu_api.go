package auth

import (
	"github.com/gin-gonic/gin"
	"memoirs/model/vo"
	"memoirs/pkg/constant"
	"memoirs/pkg/response"
	"memoirs/utils"
	"memoirs/validate"
	"strconv"
)

type MenuApi struct{}

// @Tag MenuApi
// @Summary 查询用户一级菜单列表
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {object} response.Response{data=vo.MenuTree,message=string} "返回用户token"
// @Router /menu/list [post]
func (this *MenuApi) QueryMenuList(ctx *gin.Context) {
	userId := utils.GetUserID(ctx)
	menuId := ctx.Request.URL.Query().Get("menuId")
	var parentId uint
	if menuId == "" {
		parentId = constant.SUPER_PARENT_ID
	} else {
		id, _ := strconv.Atoi(menuId)
		parentId = uint(id)
	}
	resp, err := menuService.QueryUserMenu(userId, parentId)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, resp)
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
	_ = ctx.ShouldBindJSON(&menuReq)
	/*	if err := validate.Verify(menuReq,validate.AddMenuVerify);err != nil {
		response.FailWithMessage(ctx,err.Error())
		return
	}*/
	err := menuService.AddMenu(menuReq)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
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
	resp, err := menuService.BuildMenuTree(userId)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, resp)
}

func (this *MenuApi) DeleteMenu(ctx *gin.Context) {
	var delMenu vo.DeleteMenu
	_ = ctx.ShouldBindJSON(&delMenu)
	if err := validate.Verify(delMenu, validate.DeleteMenuVerify); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err := menuService.DeleteMenu(delMenu)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

func (this *MenuApi) IsHidden(ctx *gin.Context) {
	var req vo.IsHidden
	_ = ctx.ShouldBindJSON(&req)
	if err := validate.Verify(req, validate.IsHiddenVerify); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err := menuService.IsHidden(req.ID, req.Hidden)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

func (this *MenuApi) SortMenu(ctx *gin.Context) {
	var req vo.SortMenuList
	_ = ctx.ShouldBindJSON(&req)
	err := menuService.SortMenu(req.SortMenuList)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}
