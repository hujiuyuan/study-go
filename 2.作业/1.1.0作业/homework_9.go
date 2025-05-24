package main

import "fmt"

/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。

你可以按任意顺序返回答案。

示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]

示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]

提示：

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案
*/
func main() {
	fmt.Println(demo9([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5))
	fmt.Println(demo9_1([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5))
}

// 循环两次 匹配 计算结果
func demo9(nums []int, sum int) []int {
	length := len(nums)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if sum-nums[i] == nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}

func demo9_1(nums []int, sum int) []int {
	target_map := make(map[int]int)
	// 1. 循环 数组中的每一个元素
	for index, value := range nums {
		// 3.判断 每个元素 的差值 是否之前 已经存过， 取出下标
		target, ok := target_map[sum-value]
		if ok {
			// 返回结果
			return []int{target, index}
		}
		// 2. 把每个元素 和目标值的 差值存起来， 差值做下标，value 做值
		target_map[value] = index
	}
	return nil
}
