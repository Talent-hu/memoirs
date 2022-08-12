package auth

import "memoirs/model"

type SysFunction struct {
	model.BaseModel
	Path    string `gorm:"type:varchar(128);comment:请求路径" json:"path"`
	Name    string `gorm:"type:varchar(64);comment:功能名称" json:"name"`
	Method  string `gorm:"type:varchar(10);comment:请求方式" json:"method"`
	CrtUser uint   `gorm:"comment:创建用户" json:"crtUser"`
	UpdUser uint   `gorm:"comment:更新用户" json:"updUser"`
}

func (model *SysFunction) TableName() string {
	return "sys_function"
}
