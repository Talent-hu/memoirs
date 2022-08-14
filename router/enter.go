package router

import (
	"memoirs/router/auth"
	"memoirs/router/bank"
	"memoirs/router/file"
)

type RouterGroup struct {
	auth.UserRouter
	auth.MenuRouter
	auth.RoleRouter
	auth.SystemRouter
	auth.AreaRouter
	auth.CasbinRouter
	file.FileRouter
	bank.SubjectRouter
}

var RouterGroupApp = new(RouterGroup)
