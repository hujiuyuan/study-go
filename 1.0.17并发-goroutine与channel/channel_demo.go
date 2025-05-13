package main

import (
	"fmt"
	"time"
)

func main() {
	channel_demo()
}

/*
channel

channel 是 Go 中定义的一种类型，专门用来在多个 goroutine 之间通信的线程安全的数据结构。

可以在一个 goroutine 中向一个 channel 中发送数据，从另外一个 goroutine 中接收数据。

channel 类似队列，满足先进先出原则。

定义方式：
// 仅声明
var <channel_name> chan <type_name>

// 初始化
<channel_name> := make(chan <type_name>)

// 初始化有缓冲的channel
<channel_name> := make(chan <type_name>, 3)

channel 的三种操作：发送数据，接收数据，以及关闭通道。

声明方式：
// 发送数据
channel_name <- variable_name_or_value

// 接收数据
value_name, ok_flag := <- channel_name
value_name := <- channel_name

// 关闭channel
close(channel_name)

channel 还有两个变种，可以把 channel 作为参数传递时，限制 channel 在函数或方法中能够执行的操作。

声明方式：
//仅发送数据
func <method_name>(<channel_name> chan <- <type>)

//仅接收数据
func <method_name>(<channel_name> <-chan <type>)

锁与 channel

在 Go 中，当需要 goroutine 之间协作地方，更常见的方式是使用 channel，而不是 sync 包中的 Mutex 或 RWMutex 的互斥锁。但其实它们各有侧重。

大部分时候，流程是根据数据驱动的，channel 会被使用得更频繁。

channel 擅长的是数据流动的场景：

	传递数据的所有权，即把某个数据发送给其他协程。
	分发任务，每个任务都是一个数据。
	交流异步结果，结果是一个数据。

而锁使用的场景更偏向同一时间只给一个协程访问数据的权限：

	访问缓存
	管理状态
*/
func channel_demo() {
	// 创建一个带缓冲的channel
	ch := make(chan int, 3)

	// 启动发送goroutine
	go sendOnly(ch)

	// 启动接收goroutine
	go receiveOnly(ch)

	// 使用select进行多路复用
	timeout := time.After(2 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("Channel已关闭")
				return
			}
			fmt.Printf("主goroutine接收到: %d\n", v)
		case <-timeout:
			fmt.Println("操作超时")
			return
		default:
			fmt.Println("没有数据，等待中...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// 只接收channel的函数
func receiveOnly(ch <-chan int) {
	for v := range ch {
		fmt.Printf("接收到: %d\n", v)
	}
}

// 只发送channel的函数
func sendOnly(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("发送: %d\n", i)
	}
	close(ch)
}
