package auth

import "memoirs/service"

var (
	authService   = service.ServiceGroupApp.AuthService
	userService   = service.ServiceGroupApp.UserService
	roleService   = service.ServiceGroupApp.RoleService
	menuService   = service.ServiceGroupApp.MenuService
	systemService = service.ServiceGroupApp.SystemService
)
