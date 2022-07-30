package service

import (
	"memoirs/model"
)

type RoleService struct{}

func (this *RoleService) GetRoleInfo(userId uint) ([]*model.Role, error) {
	var roleList []*model.Role
	err := db.Model(&model.Role{}).Select("role.role_code,role.role_name,role.parent_id,role.id").
		Joins("left join user_role on user_role.role_id = role.id").
		Where("user_role.user_id = ?", userId).
		Preload("Menus").
		Find(&roleList).Error
	return roleList, err
}

func (this *RoleService) DelRelation(roleId uint, menuIds []uint) error {
	var roleMenu model.RoleMenu
	err := db.Where("role_id = ? and menu_id in (?)", roleId, menuIds).
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
	err := db.Model(&model.RoleMenu{}).CreateInBatches(roleMenus, len(roleMenus)).Error
	return err
}

func (this *RoleService) AddRole(role model.Role) error {
	err := db.Create(&role).Error
	return err
}

func (this *RoleService) QueryList() ([]model.Role, error) {
	var roleList []model.Role
	err := db.Find(&roleList).Error
	return roleList, err
}

func (this *RoleService) QueryUserRoleList(userId uint) ([]model.Role, error) {
	var roleList []model.Role
	err := db.Model(&model.UserRole{}).
		Joins("left join role on role.id = user_role.role_id").
		Where("user_role.user_id=?", userId).
		Find(&roleList).Error
	return roleList, err
}

func (this *RoleService) UpdateRole(role model.Role) error {
	err := db.Table("role").
		Where("id = ?", role.ID).
		Updates(role).Error
	return err
}

func (this *RoleService) DeleteRole(roleId uint) error {
	tx := db.Begin()
	err := tx.Delete(&model.Role{}, roleId).Error
	err = tx.Delete(&model.RoleMenu{}, roleId).Error
	tx.Commit()
	if err != nil {
		tx.Callback()
	}
	return err
}
