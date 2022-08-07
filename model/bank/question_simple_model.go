package bank

import "memoirs/model"

type QuestionSimple struct {
	model.BaseQuestionModel
	Answer string `gorm:"longtext(5000);comment:问答题" json:"answer"`
}
