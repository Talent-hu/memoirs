package auth

import "memoirs/model"

type SysDictItem struct {
	model.BaseModel
	Name     string `gorm:"varchar(50);comment:字典项名称" json:"name"`
	Value    string `gorm:"varchar(50);comment:字典项值" json:"value"`
	DictCode string `gorm:"varchar(50);comment:字典编码" json:"dictCode"`
	Sort     int    `gorm:"int;default 0;comment:排序" json:"sort"`
	Status   *bool  `gorm:"default 0;comment:状态（0 停用 1正常）" json:"status"`
	Ext      string `gorm:"comment:预留字段" json:"ext"`
	Remark   string `gorm:"comment:备注" json:"remark"`
}

func (this *SysDictItem) TableName() string {
	return "sys_dict_item"
}
