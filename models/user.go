package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	NickName    string `json:"nickname" form:"nickname" default:"momo"`
	Gender      string `json:"gender" form:"gender" default:"武装直升机"`
	ContactInfo string `json:"contactInfo" form:"gender" default:"123456"`
	Username    string `json:"usernameVal" form:"usernameVal"`
	Password    string `json:"passwordVal" form:"passwordVal"`
	Signature   string `json:"signature" default:"我真帅吧"`
	Birthday    string `json:"birthday" default:"2006-01-02"`
}

// 表示配置操作数据库的表名称
// 表示把Student结构体默认操作的表改为students表
func (User) TableName() string {
	return "users"
}
