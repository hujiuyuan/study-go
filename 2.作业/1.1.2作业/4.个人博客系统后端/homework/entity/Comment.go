package entity

import (
	"fmt"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Id      uint   `gorm:"primarykey;comment:主键"`
	Content string `gorm:"type:varchar(500);comment:评论信息"`
	UserId  uint   `gorm:"comment:关联用户id"`
	PostId  uint   `gorm:"comment:关联用户id"`
}

func (comment *Comment) PrintInfo() {
	fmt.Printf("【评论信息】Id:%d,userId:%s,postId:%s,content:%s\n", comment.Id, comment.UserId, comment.PostId, comment.Content[:18]+"……")
}
