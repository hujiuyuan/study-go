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
	fmt.Println(demo6([]int{2, 9, 9, 9}))
}

func demo6(nums []int) []int {
	add_one := false
	for i := len(nums) - 1; i >= 0; i-- {
		// 如果元素 = 9 那么进一
		if nums[i] == 9 {
			nums[i] = 0
			// 如果是首位数字，那么要额外补1
			if i == 0 {
				add_one = true
				break
			}
		} else {
			nums[i] += 1
			break
		}
	}
	var result []int
	if add_one {
		result = append(result, 1)
	}
	result = append(result, nums...)
	return result
}
