package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

/*
题目1：使用SQL扩展库进行查询,

	假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
	    要求 ：
	        编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。,
	        编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/
var DB *gorm.DB

func init() {
	username := "root"
	password := "19971123"
	host := "localhost"
	port := 3306
	Dbname := "canal_demo"
	timeout := "10s"
	// root:root@tcp(127.0.0.1:3306)/canal_demo?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		//SkipDefaultTransaction: true, // true：关闭事务
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "f_", // 表明前缀
			SingularTable: true, // 是否单数表名
			//NoLowerCase:   true, // 不要小写转换
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
func main() {
	DB.AutoMigrate(&Employees{})
	employees := []*Employees{
		{Id: 1, Name: "这是一个姓名1", Department: "部门1", Salary: 10000.0},
		{Id: 2, Name: "这是一个姓名2", Department: "技术部", Salary: 12000.0},
		{Id: 3, Name: "这是一个姓名3", Department: "技术部", Salary: 13000.0},
		{Id: 4, Name: "这是一个姓名4", Department: "部门1", Salary: 14000.0},
	}
	DB.Create(&employees)

	// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。,
	var emp_2 []Employees
	DB.Where("department = '技术部'").Find(&emp_2)

	fmt.Println(emp_2)

	// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	var emp_3 Employees
	DB.Order("salary DESC").First(&emp_3)
	fmt.Println(emp_3)
}

type Employees struct {
	Id         uint
	Name       string
	Department string
	Salary     float64
}
