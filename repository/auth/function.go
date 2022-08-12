package auth

import (
	"memoirs/global"
	"memoirs/model/auth"
)

type FunctionRepository struct{}

func (repo *FunctionRepository) QueryAll() ([]auth.SysFunction, error) {
	var funcList []auth.SysFunction
	err := global.DB.Model(&auth.SysFunction{}).Scan(&funcList).Error
	return funcList, err
}
