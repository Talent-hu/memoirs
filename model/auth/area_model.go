package auth

import "memoirs/model"

type Area struct {
	model.BaseTreeModel
	Name     string `gorm:"type:varchar(64);comment:地区名称" json:"name"`
	Code     string `gorm:"type:varchar(20);comment:编码" json:"code"`
	Level    uint   `gorm:"type:int;comment:级别" json:"level"`
	Sort     uint   `gorm:"type:int;comment:排序" json:"sort"`
	Children []Area `gorm:"-" json:"children"`
}

func (this *Area) TableName() string {
	return "area"
}
