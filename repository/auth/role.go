package auth

import (
	"memoirs/global"
	"memoirs/model/auth"
)

type RoleRepoInterface interface {
	GetRoleInfo(userId uint) ([]*auth.Role, error)
	DelRelation(roleId uint, menuIds []uint) error
	AddRoleAndMenu(roleMenus []auth.RoleMenu) error
	AddRole(role auth.Role) error
	QueryList(pageSize, offset int) ([]auth.Role, int64)
	QueryUserRoleList(userId uint) ([]auth.Role, error)
	UpdateRole(role auth.Role) error
	DeleteRole(roleId uint) error
}

type RoleRepository struct{}

func (this *RoleRepository) GetRoleInfo(userId uint) ([]*auth.Role, error) {
	var roleList []*auth.Role
	err := global.DB.Model(&auth.Role{}).Select("role.role_code,role.role_name,role.parent_id,role.id").
		Joins("left join user_role on user_role.role_id = role.id").
		Where("user_role.user_id = ?", userId).
		Preload("Menus").
		Find(&roleList).Error
	return roleList, err
}

func (this *RoleRepository) DelRelation(roleId uint, menuIds []uint) error {
	var roleMenu auth.RoleMenu
	err := global.DB.Where("role_id = ? and menu_id in (?)", roleId, menuIds).
		Delete(&roleMenu).Error
	return err
}

func (this *RoleRepository) AddRoleAndMenu(roleMenus []auth.RoleMenu) error {
	err := global.DB.Model(&auth.RoleMenu{}).CreateInBatches(roleMenus, len(roleMenus)).Error
	return err
}

func (this *RoleRepository) AddRole(role auth.Role) error {
	err := global.DB.Create(&role).Error
	return err
}

func (this *RoleRepository) QueryList(pageSize, offset int) ([]auth.Role, int64) {
	var roleList []auth.Role
	var count int64
	global.DB.Limit(pageSize).Offset(offset).Find(&roleList)
	global.DB.Model(&auth.Role{}).Count(&count)
	return roleList, count
}

func (this *RoleRepository) QueryUserRoleList(userId uint) ([]auth.Role, error) {
	var roleList []auth.Role
	err := global.DB.Model(&auth.UserRole{}).
		Joins("left join role on role.id = user_role.role_id").
		Where("user_role.user_id=?", userId).
		Find(&roleList).Error
	return roleList, err
}

func (this *RoleRepository) UpdateRole(role auth.Role) error {
	err := global.DB.Table("role").
		Where("id = ?", role.ID).
		Updates(role).Error
	return err
}

func (this *RoleRepository) DeleteRole(roleId uint) error {
	tx := global.DB.Begin()
	err := tx.Delete(&auth.Role{}, roleId).Error
	err = tx.Delete(&auth.RoleMenu{}, roleId).Error
	tx.Commit()
	if err != nil {
		tx.Callback()
	}
	return err
}
