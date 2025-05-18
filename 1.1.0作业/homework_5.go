package main

import "fmt"

/*
基本值类型,
给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。

考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：

	更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
	返回 k 。
*/
func main() {
	fmt.Println(demo5([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func demo5(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	// 用来计算一共多少种元素
	count := 1
	// 下标从 第二个元素开始计算
	for index := 1; index < length; index++ {
		// 如果当前元素之前出现过，那么把这个元素放到计数的位置
		// 否则判断下一个元素是否出现过
		// 前提是 当前元素 >= 前一个元素，也就是 非严格递增排列
		if nums[index] != nums[index-1] {
			nums[count] = nums[index]
			count++
		}
	}
	return count
}
