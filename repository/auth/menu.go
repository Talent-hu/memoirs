package auth

import (
	"errors"
	"gorm.io/gorm"
	"memoirs/global"
	"memoirs/model/auth"
	"memoirs/pkg/constant"
)

type MenuRepoInterface interface {
	AddMenu(menu auth.Menu) error
	QueryMenuInfo(userId uint) ([]auth.Menu, error)
	DeleteMenu(menuIds []uint) (auth.Menu, error)
}

type MenuRepository struct{}

func (this *MenuRepository) AddMenu(menu auth.Menu) error {
	db := global.DB.Begin()
	err := db.Create(&menu).Error
	// 给超级管理员赋予权限
	var roleMenu auth.RoleMenu
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

func (this *MenuRepository) Update(menu auth.Menu) error {
	var menuModel auth.Menu
	if errors.Is(global.DB.First(&menuModel, menu.ID).Error, gorm.ErrRecordNotFound) {
		return errors.New("该菜单数据不存在")
	}
	err := global.DB.Model(&auth.Menu{}).
		Where("id=?", menu.ID).
		Updates(menu).Error
	return err
}

func (this *MenuRepository) QueryMenuInfo(userId uint) ([]auth.Menu, error) {
	var menuList []auth.Menu
	err := global.DB.Table("user_role").
		Joins("LEFT JOIN role_menu ON role_menu.role_id = user_role.role_id").
		Joins("INNER JOIN menu ON menu.id = role_menu.menu_id").
		Where("user_role.user_id = ?", userId).
		Order("menu.sort ASC,created_at ASC").
		Scan(&menuList).Error
	return menuList, err
}

func (this *MenuRepository) QueryFirstMenuInfo(userId, superMenuId uint) ([]auth.Menu, error) {
	var menuList []auth.Menu
	err := global.DB.Table("user_role").
		Joins("LEFT JOIN role_menu ON role_menu.role_id = user_role.role_id").
		Joins("INNER JOIN menu ON menu.id = role_menu.menu_id").
		Where("user_role.user_id = ? AND menu.parent_id = ? AND menu.has_btn = 0 AND menu.hidden=0", userId, superMenuId).
		Order("menu.sort ASC,created_at ASC").
		Scan(&menuList).Error
	return menuList, err
}

func (this *MenuRepository) DeleteMenu(menuIds []uint) error {
	var menus auth.Menu
	err := global.DB.Unscoped().Where("id in (?)", menuIds).Delete(&menus).Error
	if err != nil {
		return err
	}
	err = global.DB.Unscoped().Where("menu_id in (?)", menuIds).Delete(&auth.RoleMenu{}).Error
	return err
}

func (this *MenuRepository) UpdateByHidden(menuId uint, hidden *bool) error {
	var menuModel auth.Menu
	if errors.Is(global.DB.First(&menuModel, menuId).Error, gorm.ErrRecordNotFound) {
		return errors.New("该菜单数据不存在")
	}
	err := global.DB.Model(&auth.Menu{}).
		Where("id=?", menuId).
		Update("hidden", hidden).Error
	return err
}

func (this *MenuRepository) UpdateBySort(menuId, sort, parentId uint) error {
	var menuModel auth.Menu
	if errors.Is(global.DB.First(&menuModel, menuId).Error, gorm.ErrRecordNotFound) {
		return errors.New("该菜单数据不存在")
	}
	err := global.DB.Model(&auth.Menu{}).
		Where("id=?", menuId).
		Update("sort", sort).
		Update("parent_id", parentId).Error
	return err
}
