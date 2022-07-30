package service

import "memoirs/global"

type ServiceGroup struct {
	UserService
	RoleService
	MenuService
}

var ServiceGroupApp = new(ServiceGroup)

var db = global.DB
