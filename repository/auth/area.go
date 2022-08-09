package auth

import (
	"memoirs/global"
	"memoirs/model/auth"
)

type AreaRepository struct{}

func (repo *AreaRepository) QueryAll() ([]auth.Area, error) {
	var areas []auth.Area
	err := global.DB.Model(auth.Area{}).Find(&areas).Error
	return areas, err
}
