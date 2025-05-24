package main

import (
	"fmt"
	"sync"
	"time"
)

/*
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。

	考察点 ：通道的缓冲机制。
*/
func main() {

	ch := make(chan int, 100)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Printf("%s发送数据：%d\n", time.Now().Format(time.DateTime), i)
			ch <- i
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for i := range ch {
			fmt.Printf("%s接收数据：%d\n", time.Now().Format(time.DateTime), i)
		}
	}()

	wg.Wait()
}
