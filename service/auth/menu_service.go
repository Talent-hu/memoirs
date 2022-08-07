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
	err = menuMapper.AddMenu(menu)
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

func (menuService *MenuService) DeleteMenu(menuIds vo.DeleteMenu) error {
	_, err := menuMapper.DeleteMenu(menuIds.MenuIds)
	return err
}
