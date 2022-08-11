package router

import (
	"memoirs/router/auth"
	"memoirs/router/bank"
)

type RouterGroup struct {
	auth.UserRouter
	auth.MenuRouter
	auth.RoleRouter
	auth.SystemRouter
	auth.AreaRouter
	auth.CasbinRouter
	bank.SubjectRouter
}

var RouterGroupApp = new(RouterGroup)
