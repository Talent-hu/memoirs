package service

import (
	"memoirs/service/auth"
	"memoirs/service/bank"
)

type ServiceGroup struct {
	auth.AuthService
	auth.UserService
	auth.RoleService
	auth.MenuService
	bank.SubjectService
}

var ServiceGroupApp = new(ServiceGroup)
