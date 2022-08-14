package auth

import (
	"memoirs/model/auth"
	"memoirs/model/vo"
	"memoirs/pkg/constant"
	"memoirs/utils"
)

type MenuService struct{}

func (menuService *MenuService) QueryUserMenu(userId, parentId uint) ([]vo.MenuTree, error) {
	menuList, err := menuMapper.QueryFirstMenuInfo(userId, parentId)
	if err != nil {
		return nil, err
	}
	// 结构体属性copy
	var menuTree []vo.MenuTree
	for _, menu := range menuList {
		var menuInfo vo.MenuTree
		err = utils.CopyProperties(&menu, &menuInfo)
		menuInfo.ID = menu.ID
		menuInfo.Children = []vo.MenuTree{}
		menuTree = append(menuTree, menuInfo)
	}
	return menuTree, err
}

func (menuService *MenuService) AddMenu(menuReq vo.AddMenuRequest) error {
	var menu auth.Menu
	err := utils.CopyProperties(&menuReq, &menu)
	if err != nil {
		return err
	}
	if menuReq.MenuId == constant.SUPER_PARENT_ID {
		err = menuMapper.AddMenu(menu)
	} else {
		menu.ID = menuReq.MenuId
		err = menuMapper.Update(menu)
	}
	return err
}

func (menuService *MenuService) BuildMenuTree(userId uint) ([]vo.MenuTree, error) {
	menuList, err := menuMapper.QueryMenuInfo(userId)
	if err != nil {
		return nil, err
	}
	// 结构体属性copy
	var menuTree []vo.MenuTree
	for _, menu := range menuList {
		var menuInfo vo.MenuTree
		err = utils.CopyProperties(&menu, &menuInfo)
		menuInfo.ID = menu.ID
		menuInfo.Children = []vo.MenuTree{}
		menuTree = append(menuTree, menuInfo)
	}
	// 去重
	var menus []vo.MenuTree
	temp := map[uint]struct{}{}
	for _, menu := range menuTree {
		if _, ok := temp[menu.ID]; !ok {
			temp[menu.ID] = struct{}{}
			menus = append(menus, menu)
		}
	}
	resp := BuildTree(menus, constant.SUPER_PARENT_ID)
	return resp, nil
}

func BuildTree(list []vo.MenuTree, rootId uint) []vo.MenuTree {
	var nodeList []vo.MenuTree
	for _, item := range list {
		if item.ParentId == rootId {
			item.Children = BuildTree(list, item.ID)
			nodeList = append(nodeList, item)
		}
	}
	return nodeList
}

func (menuService *MenuService) DeleteMenu(menuIds vo.DeleteMenu) error {
	err := menuMapper.DeleteMenu(menuIds.MenuIds)
	return err
}

func (menuService *MenuService) IsHidden(menuId uint, hidden *bool) error {
	err := menuMapper.UpdateByHidden(menuId, hidden)
	return err
}

func (menuService *MenuService) SortMenu(req []vo.SortMenu) (err error) {
	for _, item := range req {
		err = menuMapper.UpdateBySort(item.MenuId, item.Sort, item.ParentId)
	}
	return
}
