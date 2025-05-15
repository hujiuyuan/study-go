package main

import "fmt"

func main() {
	fmt.Println(demo2(1))
}

/*
把数字从前数数，从后往前数；拆成两个数字，
比对两个数字大小，如果 相等或者 相差10倍（精确相差10倍）基本可以确定是回文数，当然得排除掉几种特殊场景
负数不是回文数，末位为 0 的非 0 数字不是回文数，0~9 的个位数是回文数
*/
func demo2(num int) bool {
	// 负数不是回文数，末位为 0 的非 0 数字不是回文数，0~9 的个位数是回文数
	if num < 0 || (num%10 == 0 && num != 0) {
		return false
	}
	reversed := 0        //初始化 reversed 为 0，用于存储反转的后半部分数字
	for num > reversed { // 循环反转数字的后半部分，num 逐步减少位数，reversed 累积反转数字，直到 num <= reversed（通常处理了大约一半位数）
		// 提取 num 的最后一位并添加到 reversed，reversed 左移一位后累积新数字
		reversed = reversed*10 + num%10
		// 移除 num 的最后一位，相当于 num 除以 10
		num /= 10
	}
	// 比较前半部分与反转后半部分，支持奇数位和偶数位回文数
	return num == reversed || num == reversed/10
}
