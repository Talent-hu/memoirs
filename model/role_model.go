package model

type Role struct {
	BaseModel
	RoleCode string `gorm:"varchar(32);unique;comment:角色编码" json:"roleCode"`
	RoleName string `gorm:"varchar(32);comment:角色名" json:"roleName"`
	ParentId uint   `gorm:"comment:父角色ID" json:"parentId"`
	Children []Role  `gorm:"-" json:"children"`
	Menus    []Menu  `gorm:"many2many:role_menu" json:"menus"`
	Users    []User  `gorm:"many2many:user_role" json:"users"`
}


func (this *Role) TableName() string {
	return  "role"
}