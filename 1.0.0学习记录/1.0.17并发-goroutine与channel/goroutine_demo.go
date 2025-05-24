package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	goroutine_demo()
	goroutine_demo2()
}

/*
go 支持并发的方式，就是通过 goroutine 和 channel 提供的简洁且高效的方式实现的。
goroutine 是轻量线程，创建一个 goroutine 所需的资源开销很小，所以可以创建非常多的 goroutine 来并发工作。

它们是由 Go 运行时调度的。调度过程就是 Go 运行时把 goroutine 任务分配给 CPU 执行的过程。

但是 goroutine 不是通常理解的线程，线程是操作系统调度的。

在 Go 中，想让某个任务并发或者异步执行，只需把任务封装为一个函数或闭包，交给 goroutine 执行即可。

声明方式 1，把方法或函数交给 goroutine 执行：
go <method_name>(<method_params>...)

声明方式 2，把闭包交给 goroutine 执行：

	go func(<method_params>...){
	    <statement_or_expression>
	    ...
	}(<params>...)
*/
func goroutine_demo() {
	go func() {
		fmt.Println("run goroutine in closure")
	}()
	go func(msg string) {
		fmt.Println(msg)
	}("gorouine: closure params")
	go say("in goroutine: world")
	say("hello")
}
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

/*
go 中并发同样存在线程安全问题，因为 Go 也是使用共享内存让多个 goroutine 之间通信。并且大部分时候为了性能，所以 go 的大多数标准库的数据结构默认是非线程安全的。
*/
func goroutine_demo2() {
	counter := UnsafeCounter{}

	// 启动100个goroutine同时增加计数
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
		}()
	}

	// 等待一段时间确保所有goroutine完成
	time.Sleep(time.Second)

	// 输出最终计数
	fmt.Printf("Final count: %d\n", counter.GetCount())
}

// 线程安全的计数器
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// 增加计数
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// 获取当前计数
func (c *SafeCounter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

type UnsafeCounter struct {
	count int
}

// 增加计数
func (c *UnsafeCounter) Increment() {
	c.count += 1
}

// 获取当前计数
func (c *UnsafeCounter) GetCount() int {
	return c.count
}
