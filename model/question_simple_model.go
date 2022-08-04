package model

type QuestionSimple struct {
	BaseQuestionModel
	Answer string `gorm:"longtext(5000);comment:问答题" json:"answer"`
}
