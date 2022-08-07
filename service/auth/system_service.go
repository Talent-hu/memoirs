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
