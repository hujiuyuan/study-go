package main

import "fmt"

func main() {
	const_demo()
}

/*
由于常量的值是在编译期确定的，所以常量定义时必须赋值，并且不能使用方法的返回值为常量赋值。

常量被定义后，其值不能再被修改。

常量（包括全局常量和局部常量）被定义后可以不使用。

常量的定义方式与变量定义的方式基本相同，只是 var 关键字被更换成了 const：
方式 1：
const <name> <type> = <value>
方式 2：
// 某些类型可以直接推导出来，不需要声明
const <name> = <value>
方式 4：
const <name5>, <name6>, ... <type> = <value5>, <value6>, ...
方式 5：
// 声明多个时，可以用小括号包裹，此模式不限制声明次数
const (

	<name1> <type1> = <value1>
	<name2> = <value2>
	<name3>, <name4>, ... = <value3>, <value4>, ...
	<name5>, <name6>, ... <type> = <value5>, <value6>, ...

)

注：Go 中，常量只能使用基本数据类型，即数字、字符串和布尔类型。不能使用复杂的数据结构，比如切片、数组、map、指针和结构体等。如果使用了非基本数据类型，会在编译期报错。
*/
func const_demo() {

	// 方式1
	const a int = 1

	// 方式2
	const b = "test"

	// 方式3
	const c, d = 2, "hello"

	// 方式4
	const e, f bool = true, false

	// 方式5
	const (
		h    byte = 3
		i         = "value"
		j, k      = "v", 4
		l, m      = 5, false
	)

	const (
		n = 6
	)
	fmt.Println(a, b, c, d, e, f, h, i, j, k, l, m, n)
}

/*
Go 中没有内置枚举类型，所以 Go 中的枚举是使用 const 来定义枚举的。
枚举的本质就是一系列的常量。所以 Go 中使用 const 定义枚举，比如：
const (

	Male = "Male"
	Female = "Female"

)
除了直接定义值以外，还会使用类型别名，让常量定义的枚举类型的作用显得更直观，比如：
type Gender string
const (

	Male   Gender = "Male"
	Female Gender = "Female"

)
当此枚举作为参数传递时，会使用 Gender 作为参数类型，而不是基础类型 string，比如：
func method(gender Gender) {}

并且使用了类型别名后，还可以为这个别名类型声明自定义方法：

	func (g *Gender) String() string {
	    switch *g {
	    case Male:
	        return "Male"
	    case Female:
	        return "Female"
	    default:
	        return "Unknown"
	    }
	}

	func (g *Gender) IsMale() bool {
	    return *g == Male
	}
*/
type Gender byte

/*
除了上面的别名类型来声明枚举类型以外，还可以使用 iota 关键字，来自动为常量赋值。
如果 iota 定义在 const 定义组中的第 n 行，那么 iota 的值为 n (从0开始算)。所以一定要注意 iota 出现在定义组中的第几行，而不是当前代码中它第几次出现。
*/
const (
	Male Gender = iota
	Female
)
