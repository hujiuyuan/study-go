package main

import (
	"fmt"
	"time"
)

/*
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。

	考察点 ： go 关键字的使用、协程的并发执行。
*/
func main() {
	go printNum1()

	go printNum2()

	time.Sleep(10000 * time.Millisecond)
}

func printNum2() {
	func() {
		for i := range 10 {
			if i%2 == 0 {
				fmt.Println("打印偶数：", i)
			}
		}
	}()
}

func printNum1() {
	func() {
		for i := range 10 {
			if i%2 != 0 {
				fmt.Println("打印奇数：", i)
			}
		}
	}()
}
