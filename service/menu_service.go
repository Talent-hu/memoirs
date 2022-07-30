package service

import (
	"memoirs/common/constant"
	"memoirs/global"
	"memoirs/model"
	"memoirs/model/vo"
	"memoirs/utils"
)

type MenuService struct{}

func (this *MenuService) AddMenu(menu *model.Menu) error {
	db := global.DB.Begin()
	err := db.Create(&menu).Error
	// 给超级管理员赋予权限
	var roleMenu model.RoleMenu
	roleMenu.MenuId = menu.ID
	roleMenu.RoleId = constant.ROOT_ROLE_ID
	err = db.Create(&roleMenu).Error
	db.Commit()
	if err != nil {
		global.Log.Error(err.Error())
		db.Callback()
	}
	return err
}

func (this *MenuService) QueryMenuInfo(userId uint) ([]vo.MenuTree, error) {
	var menuList []model.Menu
	err := global.DB.Table("user_role").
		Joins("left join role_menu on role_menu.role_id = user_role.role_id").
		Joins("left join menu on menu.id = role_menu.menu_id").
		Where("user_role.user_id = ?", userId).
		Scan(&menuList).Error
	if err != nil {
		return nil, err
	}
	// 结构体属性copy
	var menuTree []vo.MenuTree
	for _, menu := range menuList {
		var menuInfo vo.MenuTree
		_ = utils.CopyProperties(&menu, &menuInfo)
		menuInfo.ID = menu.ID
		menuInfo.Children = []vo.MenuTree{}
		menuTree = append(menuTree, menuInfo)
	}
	// 去重
	var menuLst []vo.MenuTree
	temp := map[uint]struct{}{}
	for _, menu := range menuTree {
		if _, ok := temp[menu.ID]; !ok {
			temp[menu.ID] = struct{}{}
			menuLst = append(menuLst, menu)
		}
	}
	return menuLst, err
}

func (this *MenuService) DeleteMenu(menuIds []uint) (model.Menu, error) {
	var menus model.Menu
	err := global.DB.Where("id in (?)", menuIds).Delete(&menus).Error
	return menus, err
}
