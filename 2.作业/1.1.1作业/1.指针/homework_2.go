package main

import "fmt"

/*
题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。

	考察点 ：指针运算、切片操作。
*/
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	demo1_2(&nums)
	fmt.Println(nums)
}

func demo1_2(nums *[]int) {
	arrays := *nums
	for index, value := range arrays {
		arrays[index] = value * 2
	}
}
