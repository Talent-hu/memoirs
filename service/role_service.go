package service

import (
	"memoirs/global"
	"memoirs/model"
)

type RoleService struct{}

func (this *RoleService) GetRoleInfo(userId uint) ([]*model.Role, error) {
	var roleList []*model.Role
	err := global.DB.Model(&model.Role{}).Select("role.role_code,role.role_name,role.parent_id,role.id").
		Joins("left join user_role on user_role.role_id = role.id").
		Where("user_role.user_id = ?", userId).
		Preload("Menus").
		Find(&roleList).Error
	return roleList, err
}

func (this *RoleService) DelRelation(roleId uint, menuIds []uint) error {
	var roleMenu model.RoleMenu
	err := global.DB.Where("role_id = ? and menu_id in (?)", roleId, menuIds).
		Delete(&roleMenu).Error
	return err
}

func (this *RoleService) AddRoleAndMenu(roleId uint, menuIds []uint) error {
	var roleMenus []model.RoleMenu
	for _, menuId := range menuIds {
		roleMenu := new(model.RoleMenu)
		roleMenu.RoleId = roleId
		roleMenu.MenuId = menuId
		roleMenus = append(roleMenus, *roleMenu)
	}
	err := global.DB.Model(&model.RoleMenu{}).CreateInBatches(roleMenus, len(roleMenus)).Error
	return err
}
