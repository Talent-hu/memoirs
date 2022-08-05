package auth

import (
	"errors"
	"go.uber.org/zap"
	"memoirs/global"
	"memoirs/model/vo"
	"memoirs/utils"
)

type UserService struct{}

func (userService *UserService) QueryUserInfo(userId uint) (*vo.UserInfoReply, error) {
	userInfo, err := userMapper.QueryUserInfo(userId)
	if err != nil {
		global.Log.Error("系统错误，未查询到对应数据。", zap.Error(err))
		return nil, errors.New("系统错误，未查询到对应数据。")
	}
	resp := new(vo.UserInfoReply)
	_ = utils.CopyProperties(userInfo, resp)
	for _, role := range userInfo.Roles {
		roleModel := new(vo.RoleModel)
		roleModel.RoleName = role.RoleName
		roleModel.RoleCode = role.RoleCode
		resp.Roles = append(resp.Roles, *roleModel)
	}
	return resp, err
}

func (userService *UserService) QueryUserList(query vo.ListQuery) (*vo.PageQueryReply, error) {
	offset := query.Offset()
	userList, total, err := userMapper.QueryUserList(query.PageSize, offset)
	if err != nil {
		global.Log.Error("系统错误，未查询到对应数据。")
		return nil, errors.New("系统错误，未查询到对应数据。")
	}
	userInfoList := make([]vo.UserInfoReply, query.PageSize)
	for _, user := range userList {
		var userInfo vo.UserInfoReply
		err = utils.CopyProperties(user, &userInfo)
		for _, role := range user.Roles {
			roleModel := new(vo.RoleModel)
			roleModel.RoleName = role.RoleName
			roleModel.RoleCode = role.RoleCode
			userInfo.Roles = append(userInfo.Roles, *roleModel)
		}
		userInfoList = append(userInfoList, userInfo)
	}
	resp := new(vo.PageQueryReply)
	resp.Total = total
	resp.List = userInfoList
	return resp, err
}
