package module

type RouterGroup struct {
	UserRouter
	MenuRouter
}

var RouterGroupApp = new(RouterGroup)
