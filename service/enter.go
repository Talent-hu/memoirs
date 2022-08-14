package service

import (
	"memoirs/service/auth"
	"memoirs/service/bank"
	"memoirs/service/file"
)

type ServiceGroup struct {
	auth.AuthService
	auth.UserService
	auth.RoleService
	auth.MenuService
	auth.SystemService
	auth.AreaService
	auth.CasbinService
	file.FileService
	bank.SubjectService
}

var ServiceGroupApp = new(ServiceGroup)
