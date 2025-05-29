package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

/*
题目1：模型定义,

    假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
        要求 ：
            使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。,
            编写Go代码，使用Gorm创建这些模型对应的数据库表。,
        ,
    ,

题目2：关联查询,

    基于上述博客系统的模型定义。
        要求 ：
            编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。,
            编写Go代码，使用Gorm查询评论数量最多的文章信息


题目3：钩子函数,

    继续使用博客系统的模型。
        要求 ：
            为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。,
            为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get github.com/shopspring/decimal
*/

var dbClient *gorm.DB

func main() {
	dbClient.AutoMigrate(&User{}, &Comment{}, &Post{})
	var users = []User{
		{Id: 1, Name: "用户a", UserName: "username1", Password: "这是一个加密密码", Salt: "这是一个加密盐"},
		{Id: 2, Name: "用户b", UserName: "username2", Password: "这是一个加密密码", Salt: "这是一个加密盐"},
		{Id: 3, Name: "用户c", UserName: "username3", Password: "这是一个加密密码", Salt: "这是一个加密盐"},
	}
	dbClient.Create(&users)
	var posts = []Post{
		{AuthorId: 1, Title: "这是一个标题", Data: "文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容"},
		{AuthorId: 2, Title: "这是一个标题", Data: "文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容"},
		{AuthorId: 3, Title: "这是一个标题", Data: "文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容"},
		{AuthorId: 1, Title: "这是一个标题", Data: "文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容"},
	}
	dbClient.Create(&posts)
	var comments = []Comment{
		{AuthorId: 1, PostId: 1, Remake: "这是一个评论这是一个评论这是一个评论这是一个评论这是一个评论"},
		{AuthorId: 1, PostId: 1, Remake: "这是一个评论这是一个评论这是一个评论这是一个评论这是一个评论"},
		{AuthorId: 1, PostId: 2, Remake: "这是一个评论这是一个评论这是一个评论这是一个评论这是一个评论"},
		{AuthorId: 1, PostId: 3, Remake: "这是一个评论这是一个评论这是一个评论这是一个评论这是一个评论"},
		{AuthorId: 2, PostId: 1, Remake: "这是一个评论这是一个评论这是一个评论这是一个评论这是一个评论"},
		{AuthorId: 2, PostId: 1, Remake: "这是一个评论这是一个评论这是一个评论这是一个评论这是一个评论"},
		{AuthorId: 3, PostId: 3, Remake: "这是一个评论这是一个评论这是一个评论这是一个评论这是一个评论"},
		{AuthorId: 3, PostId: 4, Remake: "这是一个评论这是一个评论这是一个评论这是一个评论这是一个评论"},
	}
	dbClient.Create(&comments)

	var user_1 User
	result := dbClient.Where("id = ?", 1).First(&user_1)

	if result.RowsAffected == 0 {
		fmt.Println("没查到账户")
	} else {
		user_1.printInfo()

		//posts := user_1.Posts
		//for _, post := range posts {
		//	post.printInfo()
		//}
		//
		//comments := Map(posts, func(post Post) []Comment {
		//	return post.Comments
		//})
		//for _, coms := range comments {
		//	for _, comment := range coms {
		//		comment.printInfo()
		//	}
		//
		//}

		var search_post []Post
		postResult := dbClient.Where("author_id = ?", user_1.Id).Find(&search_post)
		if postResult.RowsAffected == 0 {
			fmt.Printf("用户：%s, 下没有文章")
		} else {
			for _, post := range search_post {
				post.printInfo()
			}

			postIds := Map(search_post, func(post Post) uint { return post.Id })

			var comments []Comment
			dbClient.Where("post_id IN ?", postIds).Find(&comments)
			for _, comment := range comments {
				comment.printInfo()
			}
		}

	}
}

func init() {
	username := "root"
	password := "19971123"
	db_name := "canal_demo"
	host := "localhost"
	port := "3306"
	timeout := "10s"
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, db_name, timeout)

	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略 ErrRecordNotFound 错误
			Colorful:                  true,        // 启用彩色打印
		},
	)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic("链接数据库失败：error" + err.Error())
	}
	dbClient = db
}

