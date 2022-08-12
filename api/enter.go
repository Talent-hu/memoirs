package api

import (
	"memoirs/api/auth"
	"memoirs/api/bank"
)

type ApiGroup struct {
	auth.UserApi
	auth.RoleApi
	auth.MenuApi
	auth.SystemApi
	auth.AreaApi
	auth.CasbinApi
	bank.SubjectApi
}

var ApiGroupApp = new(ApiGroup)
