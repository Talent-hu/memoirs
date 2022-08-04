package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint           `gorm:"primary key;comment:主键ID" json:"id"`
	CreatedAt time.Time      `gorm:"comment:创建时间" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"comment:更新时间" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间" json:"-"`
}

type BaseQuestionModel struct {
	BaseModel
	LabelID   uint   `gorm:"comment:标签ID" json:"labelId"`
	Type      uint   `gorm:"int;unique;comment:试题类别(1:选择题,2:多选题,3:判断题,4:填空题,5:问答题)" json:"type"`
	Name      string `gorm:"varchar(64);comment:题型名称" json:"name"`
	Difficult string `gorm:"int;comment:难度(1:简单、2:中等、3:难)" json:"difficult"`
	Title     string `gorm:"longtext(1000);comment:试题题目" json:"title"`
	Analysis  string `gorm:"longtext(500);comment:解析" json:"analysis"`
}
