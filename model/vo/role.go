package vo

type RoleMenuRelation struct {
	RoleId  uint   `json:"roleId"`  // 角色ID
	MenuIds []uint `json:"menuIds"` // 菜单ID列表
}

type RoleRequest struct {
	RoleCode string `json:"roleCode"` // 角色编码
	RoleName string `json:"roleName"` // 角色名称
	ParentId uint   `json:"parentId"` // 父级ID
}

type RoleUpdate struct {
	RoleId   uint   `json:"roleId"`   // 角色ID
	RoleCode string `json:"roleCode"` // 角色编码
	RoleName string `json:"roleName"` // 角色名称
	ParentId uint   `json:"parentId"` // 父级ID
}

type DeletedRole struct {
	RoleId uint `json:"roleId"` // 角色ID
}
