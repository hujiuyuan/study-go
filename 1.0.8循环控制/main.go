package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	for_demo()
	break_demo()
	continue_demo()
	goto_demo()
}

/*
for 循环
声明方式 1：

	for <init>; <condition>;<post> {
	    <expression>
	    ...
	}

声明方式 2，可以仅声明条件判断语句：

	for <condition> {
	    <expression>
	    ...
	}

声明方式 3，无限循环：

	for {
	    <expression>
	    ...
	}

声明方式 4，搭配 range 关键字：
// 遍历切片，下标参数为可选参数

	for <position name>[, <element var name>] := range <slice/array name> {
	    <expression>
	    ...
	}

// 遍历map，value值为可选参数

	for <key var name>[, <value var name>] := range <map name> {
	    <expression>
	    ...
	}
*/
func for_demo() {
	// 方式1
	for i := 0; i < 10; i++ {
		fmt.Println("方式1，第", i+1, "次循环")
	}

	// 方式2
	b := 1
	for b < 10 {
		fmt.Println("方式2，第", b, "次循环")
		b++
	}

	//// 方式3，无限循环
	// 创建一个带有超时时间的上下文，超时时间为当前时间之后的2秒
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))

	// 定义一个布尔变量，用于标记子goroutine是否已经启动
	var started bool

	// 定义一个原子布尔变量，用于标记子goroutine是否已经停止
	// 这里使用了sync/atomic包中的原子操作，确保在并发环境下的线程安全
	var stopped atomic.Bool

	// 主goroutine进入一个无限循环
	for {
		// 如果子goroutine还未启动
		if !started {
			// 标记子goroutine已经启动
			started = true

			// 启动一个子goroutine
			go func() {
				// 子goroutine进入一个无限循环
				for {
					// 使用select语句监听ctx.Done()通道
					select {
					// 如果ctx.Done()通道收到通知，说明超时时间已到
					case <-ctx.Done():
						// 打印"ctx done"，表示子goroutine收到了停止信号
						fmt.Println("ctx done")

						// 将stopped设置为true，表示子goroutine已经停止
						stopped.Store(true)

						// 退出子goroutine
						return
					}
				}
			}()
		}

		// 主goroutine在每次循环中打印"main"
		fmt.Println("main")

		// 检查stopped的值
		if stopped.Load() {
			// 如果stopped为true，说明子goroutine已经停止
			// 主goroutine退出循环
			break
		}
	}

	// 主goroutine退出循环后，程序结束

	// 遍历数组
	var a [10]string
	a[0] = "Hello"
	for i := range a {
		fmt.Println("当前下标：", i)
	}
	for i, e := range a {
		fmt.Println("a[", i, "] = ", e)
	}

	// 遍历切片
	s := make([]string, 10)
	s[0] = "Hello"
	for i := range s {
		fmt.Println("当前下标：", i)
	}
	for i, e := range s {
		fmt.Println("s[", i, "] = ", e)
	}

	m := make(map[string]string)
	m["b"] = "Hello, b"
	m["a"] = "Hello, a"
	m["c"] = "Hello, c"
	for i := range m {
		fmt.Println("当前key：", i)
	}
	for k, v := range m {
		fmt.Println("m[", k, "] = ", v)
	}
}

/*
break 语句可用于以下几个场景：

	终止当前循环，并跳出循环。
	中断 switch 语句，提前跳出 case 代码块。
	中断 select 语句，提前跳出 select 代码块。
	嵌套循环中，可以用 label 标出想 break 的循环。
*/
func break_demo() {
	fmt.Println("====================")
	// 中断for循环
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println("第", i, "次循环")
	}

	// 中断switch
	switch i := 1; i {
	case 1:
		fmt.Println("进入case 1")
		if i == 1 {
			break
		}
		fmt.Println("i等于1")
	case 2:
		fmt.Println("i等于2")
	default:
		fmt.Println("default case")
	}

	//for {

	// 中断select
	select {
	case <-time.After(time.Second * 2):
		fmt.Println("进过了2秒")

	case <-time.After(time.Second):
		fmt.Println("过了1秒")

		if true {
			break
		}
		fmt.Println("break 之后")
	}
	//}

	// 不使用标记
	for i := 1; i <= 3; i++ {
		fmt.Printf("不使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("不使用标记,内部循环 j = %d\n", j)
			break
		}
	}

	// 使用标记
outter:
	for i := 1; i <= 3; i++ {
		fmt.Printf("使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("使用标记,内部循环 j = %d\n", j)
			break outter
		}
	}
}

/*
continue 语句类似 break 语句，但是仅能在循环中使用。

continue 语句不是中断循环，而是跳过当前循环的执行，并开始执行下一次循环语句。

而且 continue 语句会执行 for 循环的 post 语句。

在嵌套循环中，可以使用标号 label 标出想 continue 的循环。
*/
func continue_demo() {
	fmt.Println("=============================")
	// 中断for循环
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("第", i, "次循环")
	}

	// 不使用标记
	for i := 1; i <= 2; i++ {
		fmt.Printf("不使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("不使用标记,内部循环 j = %d\n", j)
			if j >= 7 {
				continue
			}
			fmt.Println("不使用标记，内部循环，在continue之后执行")
		}
	}

	// 使用标记
outter:
	for i := 1; i <= 3; i++ {
		fmt.Printf("使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("使用标记,内部循环 j = %d\n", j)
			if j >= 7 {
				continue outter
			}
			fmt.Println("不使用标记，内部循环，在continue之后执行")
		}
	}
}

/*
goto 语句可以无条件转移到指定 labal 标出的代码处。一般 goto 语句会配合条件语句使用，实现条件转移，构成循环，跳出循环的功能。

一般不推荐使用 goto 语句，goto 语句会增加代码流程的混乱，不容易理解代码和调试程序
*/
func goto_demo() {
	fmt.Println("---------------------------")
	gotoPreset := false

preset:
	a := 5

process:
	if a > 0 {
		a--
		fmt.Println("当前a的值为：", a)
		goto process
	} else if a <= 0 {
		// elseProcess:
		if !gotoPreset {
			gotoPreset = true
			goto preset
		} else {
			goto post
		}
	}

post:
	fmt.Println("main将结束，当前a的值为：", a)
}
