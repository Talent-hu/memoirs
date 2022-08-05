package bank

import (
	"memoirs/model"
	"memoirs/model/vo"
	"memoirs/utils"
)

type SubjectService struct{}

func (sub *SubjectService) QueryAll() []model.SubjectCategory {
	subList := subjectMapper.QueryAll()
	return subList
}

func (sub *SubjectService) QueryById(subId uint) model.SubjectCategory {
	subject := subjectMapper.QueryById(subId)
	return subject
}

func (sub *SubjectService) Insert(subReq vo.Subject) error {
	var subject model.SubjectCategory
	err := utils.CopyProperties(&subReq, &subject)
	if err != nil {
		return err
	}
	err = subjectMapper.Insert(subject)
	return err
}

func (sub *SubjectService) Update(subReq vo.Subject) error {
	var subject model.SubjectCategory
	err := utils.CopyProperties(&subReq, &subject)
	if err != nil {
		return err
	}
	subject.ID = subReq.ID
	err = subjectMapper.Update(subject)
	return err
}

func (sub *SubjectService) DeleteById(subId uint) error {
	err := subjectMapper.DeleteById(subId)
	return err
}
