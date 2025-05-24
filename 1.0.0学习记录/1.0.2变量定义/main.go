package main

import "fmt"

func main() {
	global_variable()
	local_variable()
	variable_demo()
}

/*
多变量定义
全局变量声明方式：
var <name1>, <name2>, ... <type> = <value1>, <value2>, ...
var <name1>, <name2>, ... <type>
var <name1>, <name2>, ... = <value1>, <value2>, ...

局部变量的声明方式
基本与全局变量的声明方式相同，额外多了不用关键字 var 的方式：
<name1>, <name2>, ... := <value1>, <value2>, ...
*/

var a, b, c int = 1, 2, 3

var e, f, g int

var h, i, j = 1, 2, "test"

func variable_demo() {
	var k, l, m int = 1, 2, 3
	var n, o, p int
	q, r, s := 1, 2, "test"
	fmt.Println(a, b, c, e, f, g, h, i, j)
	fmt.Println(k, l, m, n, o, p, q, r, s)
}

/*
局部变量的申明

方式 1，与全局变量的声明方式完全一致：
var <name> <type> = <value>

方式 2，也是与全局变量声明方式完全相同，仅声明，为类型默认零值：
var <name> <type>

方式 3，无需关键字 var，也无需声明类型，Go 通过字面量或表达式推导此变量类型：
<name> := <value>

方式 4，这种方式是全局变量没有的，可以直接在返回值中声明，相当于在方法一开始就声明了这些变量：

	func method() (<name1> <type1>, <name2> <type2>) {
	    return
	}

	func method() (<name1> <type1>, <name2> <type2>) {
	    return <value1>, <value2>
	}
*/
func local_variable() {
	local_variable_method1()
	method2_a, method2_b := local_variable_method2()
	fmt.Println("method2_a:", method2_a, "method2_b:", method2_b)
	method3_a, method3_b := local_variable_method3()
	fmt.Println("method3_a:", method3_a, "method3_b:", method3_b)
	method4_a, method4_4 := local_variable_method4()
	fmt.Println("method4_a:", method4_a, "method4_4:", method4_4)
}

func local_variable_method1() {
	// 方式1，类型推导，用得最多
	a := 1
	// 方式2，完整的变量声明写法
	var b int = 2
	// 方式3，仅声明变量，但是不赋值，
	var c int
	fmt.Println(a, b, c)
}

// 方式4，直接在返回值中声明
func local_variable_method2() (a int, b string) {
	// 这种方式必须声明return关键字
	// 并且同样不需要使用，并且也不用必须给这种变量赋值
	return 1, "test"
}

func local_variable_method3() (a int, b string) {
	a = 1
	b = "test"
	return
}

func local_variable_method4() (a int, b string) {
	return
}

/*
全局变量的使用
注：全局变量允许声明后不使用。

方式1：var <name> <type> = <value>

// 仅声明，但未赋值，为类型默认零值：
方式2：var <name> <type>

// 某些类型可以直接推导出来，不需要声明
方式3：var <name> = <value>

// 声明多个时，可以用小括号包裹，此方式不限制声明次数
方式4：
var (

	<name1> <type1> = <value1>
	<name2> <type2>
	<name3> = <value3>

)
*/
func global_variable() {
	var s1 string = "Hello"
	var zero int
	var b1 = true

	var (
		i  int = 123
		b2 bool
		s2 = "test"
	)

	var (
		group = 2
	)

	fmt.Println(s1, zero, b1, s2, i, b2, group)
}
