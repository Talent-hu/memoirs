package auth

import (
	"memoirs/model"
	"memoirs/model/vo"
	"memoirs/utils"
)

type RoleService struct{}

func (roleService *RoleService) QueryAll(queryPage vo.ListQuery) *vo.PageQueryReply {
	pageSize := queryPage.PageSize
	offset := queryPage.Offset()
	roleList, total := roleMapper.QueryList(pageSize, offset)
	var roleInfoList []vo.RoleInfo
	for _, role := range roleList {
		var roleInfo vo.RoleInfo
		_ = utils.CopyProperties(&role, &roleInfo)
		roleInfo.RoleId = role.ID
		roleInfoList = append(roleInfoList, roleInfo)
	}
	resp := new(vo.PageQueryReply)
	resp.Total = total
	resp.List = roleInfoList
	return resp
}

func (roleService *RoleService) QueryUserRole(userId uint) ([]vo.RoleInfo, error) {
	roleList, err := roleMapper.QueryUserRoleList(userId)
	if err != nil {
		return nil, err
	}
	var roleInfoList []vo.RoleInfo
	for _, role := range roleList {
		var roleInfo vo.RoleInfo
		err = utils.CopyProperties(&role, &roleInfo)
		roleInfo.RoleId = role.ID
		roleInfoList = append(roleInfoList, roleInfo)
	}
	return roleInfoList, err
}

func (roleService *RoleService) AddRole(roleReq vo.RoleRequest) error {
	var role model.Role
	err := utils.CopyProperties(&roleReq, &role)
	if err != nil {
		return err
	}
	err = roleMapper.AddRole(role)
	return err
}

func (roleService *RoleService) AddRoleAndMenu(relation vo.RoleMenuRelation) error {
	var roleMenus []model.RoleMenu
	for _, menuId := range relation.MenuIds {
		roleMenu := new(model.RoleMenu)
		roleMenu.RoleId = relation.RoleId
		roleMenu.MenuId = menuId
		roleMenus = append(roleMenus, *roleMenu)
	}
	err := roleMapper.AddRoleAndMenu(roleMenus)
	return err
}

func (roleService *RoleService) UpdateRole(roleReq vo.RoleInfo) error {
	var role model.Role
	_ = utils.CopyProperties(&roleReq, &role)
	role.ID = roleReq.RoleId
	err := roleMapper.UpdateRole(role)
	return err
}

func (roleService *RoleService) DeleteRole(roleId uint) error {
	err := roleMapper.DeleteRole(roleId)
	return err
}

func (roleService *RoleService) DeleteRoleAndMenu(relation vo.RoleMenuRelation) error {
	err := roleMapper.DelRelation(relation.RoleId, relation.MenuIds)
	return err
}
