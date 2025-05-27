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
	// 自动迁移数据库结构
	Db.AutoMigrate(&Account{}, &Transaction{})

	//// 初始化账户
	//amount, _ := decimal.NewFromString("1000.50")
	//Db.Create(&Account{Id: 1, Balance: amount})
	//Db.Create(&Account{Id: 2, Balance: amount})

	// 开始事务
	tx := Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			fmt.Println("Transaction rolled back due to panic:", r)
		}
	}()

	// 检查是否存在A、B两个账户
	var A, B Account
	if err := tx.Where("id = ?", 1).First(&A).Error; err != nil {
		tx.Rollback()
		fmt.Println("Transaction rolled back due to error:", err)
	}

	if err := tx.Where("id = ?", 2).First(&B).Error; err != nil {
		tx.Rollback()
		fmt.Println("Transaction rolled back due to error:", err)
	}

	A.printInfo()
	B.printInfo()

	// 定义转账金额
	money := decimal.NewFromInt(500)
	//fmt.Println("判断A的余额")
	if A.Balance.LessThan(money) {
		tx.Rollback()
		fmt.Println("A的余额不足不能转账")
	}

	//fmt.Println("A的账户扣款")
	if err := tx.Model(&A).Where("id = ?", A.Id).Update("balance", A.Balance.Sub(money)).Error; err != nil {
		tx.Rollback()
		fmt.Println("A的账户扣款失败", err)
	}
	//fmt.Println("B的账户收款")
	if err := tx.Model(&B).Where("id = ?", B.Id).Update("balance", B.Balance.Add(money)).Error; err != nil {
		tx.Rollback()
		fmt.Println("B的账户收款失败", err)
	}

	//fmt.Println("记录转账记录")
	if err := tx.Create(&Transaction{FromAccountId: A.Id, ToAccountId: B.Id, Amount: money}).Error; err != nil {
		tx.Rollback()
		fmt.Println("记录转账记录失败", err)
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		fmt.Println("事务提交失败", err)
	}

	var amounts []Account
	Db.Find(&amounts)

	for _, account := range amounts {
		account.printInfo()
	}

	var transactions []Transaction
	Db.Find(&transactions)
	for _, t := range transactions {
		t.printInfo()
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
	Id            uint64          `gorm:"primaryKey; comment:主键;"`
	FromAccountId uint64          `gorm:"type: bigint; comment:关联转出账户ID;"`
	ToAccountId   uint64          `gorm:"type: bigint; comment:关联转入账户ID;"`
	Amount        decimal.Decimal `gorm:"type: decimal(10,2);comment:账户余额;"`
}

type Account struct {
	Id      uint64          `gorm:"primaryKey; comment:主键;"`
	Balance decimal.Decimal `gorm:"type: decimal(10,2);comment:账户余额;"`
}

func (t *Transaction) printInfo() {
	fmt.Printf("Transaction info: Id:%d; FromAccountId: %d; ToAccountId:%d;Amount:%s;", t.Id, t.FromAccountId, t.ToAccountId, t.Amount)
}

func (a *Account) printInfo() {
	fmt.Printf("Account info: Id:%d Balance:%s \n", a.Id, a.Balance)
}
