package main

import "fmt"

/*
控制流程

 136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。

找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
func main() {
	num := demo([]int{1, 2, 5, 4, 8, 4, 6, 5, 2, 1})
	fmt.Println(num)
}

func demo(arr []int) int {
	m := make(map[int]int)

	fmt.Println(m)
	for _, i := range arr {
		m[i] = m[i] + 1
	}

	for key, value := range m {
		if value == 1 {
			return key
		}
	}
	return 0
}
