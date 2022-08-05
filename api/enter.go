package api

import (
	"memoirs/service"
)

type ApiGroup struct {
	UserApi
	RoleApi
	MenuApi
	SystemApi
}

var ApiGroupApp = new(ApiGroup)

var (
	authService   = service.ServiceGroupApp.AuthService
	userService   = service.ServiceGroupApp.UserService
	roleService   = service.ServiceGroupApp.RoleService
	menuService   = service.ServiceGroupApp.MenuService
	systemService = service.ServiceGroupApp.SystemService
)
