package auth

import (
	"memoirs/global"
	"memoirs/model/auth"
	"memoirs/model/vo"
)

type DictItemRepository struct{}

func (repo *DictItemRepository) QueryAll(dict auth.SysDictItem, page vo.ListQuery) (list []auth.SysDictItem, total int64, err error) {
	pageSize := page.PageSize
	offset := page.Offset()
	var dictItemList []auth.SysDictItem
	db := global.DB.Model(&auth.SysDictItem{})
	if dict.DictCode != "" {
		db = db.Where("dict_code=?", dict.DictCode)
	}
	if dict.Name != "" {
		db = db.Where("`name` LIKE  ?", "%"+dict.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Find(&dictItemList).Error
	return dictItemList, total, err
}

func (repo *DictItemRepository) Insert(dict auth.SysDictItem) error {
	err := global.DB.Create(&dict).Error
	return err
}

func (repo *DictItemRepository) Update(dict auth.SysDictItem) error {
	db := global.DB.Model(&auth.SysDictItem{}).Where("id = ?", dict.ID)
	if dict.Name != "" {
		db = db.Update("name", dict.Name)
	}
	if dict.Value != "" {
		db = db.Update("value", dict.Value)
	}
	if dict.Status != nil {
		db = db.Update("status", dict.Status)
	}
	err := db.Error
	return err
}

func (repo *DictItemRepository) Delete(dictId uint) error {
	var dictItem auth.SysDictItem
	err := global.DB.Delete(&dictItem, dictId).Error
	return err
}
