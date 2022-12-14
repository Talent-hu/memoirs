package auth

import "memoirs/model"

type SysDict struct {
	model.BaseModel
	Code         string        `gorm:"varchar(64);unique;comment:类型编码" json:"code"`
	Name         string        `gorm:"varchar(255);comment:类型名称" json:"value"`
	Status       *bool         `gorm:"comment:状态（0 停用 1正常）" json:"status"`
	Remark       string        `gorm:"comment:备注" json:"remark"`
	SysDictItems []SysDictItem `gorm:"foreignKey:DictCode" json:"-"`
}

func (this *SysDict) TableName() string {
	return "sys_dict"
}
