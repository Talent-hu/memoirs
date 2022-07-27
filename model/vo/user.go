package vo

type Empty struct{}

type LoginRequest struct {
	Username string `json:"username"` // 账号|手机号|邮箱
	Password string `json:"password"` // 登录密码
}

type LoginResponse struct {
	Token string `json:"token"`
}

type GetRsaKeyResponse struct {
	PublicKey string `json:"publicKey"` // 公钥
}

type GetUserInfoResponse struct {
	Identity string `json:"identity"` // 用户唯一标识
	NickName string `json:"nickName"` // 昵称
	Gender   string `json:"gender"`   // 性别
	Avatar   string `json:"avatar"`   // 头像
	Phone    string `json:"phone"`    // 电话号码
	Email    string `json:"email"`    // 邮箱
	Roles   []RoleModel `json:"roles"` // 角色列表
}

type RoleModel struct {
	RoleCode string `json:"roleCode"` // 角色编码
	RoleName string `json:"roleName"` // 角色名称
}