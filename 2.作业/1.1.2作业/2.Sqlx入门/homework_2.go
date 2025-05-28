package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

/*
题目2：实现类型安全映射,

	假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
	    要求 ：
	        定义一个 Book 结构体，包含与 books 表对应的字段。,
	        编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/
var mysqlClient *gorm.DB

func main() {
	// 迁移表结构
	//mysqlClient.AutoMigrate(&Book{})
	//var initBook = []Book{
	//	{Id: 1, Title: "标题1", Author: "作者1", Price: decimal.NewFromFloat(50.1)},
	//	{Id: 2, Title: "标题2", Author: "作者2", Price: decimal.NewFromFloat(11.2)},
	//	{Id: 3, Title: "标题3", Author: "作者3", Price: decimal.NewFromFloat(241.3)},
	//}
	//mysqlClient.Create(&initBook)

	var queryBook = []Book{}
	mysqlClient.Model(&queryBook).Where("price > ?", 50).Find(&queryBook)

	for _, book := range queryBook {
		book.printInfo()
	}
}

func init() {
	username := "root"
	password := "19971123"
	db_name := "canal_demo"
	port := 3306
	timeout := "10s"
	host := "localhost"

	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, db_name, timeout)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//SingularTable: true,
		},
	})

	if err != nil {
		panic("failed to connect database")
	}
	mysqlClient = db
}

type Book struct {
	Id     uint            `gorm:"primarykey; comment: 主键"`
	Title  string          `gorm:"varchar(255); comment: 标题"`
	Author string          `gorm:"varchar(255); comment: 作者"`
	Price  decimal.Decimal `gorm:"decimal(10,2); comment: 价格"`
}

func (b *Book) printInfo() {
	fmt.Printf("【书籍信息】ID:%d, 标题：%s, 作者：%s, 价格: %s\n", b.Id, b.Title, b.Author, b.Price.String())
}
