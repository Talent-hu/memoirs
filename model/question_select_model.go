package model

type QuestionSelect struct {
	BaseQuestionModel
	OptionA string `gorm:"varchar(255);comment:选择A" json:"optionA"`
	OptionB string `gorm:"varchar(255);comment:选择B" json:"optionB"`
	OptionC string `gorm:"varchar(255);comment:选择C" json:"optionC"`
	OptionD string `gorm:"varchar(255);comment:选择D" json:"optionD"`
	OptionE string `gorm:"varchar(255);comment:选择E" json:"optionE"`
	OptionF string `gorm:"varchar(255);comment:选择F" json:"optionF"`
	Answer  string `gorm:"varchar(6);comment:答案" json:"answer"`
}

func (this *QuestionSelect) TableName() string {
	return "question_select"
}
