package service

type ServiceGroup struct {
	UserService
	RoleService
	MenuService
}

var ServiceGroupApp = new(ServiceGroup)
