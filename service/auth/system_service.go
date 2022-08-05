package auth

import (
	"memoirs/model/vo"
	"memoirs/pkg/constant"
)

type SystemService struct{}

func (systemService *SystemService) QuerySetting() (map[string]*vo.QuerySetting, error) {
	// 先查询出侧边栏菜单设置
	menu, _ := menuMapper.QuerySetting(constant.IS_MENU)
	menuSetting := new(vo.QuerySetting)
	menuSetting.FontSize = menu.FontSize
	menuSetting.FontType = menu.FontType

	// 查询出按钮菜单设置
	btn, _ := menuMapper.QuerySetting(constant.IS_BTN)
	btnSetting := new(vo.QuerySetting)
	btnSetting.FontSize = btn.FontSize
	btnSetting.FontType = btn.FontType

	data := make(map[string]*vo.QuerySetting, 2)
	data["menuSetting"] = menuSetting
	data["btnSetting"] = btnSetting
	return data, nil
}
