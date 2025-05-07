package main

import "fmt"

func int_demo() {
	// int，int8，int16，int32，int64，uint，uint8，uint16，uint32，uint64，uintptr。
	// 十六进制
	var a uint8 = 0xF
	var b uint8 = 0xf

	// 八进制
	var c uint8 = 017
	var d uint8 = 0o17
	var e uint8 = 0o17

	// 二进制
	var f uint8 = 0b1111
	var g uint8 = 0b1111

	// 十进制
	var h uint8 = 15

	println(a, b, c, d, e, f, g, h)
}

func float_demo() {

	/*
		对于浮点类型需要被自动推到的变量，其类型都会被自动设置为 float64，而不管它的字面量是否是单精度。
	*/
	var float1 float32 = 10
	float2 := 10.0

	//float1 = float2
	// cannot use float2 (variable of type float64) as float32 value in assignment

	float1 = float32(float2)

	println(float1, float2)
}

func complex_demo() {
	/**
	虚数单位：
		数的世界里还有一个“空白”。比如，你遇到这样一个问题：“一个数的平方等于-1，这个数是多少？”在实数的世界里，任何数的平方都是非负的（比如 2^2 = 4 ， (-2)^2 = 4 ），
	所以没有一个实数的平方是-1。这就像是一个“无解”的问题。但是，数学家们不喜欢“无解”，他们决定发明一种新的数来解决这个问题。于是，他们发明了**虚数单位 i **，并规定 i^2 = -1 。
	这样，问题就解决了： i 的 平方等于 -1

	复数：
		有了虚数单位 i 之后，数学家们发现，实数和虚数可以组合在一起，形成一种新的数，叫做复数。复数的形式是 a + bi ，其中：• a 是实数部分（实部）。• b 是虚数部分（虚部）。• i 是虚数单位
	*/
	var c1 complex64
	c1 = 1.10 + 0.1i
	c2 := 1.10 + 0.1i
	c3 := complex(1.10, 0.1) // c2与c3是等价的

	x := real(c2)
	y := imag(c2)
	fmt.Println(c1, c2, c3, x, y)
}

func byte_demo() {
	var s string = "Hello, world!"
	var bytes []byte = []byte(s)
	fmt.Println("convert \"Hello, world!\" to bytes: ", bytes)
}

func rune_demo() {
	var r1 rune = 'a'
	var r2 rune = '世'
	var s string = "abc，你好，世界！"
	var runes []rune = []rune(s)

	fmt.Println(r1, r2, s, runes)
}

func string_demo() {
	var s1 string = "Hello\nworld!\n"
	var s2 string = `Hello
world!
`
	fmt.Println(s1 == s2)
}

func byte_rune_string() {
	var s string = "Go语言"
	var bytes []byte = []byte(s)
	var runes []rune = []rune(s)

	fmt.Println("string length: ", len(s))
	fmt.Println("bytes length: ", len(bytes))
	fmt.Println("runes length: ", len(runes))

	fmt.Println("string sub: ", s[0:7])
	fmt.Println("bytes sub: ", string(bytes[0:7]))
	fmt.Println("runes sub: ", string(runes[0:3]))
}

func bool_demo() {
	var a bool = true
	var b bool = false
	fmt.Println(a && b)
}

func main() {
	int_demo()
	float_demo()
	complex_demo()
	byte_demo()
	rune_demo()
	string_demo()
	byte_rune_string()
}
