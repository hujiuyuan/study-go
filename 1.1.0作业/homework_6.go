package main

import "fmt"

/*
加一
简单
相关标签
相关企业

给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。
*/
func main() {
	fmt.Println(demo6([]int{1, 1, 1, 2, 2, 3, 3, 9}))
}

func demo6(nums []int) []int {
	return add_1(nums, len(nums)-1)
}

func add_1(nums []int, i int) []int {
	index := i
	result := []int{}
	if nums[index] == 9 && index > 0 {
		result = add_1(nums, index-1)
	} else if nums[index] == 9 && index == 0 {
		result = append(result, 1)
		result = append(result, nums[:index]...)
	} else if nums[index] < 9 {
		result = append(result, nums[:index]...)
		result = append(result, nums[index]+1)
	}
	return result
}
