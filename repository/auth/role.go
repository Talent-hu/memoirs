package auth

import (
	"memoirs/global"
	"memoirs/model"
)

type RoleRepoInterface interface {
	GetRoleInfo(userId uint) ([]*model.Role, error)
	DelRelation(roleId uint, menuIds []uint) error
	AddRoleAndMenu(roleMenus []model.RoleMenu) error
	AddRole(role model.Role) error
	QueryList(pageSize, offset int) ([]model.Role, int64)
	QueryUserRoleList(userId uint) ([]model.Role, error)
	UpdateRole(role model.Role) error
	DeleteRole(roleId uint) error
}

type RoleRepository struct{}

func (this *RoleRepository) GetRoleInfo(userId uint) ([]*model.Role, error) {
	var roleList []*model.Role
	err := global.DB.Model(&model.Role{}).Select("role.role_code,role.role_name,role.parent_id,role.id").
		Joins("left join user_role on user_role.role_id = role.id").
		Where("user_role.user_id = ?", userId).
		Preload("Menus").
		Find(&roleList).Error
	return roleList, err
}

func (this *RoleRepository) DelRelation(roleId uint, menuIds []uint) error {
	var roleMenu model.RoleMenu
	err := global.DB.Where("role_id = ? and menu_id in (?)", roleId, menuIds).
		Delete(&roleMenu).Error
	return err
}

func (this *RoleRepository) AddRoleAndMenu(roleMenus []model.RoleMenu) error {
	err := global.DB.Model(&model.RoleMenu{}).CreateInBatches(roleMenus, len(roleMenus)).Error
	return err
}

func (this *RoleRepository) AddRole(role model.Role) error {
	err := global.DB.Create(&role).Error
	return err
}

func (this *RoleRepository) QueryList(pageSize, offset int) ([]model.Role, int64) {
	var roleList []model.Role
	var count int64
	global.DB.Limit(pageSize).Offset(offset).Find(&roleList)
	global.DB.Model(&model.Role{}).Count(&count)
	return roleList, count
}

func (this *RoleRepository) QueryUserRoleList(userId uint) ([]model.Role, error) {
	var roleList []model.Role
	err := global.DB.Model(&model.UserRole{}).
		Joins("left join role on role.id = user_role.role_id").
		Where("user_role.user_id=?", userId).
		Find(&roleList).Error
	return roleList, err
}

func (this *RoleRepository) UpdateRole(role model.Role) error {
	err := global.DB.Table("role").
		Where("id = ?", role.ID).
		Updates(role).Error
	return err
}

func (this *RoleRepository) DeleteRole(roleId uint) error {
	tx := global.DB.Begin()
	err := tx.Delete(&model.Role{}, roleId).Error
	err = tx.Delete(&model.RoleMenu{}, roleId).Error
	tx.Commit()
	if err != nil {
		tx.Callback()
	}
	return err
}
