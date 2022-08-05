package bank

import (
	"memoirs/global"
	"memoirs/model"
)

type SubjectRepository struct{}

func (sub *SubjectRepository) QueryAll() []model.SubjectCategory {
	var subList []model.SubjectCategory
	global.DB.Find(&subList)
	return subList
}

func (sub *SubjectRepository) QueryById(subId uint) model.SubjectCategory {
	var subject model.SubjectCategory
	global.DB.Find(&subject)
	return subject
}

func (sub *SubjectRepository) Insert(subject model.SubjectCategory) error {
	err := global.DB.Create(&subject).Error
	return err
}

func (sub *SubjectRepository) Update(subject model.SubjectCategory) error {
	db := global.DB.Model(&model.SubjectCategory{})
	db.Begin()
	db.Where("id=?", subject.ID)
	if subject.Code != "" {
		db.Update("code", subject.Code)
	}
	if subject.Name != "" {
		db.Update("name", subject.Name)
	}
	err := db.Error
	db.Commit()
	if err != nil {
		db.Callback()
	}
	return err
}

func (sub *SubjectRepository) DeleteById(subId uint) error {
	err := global.DB.Delete(&model.SubjectCategory{}, subId).Error
	return err
}
