package bank

import (
	"memoirs/global"
	"memoirs/model/bank"
)

type SelectRepository struct{}

func (repo *SelectRepository) QueryPage(pageSize, offset int) ([]bank.QuestionSelect, error) {
	var selectList []bank.QuestionSelect
	err := global.DB.Limit(pageSize).Offset(offset).Scan(&selectList).Error
	return selectList, err
}

func (repo *SelectRepository) Insert(questionSelect bank.QuestionSelect) error {
	err := global.DB.Create(&questionSelect).Error
	return err
}

/*func (repo *SelectRepository) Update(questionSelect model.QuestionSelect) error {
	db := global.DB.Model(&model.QuestionSelect{})
	db.Begin()

}*/
