package service

import (
	"memoirs/global"
	"memoirs/model"
)

type UserService struct{}

func (this *UserService) Login(username string) (*model.User, error) {
	var user model.User
	err := global.DB.Where("username=?", username).
		Or("phone=?", username).
		Or("email=?", username).
		First(&user).Error
	return &user, err
}

func (this *UserService) GetUserInfo(userId uint) (*model.User, error) {
	var user model.User
	err := global.DB.Preload("Roles").First(&user, userId).Error
	return &user, err
}

func (this *UserService) QueryUserList(pageSize, offset int) ([]model.User, error) {
	var userList []model.User
	err := global.DB.Preload("Roles").Limit(pageSize).Offset(offset).
		Find(&userList).Error
	return userList, err
}
