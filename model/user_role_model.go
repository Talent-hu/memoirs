package model

type UserRole struct {
	UserId uint `gorm:"column:user_id;comment:用户ID"`
	RoleId uint `gorm:"column:role_id;comment:角色ID"`
}


func (this *UserRole) TableName() string{
	return "user_role"
}