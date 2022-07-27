package service

import (
	"memoirs/global"
	"memoirs/model"
)

type RoleService struct {}


func (this *RoleService) GetRoleInfo(userId uint) ([]*model.Role, error) {
	var roleList []*model.Role
	err := global.DB.Model(&model.Role{}).Select("role.role_code,role.role_name,role.parent_id,role.id").
		Joins("left join user_role on user_role.role_id = role.id").
		Where("user_role.user_id = ?", userId).
		Preload("Menus").
		Find(&roleList).Error
    return roleList,err
}