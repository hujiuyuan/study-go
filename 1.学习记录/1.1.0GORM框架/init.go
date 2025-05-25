package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

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
	// 链接MYSQL，获得DB类型实例，
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		//SkipDefaultTransaction: true, // true：关闭事务
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "f_", // 表明前缀
			//SingularTable: true, // 是否单数表名
			NoLowerCase: true, // 不要小写转换
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
	fmt.Println(DB)
	// 迁移 schema
	DB.AutoMigrate(&Student{})
	// 插入一条数据
	DB.Create(&Student{ID: 1, Name: "这是一个姓名", Age: 12})
	// 查询数据
	var stu Student
	DB.First(&stu, 1)
	// 打印学生信息
	stu.PrintInfo()
	// 更新数据
	// 更新多个字段
	DB.Model(&stu).Updates(Student{Name: "这是一个新名字", Age: 15})
	// 打印学生信息
	stu.PrintInfo()
	// 仅更新非零值字段
	DB.Model(&stu).Updates(map[string]interface{}{"Age": 200, "Name": "F42"})
	// 打印学生信息
	stu.PrintInfo()
}

type Student struct {
	ID   uint   `gorm:"primaryKey;comment:主键ID"`
	Name string `gorm:"type:varchar(200);comment:姓名"`
	Age  int    `gorm:"type:int(11);comment:年龄"`
}

func (stu *Student) PrintInfo() {
	fmt.Printf("学生信息 id: %d, name: %s, age: %d\n", stu.ID, stu.Name, stu.Age)
}
