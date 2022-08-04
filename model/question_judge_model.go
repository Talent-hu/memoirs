package model

type QuestionJudge struct {
	BaseQuestionModel
	Answer bool `gorm:"comment:答案" json:"answer"`
}

func (this *QuestionJudge) TableName() string {
	return "question_judge"
}
