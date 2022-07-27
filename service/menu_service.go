package service

import (
	"memoirs/global"
	"memoirs/model"
)

type MenuService struct {}


func (this *MenuService) BuildMenuTree(userId uint){


}

func (this *MenuService) AddMenu(menu *model.Menu) error{
	err := global.DB.Create(menu).Error
	return err
}