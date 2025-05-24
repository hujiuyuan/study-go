package main

import (
	"fmt"
	"math"
)

/*
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。

		在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。

	    考察点 ：接口的定义与实现、面向对象编程风格
*/
func main() {
	r := Rectangle{2, 3}
	fmt.Printf("Rectangle width:%f, height:%f Area:%f Perimeter:%f \n", r.width, r.height, r.Area(), r.Perimeter())
	c := Circle{5}
	fmt.Printf("Circle radius:%f, Area:%f Perimeter:%f \n", c.radius, c.Area(), c.Perimeter())
}

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return (r.width + r.height) * 2
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
