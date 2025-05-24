package main

import "fmt"

/*
字符串,

	有效的括号 ,

考察：字符串处理、栈的使用

题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效

链接：https://leetcode-cn.com/problems/valid-parentheses/
*/
func main() {
	//str := "{(}[{}()])"
	str := "(){}[]"
	//bytes := []byte(str)
	//for i := range bytes {
	//	fmt.Println(bytes[i])
	//}

	fmt.Println(demo_3(str))
}

func demo_3(str string) bool {
	char_arr := []rune(str)
	stake := []rune{}
	//left := []rune{'(', '{', '['}

	for i := 0; i < len(char_arr); i++ {
		ch := char_arr[i]
		if '(' == ch || '[' == ch || '{' == ch || len(stake) == 0 {
			// 如果是左括号就入栈
			stake = append(stake, ch)
		} else if ch == ')' || ch == ']' || ch == '}' {
			// 如果是右括号
			stake_top := stake[len(stake)-1]
			if (stake_top == '(' && ch == ')') || (stake_top == '[' && ch == ']') || (stake_top == '{' && ch == '}') {
				// 右括号和栈顶的元素一致,则栈顶元素出栈
				stake = stake[:len(stake)-1]
			}
		}
	}
	return len(stake) == 0
}
