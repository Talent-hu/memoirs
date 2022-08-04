package model

type QuestionFillBack struct {
	BaseQuestionModel
	Answer string `gorm:"varchar(255);comment:答案" json:"answer"`
}

func (this *QuestionFillBack) TableName() string {
	return "question_fill_back"
}
