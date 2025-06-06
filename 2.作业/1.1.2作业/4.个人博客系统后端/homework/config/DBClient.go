package config

import (
	"blog-backend/entity"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var DB *gorm.DB

func init() {
	username := "root"
	password := "19971123"
	host := "localhost"
	port := 3306
	Dbname := "canal_demo"
	timeout := "10s"

	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略 ErrRecordNotFound 错误
			Colorful:                  true,        // 启用彩色打印
		},
	)

	// root:root@tcp(127.0.0.1:3306)/canal_demo?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	// 链接MYSQL，获得DB类型实例，
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		//SkipDefaultTransaction: true, // true：关闭事务
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "block_", // 表明前缀
			//SingularTable: true, // 是否单数表名
			NoLowerCase: true, // 不要小写转换
		},
		Logger: newLogger,
	})

	if err != nil {
		panic("链接数据库失败：error" + err.Error())
	}
	// 链接成功
	fmt.Println("数据库链接成功")
	//fmt.Println(db)
	DB = db

	db.AutoMigrate(&entity.User{}, &entity.Comment{}, &entity.Post{})
}
