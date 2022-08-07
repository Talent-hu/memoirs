package auth

import (
	"gorm.io/gorm"
	"memoirs/global"
	"memoirs/model/auth"
	"memoirs/model/vo"
)

type DictRepository struct{}

func (repo *DictRepository) QueryDict(page vo.ListQuery) (list []auth.SysDict, total int64, err error) {
	pageSize := page.PageSize
	offset := page.Offset()
	var dictList []auth.SysDict
	db := global.DB.Model(&auth.SysDict{})
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Find(&dictList).Error
	return dictList, total, err
}

func (repo *DictRepository) Insert(dict auth.SysDict) error {
	err := global.DB.Create(&dict).Error
	return err
}

func (repo *DictRepository) Update(dict auth.SysDict) error {
	db := global.DB.Model(auth.SysDict{}).Where("id", dict.ID)
	if dict.Name != "" {
		db = db.Update("name", dict.Name)
	}
	if dict.Status != nil {
		db = db.Update("status", dict.Status)
	}
	err := db.Error
	return err
}

func (repo *DictRepository) Delete(dictId uint) (err error) {
	var dict auth.SysDict
	err = global.DB.Where("id=?", dictId).Preload("SysDictItems").First(&dict).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return err
	}
	err = global.DB.Delete(auth.SysDict{}, dictId).Error
	if err != nil {
		return
	}
	err = global.DB.Where("dict_code", dict.Code).Delete(auth.SysDictItem{}).Error
	if err != nil {
		return
	}
	return err
}
