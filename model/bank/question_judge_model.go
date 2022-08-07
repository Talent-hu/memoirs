package bank

import "memoirs/model"

type QuestionJudge struct {
	model.BaseQuestionModel
	Answer bool `gorm:"comment:答案" json:"answer"`
}

func (this *QuestionJudge) TableName() string {
	return "question_judge"
}
