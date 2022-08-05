package repository

import (
	"memoirs/global"
	"memoirs/model"
)

type UserRepoInterface interface {
	Login(username string) (*model.User, error)
	QueryUserInfo(userId uint) (*model.User, error)
	QueryUserList(pageSize, offset int) ([]model.User, int64, error)
}

type UserRepository struct{}

func (this *UserRepository) Login(username string) (*model.User, error) {
	var user model.User
	err := global.DB.Where("username=?", username).
		Or("phone=?", username).
		Or("email=?", username).
		First(&user).Error
	return &user, err
}

func (this *UserRepository) QueryUserInfo(userId uint) (*model.User, error) {
	var user model.User
	err := global.DB.Preload("Roles").First(&user, userId).Error
	return &user, err
}

func (this *UserRepository) QueryUserList(pageSize, offset int) ([]model.User, int64, error) {
	var userList []model.User
	var count int64
	err := global.DB.Preload("Roles").Limit(pageSize).Offset(offset).
		Find(&userList).Error
	global.DB.Model(&model.User{}).Count(&count)
	return userList, count, err
}
