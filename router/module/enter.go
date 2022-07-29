package module

type RouterGroup struct {
	UserRouter
	MenuRouter
	RoleRouter
}

var RouterGroupApp = new(RouterGroup)
