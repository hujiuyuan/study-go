package main

import (
	"fmt"
	"sync"
	"time"
)

/*
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。

	考察点 ：协程原理、并发任务调度
*/
func main() {
	tasks := []Task{
		{
			Name: "任务1",
			Func: func() {
				time.Sleep(5 * time.Second)
				fmt.Printf("时间：%s，完成任务：%s", time.Now(), "任务1")
			},
		},
		{
			Name: "任务2",
			Func: func() {
				time.Sleep(2 * time.Second)
				fmt.Printf("时间：%s，完成任务：%s", time.Now(), "任务2")
			},
		},
		{
			Name: "任务3",
			Func: func() {
				time.Sleep(4 * time.Second)
				fmt.Printf("时间：%s，完成任务：%s", time.Now(), "任务3")
			},
		},
	}

	ExecuteTasks(tasks)
}

func ExecuteTasks(tasks []Task) {
	// 创建 任务缓冲通道
	resultChan := make(chan Task, len(tasks))

	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go func(t Task) {
			defer wg.Done()
			t.StartTime = time.Now()
			// 执行任务
			t.Func()
			t.Duration = time.Since(t.StartTime)
			// 任务执行结果丢入缓冲通道
			resultChan <- t
		}(task)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// 收集并输出结果
	fmt.Println("任务执行统计：")
	for result := range resultChan {
		fmt.Println()
		fmt.Printf("日志-任务: %s | 开始时间: %s | 耗时: %v\n",
			result.Name,
			result.StartTime.Format("15:04:05.000"),
			result.Duration.Round(time.Millisecond))
	}
}

/*
定义一个任务 的类型
*/
type Task struct {
	Name      string
	Duration  time.Duration
	StartTime time.Time
	Func      func()
}
