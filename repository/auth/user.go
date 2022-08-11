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

func (repo *UserRepository) Login(username string) (*auth.User, error) {
	var user auth.User
	err := global.DB.Where("username=?", username).
		Or("phone=?", username).
		Or("email=?", username).
		First(&user).Error
	return &user, err
}

func (repo *UserRepository) QueryUserInfo(userId uint) (*auth.User, error) {
	var user auth.User
	err := global.DB.Preload("Roles").First(&user, userId).Error
	return &user, err
}

func (repo *UserRepository) QueryUserList(pageSize, offset int) ([]auth.User, int64, error) {
	var userList []auth.User
	var count int64
	err := global.DB.Preload("Roles").Limit(pageSize).Offset(offset).
		Find(&userList).Error
	global.DB.Model(&auth.User{}).Count(&count)
	return userList, count, err
}

func (repo *UserRepository) AddUserAndRole(userId, roleId uint) error {
	userRole := new(auth.UserRole)
	userRole.RoleId = roleId
	userRole.UserId = userId
	err := global.DB.Create(userRole).Error
	return err
}

func (repo *UserRepository) DelUserAndRole(userId, roleId uint) error {
	err := global.DB.Unscoped().Where("user_id=? AND role_id=?", userId, roleId).
		Delete(&auth.UserRole{}).Error
	return err
}
