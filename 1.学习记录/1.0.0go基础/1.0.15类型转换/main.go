package main

import (
	"fmt"
	"strconv"
)

func main() {
	change_type_demo()
	change_type_demo2()
	change_type_demo3()

	interface_type_demo()
	interface_type_demo2()

	struct_type_demo()
}

/*
类型转换

类型转换用于将一种数据类型的变量转换为另外一种类型的变量。

在 Go 中，类型转换的基本格式如下：
<type_name>(<expression>)
type_name 为类型。
expression 为有返回值的类型。
*/
func change_type_demo() {
	var i int32 = 17
	var b byte = 5
	var f float32

	// 数字类型可以直接强转
	f = float32(i) / float32(b)
	fmt.Printf("f 的值为: %f\n", f)

	// 当int32类型强转成byte时，高位被直接舍弃
	// int32 的256的二进制表现形式是 00000001 00000000
	// 转成 8 个字节 的 二进制之后 前面的数据被丢弃变成 00000000，即 0
	var i2 int32 = 256
	var b2 byte = byte(i2)
	fmt.Printf("b2 的值为: %d\n", b2)
}

/*
字符串类型转换
前面的部分章节会提到 string 类型、[]byte 类型与[]rune 类型之间可以类似数字类型那样相互转换，并且数据不会有任何丢失。
*/
func change_type_demo2() {
	str := "hello, 123, 你好"
	var bytes []byte = []byte(str)
	var runes []rune = []rune(str)
	fmt.Printf("bytes 的值为: %v \n", bytes)
	fmt.Printf("runes 的值为: %v \n", runes)

	str2 := string(bytes)
	str3 := string(runes)
	fmt.Printf("str2 的值为: %v \n", str2)
	fmt.Printf("str3 的值为: %v \n", str3)
}

/*
但是也会经常有数字与字符串相互转换的需求，这时需要使用到 go 提供的标准库 strconv。

strconv 可以把数字转成字符串，也可以把字符串转换成数字。

最常见的转换是字符串与 int 类型之间相互转换。也就是 Atoi 方法与 Itoa 方法。

当需要把字符串转换成无符号数字时，目前只能转换成 uint64 类型，需要其他位的数字类型需要从 uint64 类型转到所需的数字类型。

同时可以看到当使用 ParseUint 方法把字符串转换成数字时，或者使用 FormatUint 方法把数字转换成字符串时，都需要提供第二个参数 base，这个参数表示的是数字的进制，即标识字符串输出或输入的数字进制。
*/

func change_type_demo3() {
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	fmt.Printf("字符串转换为int: %d \n", num)
	str1 := strconv.Itoa(num)
	fmt.Printf("int转换为字符串: %s \n", str1)

	ui64, err := strconv.ParseUint(str, 10, 32)
	fmt.Printf("字符串转换为uint64: %d \n", num)

	str2 := strconv.FormatUint(ui64, 2)
	fmt.Printf("uint64转换为字符串: %s \n", str2)
}

/*
接口类型转换

接口类型只能通过断言将转换为指定类型。
<variable_name>.(<type_name>)
variable_name 是变量名称，type_name 是类型名称。
通过断言方式可以同时得到转换后的值以及转换是否成功的标识。
代码示例：
*/
func interface_type_demo() {
	fmt.Println("接口类型转换=================")
	var i interface{} = 3
	// 尝试将 i 判断为 int 类型， 如果是int类型那么 a = i，ok为true，如果不是 int 类型，那么 a 就相当于 int 的初始值 0 ，然后 ok 为 false
	a, ok := i.(int)
	if ok {
		fmt.Printf("'%d' is a int \n", a)
	} else {
		fmt.Println("conversion failed")
	}
}

/*
断言可以直接用做 switch判断
*/
func interface_type_demo2() {
	fmt.Println("断言用作switch判断==============")
	var i interface{} = 3
	// 此处 i.(type) 会拿到 i 的实际类型， switch 会根据实际类型，选择具体的分支去执行代码逻辑
	switch v := i.(type) {
	case int:
		fmt.Println("i is a int", v)
	case string:
		fmt.Println("i is a string", v)
	default:
		fmt.Println("i is unknown type", v)
	}
}

/*
结构体类型转换

结构体类型之间在一定条件下也可以转换的。

当两个结构体中的字段名称以及类型都完全相同，仅结构体名称不同时，这两个结构体类型即可相互转换。
*/
type SameFieldA struct {
	name  string
	value int
}

type SameFieldB struct {
	name  string
	value int
}

func (s *SameFieldB) getValue() int {
	return s.value
}

func struct_type_demo() {
	fmt.Println("结构体转换=============")
	a := SameFieldA{
		name:  "a",
		value: 1,
	}

	b := SameFieldB(a)
	fmt.Printf("conver SameFieldA to SameFieldB, value is : %d \n", b.getValue())

	// 只能结构体类型实例之间相互转换，指针不可以相互转换
	var c interface{} = &a
	_, ok1 := c.(*SameFieldA)
	fmt.Printf("c is *SameFieldA: %v \n", ok1)

	_, ok2 := c.(*SameFieldB)
	fmt.Printf("c is *SameFieldB: %v \n", ok2)
}
