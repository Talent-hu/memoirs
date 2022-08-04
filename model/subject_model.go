package model

type SubjectCategory struct {
	BaseModel
	Code string `gorm:"varchar(127);comment:学科编码" json:"code"`
	Name string `gorm:"varchar(64);unique;comment:学科名称" json:"name"`
}

func (this *SubjectCategory) TableName() string {
	return "subject_category"
}
