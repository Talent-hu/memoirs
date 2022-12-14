package auth

import "memoirs/model"

type Menu struct {
	model.BaseModel
	Path       string `gorm:"type:varchar(64);unique;comment:路由path" json:"path"`
	Name       string `gorm:"type:varchar(64);comment:路由名称" json:"name"`
	Component  string `gorm:"type:varchar(64);comment:对应前端文件路径" json:"component"`
	Hidden     *bool  `gorm:"default:0;comment:是否隐藏列表" json:"hidden"`
	Sort       uint   `gorm:"type:int;comment:排序标记" json:"sort"`
	Title      string `gorm:"type:varchar(32);comment:菜单名称" json:"title"`
	Icon       string `gorm:"comment:菜单图标" json:"icon"`
	ParentId   uint   `gorm:"comment:父级菜单ID" json:"parentId"`
	HasBtn     *bool  `gorm:"comment:是否是按钮" json:"hasBtn"`
	Permission string `gorm:"type:varchar(64);comment:权限数据" json:"permission"`
	MenuLevel  uint   `gorm:"type:int;comment:菜单级别" json:"-"`
	Children   []Menu `gorm:"-" json:"children"`
	Roles      []Role `gorm:"many2many:role_menu"`
}

func (this *Menu) TableName() string {
	return "menu"
}
