package model

type User struct {
	BaseModel
	Identity string `gorm:"unique;comment:用户唯一标识" json:"identity"`
	Username string `gorm:"comment:用户登录名" json:"username"`
	Password string `gorm:"comment:用户登录密码" json:"-"`
	NickName string `gorm:"type:varchar(32);comment:用户昵称" json:"nickName"`
	Gender   string `gorm:"type:varchar(2);comment:性别" json:"gender"`
	Avatar   string `gorm:"comment:头像" json:"avatar"`
	Phone    string `gorm:"type:varchar(22);comment:用户手机号" json:"phone"`
	Email    string `gorm:"type:varchar(127);comment:电子邮箱" json:"email"`
	Roles    []Role `gorm:"many2many:user_role" json:"roles"`
}

func (this *User) TableName() string {
	return "user"
}
