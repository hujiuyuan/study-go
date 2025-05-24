package main

import "fmt"

func main() {
	// 基础运算符
	demo1()
	// 计算时数字运算需要转换类型
	demo2()
	// 逻辑运算符只会有 bool 一种返回
	demo3()
	// 位运算
	demo4()
	// 赋值运算
	demo5()
	// 1.指针
	demo6()
	// 运算优先级
	demo7()
}

func demo7() {
	fmt.Println("===================")
	var a int = 21
	var b int = 10
	var c int = 16
	var d int = 5
	var e int

	e = (a + b) * c / d // ( 31 * 16 ) / 5
	fmt.Printf("(a + b) * c / d 的值为 : %d\n", e)

	e = ((a + b) * c) / d // ( 31 * 16 ) / 5
	fmt.Printf("((a + b) * c) / d 的值为  : %d\n", e)

	e = (a + b) * (c / d) // 31 * (16/5)
	fmt.Printf("(a + b) * (c / d) 的值为  : %d\n", e)

	// 21 + (160/5)
	e = a + (b*c)/d
	fmt.Printf("a + (b * c) / d 的值为  : %d\n", e)

	// 2 & 2 = 2; 2 * 3 = 6; 6 << 1 = 12; 3 + 4 = 7; 7 ^ 3 = 4;4 | 12 = 12
	f := 3 + 4 ^ 3 | 2&2*3<<1
	fmt.Println(f == 12)
}

func demo6() {
	fmt.Println("===================")
	a := 4
	var ptr *int
	fmt.Println(a)

	ptr = &a
	fmt.Printf("*ptr 为 %d\n", *ptr)
}

func demo4() {
	fmt.Println("===================")
	fmt.Println(0 & 0)
	fmt.Println(0 | 0)
	fmt.Println(0 ^ 0)

	fmt.Println(0 & 1)
	fmt.Println(0 | 1)
	fmt.Println(0 ^ 1)

	fmt.Println(1 & 1)
	fmt.Println(1 | 1)
	fmt.Println(1 ^ 1)

	fmt.Println(1 & 0)
	fmt.Println(1 | 0)
	fmt.Println(1 ^ 0)
}

func demo3() {
	fmt.Println("===================")
	a := true
	b := false

	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(!(a && b))
}

func demo2() {
	fmt.Println("===================")

	a := 10 + 0.1
	b := byte(1) + 1
	fmt.Println(a, b)

	sum := a + float64(b)
	fmt.Println(sum)

	sub := byte(a) - b
	fmt.Println(sub)

	mul := a * float64(b)
	div := int(a) / int(b)
	//div := int(a) / (b); 错误

	fmt.Println(mul, div)
}

/*
另外，自增与自减只能以 <var name>++ 或者 <var name>-- 的模式声明，并且只能单独存在，不能在自增或自减的同时做加减乘除的计算：
a := 1
// 正确写法
a++
a--

// 错误的使用方式
++a
--a

// 错误使用方式，不可以自增时计算,也不能赋值
b := a++ + 1
c := a--
*/
func demo1() {
	a, b := 1, 2
	sum := a + b
	sub := a - b
	mul := a * b
	div := a / b
	mod := a % b

	fmt.Println(sum, sub, mul, div, mod)
}

func demo5() {
	fmt.Println("===================")

	a, b := 1, 2
	var c int
	c = a + b
	fmt.Println("c = a + b, c =", c)

	c += a // c = c + a
	fmt.Println("c += a, c =", c)

	c -= a // c = c - a
	fmt.Println("c -= a, c =", c)

	c *= a // c = c * a
	fmt.Println("c *= a, c =", c)

	c /= a // c = c / a
	fmt.Println("c /= a, c =", c)

	c %= a // c = c % a
	fmt.Println("c %= a, c =", c)

	c <<= a // c = c << a
	fmt.Println("c <<= a, c =", c)

	c >>= a // c = c >> a
	fmt.Println("c >>= a, c =", c)

	c &= a // c = c & a
	fmt.Println("c &= a, c =", c)

	c |= a // c = c | a
	fmt.Println("c |= a, c =", c)

	c ^= a // c = c ^ a
	fmt.Println("c ^= a, c =", c)
}
