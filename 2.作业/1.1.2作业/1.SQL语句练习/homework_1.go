package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

/*
基本CRUD操作,

    假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
        要求 ：
            编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。,
            编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。,
            编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。,
            编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/

var DB *gorm.DB

func main() {
	DB.AutoMigrate(&Student{})
	var stu Student
	// 插入一条新纪录， 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。,
	DB.Model(&stu).Create(&Student{1, "张三", 20, "三年级"})
	DB.Model(&stu).Create(&Student{2, "张三2", 1, "1年级"})

	// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。,
	var stus []Student
	DB.Where("age > 18").Find(&stus)
	fmt.Println(stus)

	// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。,
	DB.Where("name = ?", "张三").Updates(&Student{Grade: "四年级"})

	// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	DB.Where("age > 18").Find(&stus)
	fmt.Println(stus)

	// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	DB.Where("age > 15").Delete(&stu)

	DB.Find(&stus)
	fmt.Println(stus)
}

type Student struct {
	Id    uint   `gorm:"primarykey;comment: 主键"`
	Name  string `gorm:"type: varchar(50);comment: 主键"`
	Age   int    `gorm:"type: int(11);comment: 年龄"`
	Grade string `gorm:"type: varchar(20); comment:年纪"`
}

func init() {
	username := "root"
	password := "19971123"
	host := "localhost"
	port := 3306
	Dbname := "canal_demo"
	timeout := "10s"

	// root:root@tcp(127.0.0.1:3306)/canal_demo?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	// 链接MYSQL，获得DB类型实例，
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		//SkipDefaultTransaction: true, // true：关闭事务
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "f_", // 表明前缀
			//SingularTable: true, // 是否单数表名
			//NoLowerCase: true, // 不要小写转换
		},
	})

	if err != nil {
		panic("链接数据库失败：error" + err.Error())
	}
	// 链接成功
	fmt.Println("数据库链接成功")
	//fmt.Println(db)
	DB = db
}
