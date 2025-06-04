package entity

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(64);comment:账号"`
	Password string `gorm:"type:varchar(64);comment:密码"`
	Salt     string `gorm:"type:varchar(64);comment:密码(盐)"`
	Email    string `gorm:"type:varchar(64);comment:邮箱"`
	Posts    []Post `gorm:"foreignkey:UserId;references:ID;"`
}

func (user *User) PrintInfo() {
	fmt.Printf("【用户信息】Id:%d,Username:%s,Email:%s\n", user.ID, user.Username, user.Email)
}