type User struct {
	Id         uint      `gorm:"primarykey; comment:主键"`
	Name       string    `gorm:"type:varchar(255); comment:姓名"`
	UserName   string    `gorm:"type:varchar(255); comment:账号"`
	Password   string    `gorm:"type:varchar(255); comment:密码"`
	Salt       string    `gorm:"type:varchar(255); comment:密码加密salt"`
	CreateTime time.Time `gorm:"type:datetime; autoCreateTime; comment:创建时间"`
	UpdateTime time.Time `gorm:"type:datetime; autoUpdateTime; comment:更新时间"`
	PostCount  int       `gorm:"type:int; comment:文章数量"`
	Posts      []Post    `gorm:"foreignKey:AuthorId;references:Id;"`
}

func (user *User) printInfo() {
	fmt.Printf("【用户信息】Id:%d, Name:%s, PostCount:%d, UserName:%s, CreateTime:%s, UpdateTime:%s\n", user.Id, user.Name, user.PostCount, user.UserName, user.CreateTime.Format(time.DateTime), user.UpdateTime.Format(time.DateTime))
}

type Post struct {
	Id            uint      `gorm:"primarykey; comment:主键"`
	AuthorId      uint      `gorm:"comment:作者id"`
	Title         string    `gorm:"type:varchar(255); comment:标题"`
	Data          string    `gorm:"type:text; comment:文章"`
	CreateTime    time.Time `gorm:"type:datetime; autoCreateTime; comment:创建时间"`
	UpdateTime    time.Time `gorm:"type:datetime; autoUpdateTime; comment:更新时间"`
	CommentStatus string    `gorm:"type:varchar(50); comment:评论状态"`
	Comments      []Comment `gorm:"foreignkey:PostId"`
}

func (post *Post) AfterCreate(scope *gorm.DB) error {
	var user User
	scope.First(&user, post.AuthorId)
	user.PostCount++
	scope.Save(&user)
	return nil
}

func (post *Post) printInfo() {
	fmt.Printf("【用户信息】Id:%d, AuthorId:%d, CommentStatus:%s, Title:%s, Data:%s, CreateTime:%s, UpdateTime:%s\n", post.Id, post.AuthorId, post.CommentStatus, post.Title, post.Data[:12]+"……", post.CreateTime.Format(time.DateTime), post.UpdateTime.Format(time.DateTime))
}

type Comment struct {
	Id         uint      `gorm:"primarykey; comment:主键"`
	AuthorId   uint      `gorm:"comment:作者Id"`
	PostId     uint      `gorm:"comment:文章Id"`
	Remake     string    `gorm:"type:varchar(500); comment:评论"`
	CreateTime time.Time `gorm:"type:datetime; autoCreateTime; comment:创建时间"`
	UpdateTime time.Time `gorm:"type:datetime; autoUpdateTime; comment:更新时间"`
}

func (comment *Comment) AfterCreate(scope *gorm.DB) error {
	var post Post
	var countNum int64
	scope.Model(comment).Where("post_id = ?", comment.PostId).Find(&post)
	scope.Model(comment).Where("post_id = ?", comment.PostId).Count(&countNum)
	var commentStatus string
	if countNum == 0 {
		commentStatus = "无评论"
	} else {
		commentStatus = fmt.Sprintf("有%d条评论", countNum)
	}
	scope.Model(&post).Where("id = ?", comment.PostId).Updates(&Post{CommentStatus: commentStatus})
	return nil
}
func (comment *Comment) AfterDelete(scope *gorm.DB) error {
	var post Post
	var countNum int64
	tx := scope.Model(&comment).Where("post_id = ?", comment.PostId).Find(&post)
	tx.Count(&countNum)
	var commentStatus string
	if countNum == 0 {
		commentStatus = "无评论"
	} else {
		commentStatus = fmt.Sprintf("有%d条评论", countNum)
	}
	scope.Model(&post).Where("id = ?", comment.PostId).Updates(&Post{CommentStatus: commentStatus})
	return nil
}

func (comment *Comment) printInfo() {
	fmt.Printf("【用户信息】Id:%d, AuthorId:%d, PostId:%d, Remake:%s, CreateTime:%s, UpdateTime:%s\n", comment.Id, comment.AuthorId, comment.PostId, comment.Remake[:18]+"……", comment.CreateTime.Format(time.DateTime), comment.UpdateTime.Format(time.DateTime))
}

func Map[T, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}
