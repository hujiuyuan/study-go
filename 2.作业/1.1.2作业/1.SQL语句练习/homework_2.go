package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
题目2：事务语句,

	假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
	    要求 ：
	        编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/

var Db *gorm.DB

func main() {
	Db.AutoMigrate(&Account{}, &Transaction{})

	amount, _ := decimal.NewFromString("1000.50")
	Db.Create(&Account{Id: 1, Balance: amount})
	Db.Create(&Account{Id: 2, Balance: amount})

	tx := Db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	var A, B Account
	tx.Model(&A).Where("id = ?", 1)
	tx.Model(&B).Where("id = ?", 2)
	money := decimal.NewFromInt(100)
	if decimal.Max(A.Balance, money) == money {
		tx.Create(&Transaction{Amount: amount, ToAccountId: B.Id, FromAccountId: A.Id})
		tx.Model(&A).Updates(Account{Balance: amount.Sub(money)})
		tx.Model(&B).Updates(Account{Balance: amount.Add(money)})
	}

	if err := tx.Commit().Error; err != nil {
		fmt.Println("Failed to commit transaction:", err)
		return
	}
}

func init() {
	Dbname := "canal_demo"
	username := "root"
	password := "19971123"
	host := "127.0.0.1"
	port := "3306"
	timeout := "10s"

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic("链接数据库失败：error" + err.Error())
	}

	fmt.Println("数据库链接成功")
	Db = db
}

type Transaction struct {
	Id            uint64          `gorm:"primaryKey comment:主键;"`
	FromAccountId uint64          `gorm:"type: bigint; comment:关联转出账户ID;"`
	ToAccountId   uint64          `gorm:"type: bigint; comment:关联转入账户ID;"`
	Amount        decimal.Decimal `gorm:"type: decimal(10,2);comment:账户余额;"`
}

type Account struct {
	Id      uint64          `gorm:"primaryKey comment:主键;"`
	Balance decimal.Decimal `gorm:"type: decimal(10,2);comment:账户余额;"`
}
