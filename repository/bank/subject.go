package bank

import (
	"memoirs/global"
	"memoirs/model/bank"
)

type SubjectRepository struct{}

func (sub *SubjectRepository) QueryAll() []bank.SubjectCategory {
	var subList []bank.SubjectCategory
	global.DB.Find(&subList)
	return subList
}

func (sub *SubjectRepository) QueryById(subId uint) bank.SubjectCategory {
	var subject bank.SubjectCategory
	global.DB.Find(&subject)
	return subject
}

func (sub *SubjectRepository) Insert(subject bank.SubjectCategory) error {
	err := global.DB.Create(&subject).Error
	return err
}

func (sub *SubjectRepository) Update(subject bank.SubjectCategory) error {
	db := global.DB.Model(&bank.SubjectCategory{})
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
	err := global.DB.Delete(&bank.SubjectCategory{}, subId).Error
	return err
}
