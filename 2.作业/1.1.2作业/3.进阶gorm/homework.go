package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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

go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get github.com/shopspring/decimal
*/

var dbClient *gorm.DB

func main() {

}

func init() {
	username := "root"
	password := "19971123"
	db_name := "canal_demo"
	host := "localhost"
	port := "3306"
	timeout := "10s"
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, db_name, timeout)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
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
	CreateTime time.Time `gorm:"type:datetime; comment:创建时间"`
	UpdateTime time.Time `gorm:"type:datetime; comment:更新时间"`
}

func (user *User) printInfo() {
	fmt.Printf("【用户信息】Id:%d, Name:%s, UserName:%s, CreateTime:%s, UpdateTime:%s\n", user.Id, user.Name, user.UserName, user.CreateTime, user.UpdateTime)
}
