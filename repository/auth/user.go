package auth

import (
	"memoirs/global"
	"memoirs/model/auth"
)

type UserRepoInterface interface {
	Login(username string) (*auth.User, error)
	QueryUserInfo(userId uint) (*auth.User, error)
	QueryUserList(pageSize, offset int) ([]auth.User, int64, error)
}

type UserRepository struct{}

func (this *UserRepository) Login(username string) (*auth.User, error) {
	var user auth.User
	err := global.DB.Where("username=?", username).
		Or("phone=?", username).
		Or("email=?", username).
		First(&user).Error
	return &user, err
}

func (this *UserRepository) QueryUserInfo(userId uint) (*auth.User, error) {
	var user auth.User
	err := global.DB.Preload("Roles").First(&user, userId).Error
	return &user, err
}

func (this *UserRepository) QueryUserList(pageSize, offset int) ([]auth.User, int64, error) {
	var userList []auth.User
	var count int64
	err := global.DB.Preload("Roles").Limit(pageSize).Offset(offset).
		Find(&userList).Error
	global.DB.Model(&auth.User{}).Count(&count)
	return userList, count, err
}
