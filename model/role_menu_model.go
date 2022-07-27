package model

type RoleMenu struct {
	MenuId uint `gorm:"column:menu_id;comment:菜单ID"`
	RoleId uint `gorm:"column:role_id;comment:角色ID"`
}

func (this *RoleMenu) TableName()string {
	return "role_menu"
}