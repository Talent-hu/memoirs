package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
)

type BaseModel struct {
	ID        uint           `gorm:"primary key;comment:主键ID" json:"id"`
	CreatedAt MyTime         `gorm:"type:datetime(3);comment:创建时间" json:"createdAt"`
	UpdatedAt MyTime         `gorm:"type:datetime(3);comment:更新时间" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间" json:"-"`
}

type BaseQuestionModel struct {
	BaseModel
	LabelID   string `gorm:"comment:标签ID" json:"labelId"`
	Type      uint   `gorm:"int;unique;comment:试题类别(1:选择题,2:多选题,3:判断题,4:填空题,5:问答题)" json:"type"`
	Name      string `gorm:"varchar(64);comment:题型名称" json:"name"`
	Difficult string `gorm:"int;comment:难度(1:简单、2:中等、3:难)" json:"difficult"`
	Title     string `gorm:"longtext(1000);comment:试题题目" json:"title"`
	Analysis  string `gorm:"longtext(500);comment:解析" json:"analysis"`
}

type BaseTreeModel struct {
	ID        string         `gorm:"type:varchar(16);primary key;comment:主键ID" json:"id"`
	ParentId  string         `gorm:"type:varchar(16);comment:父级编码" json:"parentId"`
	CreatedAt MyTime         `gorm:"type:datetime(3);comment:创建时间" json:"createdAt"`
	UpdatedAt MyTime         `gorm:"type:datetime(3);comment:更新时间" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间" json:"-"`
}

type MyTime time.Time

func (t *MyTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	str := string(data)
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = MyTime(t1)
	return err
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t MyTime) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Now()
	return tTime.Format("2006-01-02 15:04:05"), nil
}

func (t *MyTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = MyTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *MyTime) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}
