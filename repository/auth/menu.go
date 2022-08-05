package auth

import (
	"memoirs/global"
	"memoirs/model"
	"memoirs/pkg/constant"
)

type MenuRepoInterface interface {
	AddMenu(menu model.Menu) error
	QueryMenuInfo(userId uint) ([]model.Menu, error)
	DeleteMenu(menuIds []uint) (model.Menu, error)
}

type MenuRepository struct{}

func (this *MenuRepository) AddMenu(menu model.Menu) error {
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

func (this *MenuRepository) QueryMenuInfo(userId uint) ([]model.Menu, error) {
	var menuList []model.Menu
	err := global.DB.Table("user_role").
		Joins("left join role_menu on role_menu.role_id = user_role.role_id").
		Joins("left join menu on menu.id = role_menu.menu_id").
		Where("user_role.user_id = ?", userId).
		Scan(&menuList).Error
	return menuList, err
}

func (this *MenuRepository) QueryFirstMenuInfo(userId, superMenuId uint) ([]model.Menu, error) {
	var menuList []model.Menu
	err := global.DB.Table("user_role").
		Joins("left join role_menu on role_menu.role_id = user_role.role_id").
		Joins("left join menu on menu.id = role_menu.menu_id").
		Where("user_role.user_id = ? and menu.parent_id = ? and menu.has_btn = 1", userId, superMenuId).
		Scan(&menuList).Error
	return menuList, err
}

func (this *MenuRepository) DeleteMenu(menuIds []uint) (model.Menu, error) {
	var menus model.Menu
	err := global.DB.Where("id in (?)", menuIds).Delete(&menus).Error
	return menus, err
}

func (this *MenuRepository) QuerySetting(hasBtn uint) (model.Menu, error) {
	var menu model.Menu
	err := global.DB.Model(&model.Menu{}).
		Where("has_btn = ?", hasBtn).First(&menu).Error
	return menu, err
}
