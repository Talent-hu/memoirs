package auth

import (
	"memoirs/model/auth"
	"memoirs/model/vo"
	"memoirs/utils"
)

type SystemService struct{}

func (srv *SystemService) QueryDict(page vo.ListQuery) ([]auth.SysDict, int64, error) {
	list, total, err := dictMapper.QueryDict(page)
	return list, total, err
}

func (srv *SystemService) AddDict(dict vo.Dict) error {
	var sysDict auth.SysDict
	err := utils.CopyProperties(&dict, &sysDict)
	if err != nil {
		return err
	}
	err = dictMapper.Insert(sysDict)
	return err
}

func (srv *SystemService) UpdateDict(dict vo.Dict) error {
	var sysDict auth.SysDict
	err := utils.CopyProperties(&dict, &sysDict)
	if err != nil {
		return err
	}
	err = dictMapper.Update(sysDict)
	return err
}

func (srv *SystemService) DeleteDict(dictId uint) error {
	err := dictMapper.Delete(dictId)
	return err
}

func (srv *SystemService) QueryDictItemList(pageItem vo.PageDictItem) (list []auth.SysDictItem, total int64, err error) {
	item := pageItem.DictItem
	query := pageItem.ListQuery
	var dictItem auth.SysDictItem
	err = utils.CopyProperties(&item, &dictItem)
	if err != nil {
		return
	}
	list, total, err = dictItemMapper.QueryAll(dictItem, query)
	return list, total, err
}

func (srv *SystemService) InsertDictItem(item vo.DictItem) (err error) {
	var dictItem auth.SysDictItem
	err = utils.CopyProperties(&item, &dictItem)
	if err != nil {
		return
	}
	err = dictItemMapper.Insert(dictItem)
	return err
}

func (srv *SystemService) UpdateDictItem(item vo.DictItem) (err error) {
	var dictItem auth.SysDictItem
	err = utils.CopyProperties(&item, &dictItem)
	if err != nil {
		return
	}
	err = dictItemMapper.Update(dictItem)
	return err
}

func (srv *SystemService) DeleteDictItem(itemId uint) (err error) {
	err = dictItemMapper.Delete(itemId)
	return err
}
