package api

import (
	"memoirs/api/auth"
	"memoirs/api/bank"
	"memoirs/api/file"
)

type ApiGroup struct {
	auth.UserApi
	auth.RoleApi
	auth.MenuApi
	auth.SystemApi
	auth.AreaApi
	auth.CasbinApi
	file.ImageApi
	bank.SubjectApi
}

var ApiGroupApp = new(ApiGroup)
