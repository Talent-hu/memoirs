package vo

type RoleMenuRelation struct {
	RoleId  uint   `json:"roleId"`  // 角色ID
	MenuIds []uint `json:"menuIds"` // 菜单ID列表
}
