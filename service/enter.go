package service

import "memoirs/service/auth"

type ServiceGroup struct {
	auth.AuthService
	auth.UserService
	auth.RoleService
	auth.MenuService
	auth.SystemService
}

var ServiceGroupApp = new(ServiceGroup)
