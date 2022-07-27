package api

import "memoirs/service"

type ApiGroup struct {
	UserApi
	RoleApi
	MenuApi
}

var ApiGroupApp = new(ApiGroup)

var (
	userService = service.ServiceGroupApp.UserService
	roleService = service.ServiceGroupApp.RoleService
	menuService = service.ServiceGroupApp.MenuService
)