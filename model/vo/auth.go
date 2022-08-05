package vo

type Empty struct{}

type LoginRequest struct {
	Account  string `json:"account"`  // 账号|手机号|邮箱
	Password string `json:"password"` // 登录密码
}

type LoginReply struct {
	Token string `json:"token"`
}

type RsaKeyReply struct {
	PublicKey string `json:"publicKey"` // 公钥
}

type UserInfoReply struct {
	Identity string      `json:"identity"` // 用户唯一标识
	NickName string      `json:"nickName"` // 昵称
	Gender   string      `json:"gender"`   // 性别
	Avatar   string      `json:"avatar"`   // 头像
	Phone    string      `json:"phone"`    // 电话号码
	Email    string      `json:"email"`    // 邮箱
	Roles    []RoleModel `json:"roles"`    // 角色列表
}

type RoleModel struct {
	RoleCode string `json:"roleCode"` // 角色编码
	RoleName string `json:"roleName"` // 角色名称
}

type UserRegister struct {
	Account  string `json:"account"`  // 账号
	Password string `json:"password"` // 密码
	NickName string `json:"nickName"` // 昵称
	Mode     int    `json:"mode"`     // 注册模式 1：账号；2电话；3邮箱
}

type RoleMenuRelation struct {
	RoleId  uint   `json:"roleId"`  // 角色ID
	MenuIds []uint `json:"menuIds"` // 菜单ID列表
}

type RoleRequest struct {
	RoleCode string `json:"roleCode"` // 角色编码
	RoleName string `json:"roleName"` // 角色名称
	ParentId uint   `json:"parentId"` // 父级ID
}

type RoleInfo struct {
	RoleId   uint   `json:"roleId"`   // 角色ID
	RoleCode string `json:"roleCode"` // 角色编码
	RoleName string `json:"roleName"` // 角色名称
	ParentId uint   `json:"parentId"` // 父级ID
}

type DeletedRole struct {
	RoleId uint `json:"roleId"` // 角色ID
}

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

type DeleteMenu struct {
	MenuIds []uint `json:"menuIds"` // 菜单ID列表
}

type QuerySetting struct {
	FontType string `json:"fontType"` // 字体类型
	FontSize uint   `json:"fontSize"` // 字体大小
}
