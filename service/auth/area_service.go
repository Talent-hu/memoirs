package auth

import (
	"memoirs/global"
	"memoirs/model/auth"
	"memoirs/pkg/constant"
)

type AreaService struct{}

func (srv *AreaService) QueryList() (tree []auth.Area, err error) {
	global.Log.Info("获取所有的地区树")
	areas, err := areaMapper.QueryAll()
	if err != nil {
		return
	}
	treeList := ListToTree(areas, constant.SUPER_AREA_PARENT_ID)
	return treeList, err
}

func ListToTree(list []auth.Area, parentId string) []auth.Area {
	var nodeList []auth.Area
	for _, item := range list {
		if item.ParentId == parentId {
			item.Children = ListToTree(list, item.ID)
			nodeList = append(nodeList, item)
		}
	}
	return nodeList
}
