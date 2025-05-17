package main

import "fmt"

/*
	最长公共前缀,

考察：字符串处理、循环嵌套

题目：查找字符串数组中的最长公共前缀

链接：https://leetcode-cn.com/problems/longest-common-prefix/
*/
func main() {
	str_arr := []string{"flower", "flow", "flight"}

	fmt.Println(demo4(str_arr))
}

func demo4(arr []string) string {
	if len(arr) > 2 {
		str1 := arr[0]
		str2 := arr[1]
		default_prefix := get_same_prefix(str1, str2)

		for index, value := range arr {
			if index > 1 {
				default_prefix = get_same_prefix(default_prefix, value)
			}
		}
		return default_prefix
	}
	return ""
}

func get_same_prefix(str1, str2 string) string {
	fmt.Println("compare str1 = ", str1, ", str2 = ", str2)
	// 获取字符串的公共前缀
	prefix := []rune{}

	rune_1 := []rune(str1)
	rune_2 := []rune(str2)

	minLength := len(rune_1)
	if len(rune_2) < minLength {
		minLength = len(rune_2)
	}

	for index := range minLength {
		fmt.Println("index", index, "value1", string(rune_1[index]), "value2", string(rune_2[index]))
		if rune_1[index] == rune_2[index] {
			prefix = append(prefix, rune_1[index])
		} else {
			break
		}
	}
	prefix_str := string(prefix)
	fmt.Println("same prefix", prefix_str)
	return prefix_str
}
