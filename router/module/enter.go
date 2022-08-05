package module

type RouterGroup struct {
	UserRouter
	MenuRouter
	RoleRouter
	SystemRouter
}

var RouterGroupApp = new(RouterGroup)
