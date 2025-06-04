package entity

import (
	"fmt"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title   string `gorm:"type:varchar(64);comment:标题"`
	Content string `gorm:"type:varchar(64);comment:文章内容"`
	UserId  string `gorm:"type:varchar(64);comment:关联用户id"`
}

func (post *Post) PrintInfo() {
	fmt.Printf("【文章信息】Id:%d,userId:%s,title:%s,content:%s\n", post.ID, post.UserId, post.Title, post.Content[:18]+"……")
}
