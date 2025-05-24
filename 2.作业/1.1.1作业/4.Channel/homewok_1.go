package main

import (
	"fmt"
	"sync"
	"time"
)

/*
编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。

	考察点 ：通道的基本使用、协程间通信。
*/
func main() {
	ch := make(chan int)
	// 创建一个 WaitGroup 用于同步
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 1; i <= 10; i++ {
			fmt.Printf("%s发送数据：%d\n", time.Now().Format(time.DateTime), i)
			ch <- i
		}

		// 发送完之后关闭通道
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := range ch {
			fmt.Printf("%s接收到数据：%d\n", time.Now().Format(time.DateTime), i)
		}
	}()

	wg.Wait()

}
