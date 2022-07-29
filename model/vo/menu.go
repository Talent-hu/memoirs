package vo

type AddMenuRequest struct {
	Path      string `json:"path"`      // 路由path
	Name      string `json:"name"`      // 路由名称
	Component string `json:"component"` // 对应前端文件路径
	Hidden    bool   `json:"hidden"`    // 是否隐藏列表
	Sort      int    `json:"sort"`      // 排序标记
	Title     string `json:"title"`     // 菜单名称
	Icon      string `json:"icon"`      // 菜单图标
	ParentId  uint   `json:"parentId"`  // 父级菜单ID
	FontType  string `json:"fontType"`  // 字体类型
	FontSize  uint   `json:"fontSize"`  // 字体大小
	HasBtn    bool   `json:"hasBtn"`    // 是否是按钮
}

type MenuTree struct {
	Path      string     `json:"path"`      // 路由path
	Name      string     `json:"name"`      // 路由名称
	Component string     `json:"component"` // 对应前端文件路径
	Hidden    bool       `json:"hidden"`    // 是否隐藏列表
	Sort      int        `json:"sort"`      // 排序标记
	Title     string     `json:"title"`     // 菜单名称
	Icon      string     `json:"icon"`      // 菜单图标
	ParentId  uint       `json:"parentId"`  // 父级菜单ID
	FontType  string     `json:"fontType"`  // 字体类型
	FontSize  uint       `json:"fontSize"`  // 字体大小
	HasBtn    bool       `json:"hasBtn"`    // 是否是按钮
	ID        uint       `json:"menuId"`
	Children  []MenuTree `json:"children"`
}

type RemoveMenu struct {
	MenuIds []uint `json:"menuIds"` // 菜单ID列表
}
