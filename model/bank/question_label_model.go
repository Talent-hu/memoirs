package bank

import "memoirs/model"

type QuestionLabel struct {
	model.BaseModel
	SubjectId uint   `gorm:"unique;comment:学科ID" json:"subjectId"`
	Label     string `gorm:"varchar(64);unique;comment:标签" json:"label"`
}

func (this *QuestionLabel) TableName() string {
	return "question_label"
}
