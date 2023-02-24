package main

import (
	"fmt"
	"math"
)

// Geometry 几何接口， 方法的签名集合
// 接口是一系列方法签名的集合
type Geometry interface {
	Area() float64
	Perimeter() float64
}

type rect struct {
	width, height float64
}

// Area 非侵入式实现接口
func (r rect) Area() float64 {
	return r.width * r.height
}

func (r rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// 实现接口的方法

type circle struct {
	radius float64
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// 实现了接口的方法集, 结构体的实例可以赋值给接口变量
// use instances of these structs as arguments to measure.、
// 但是要注意, 值接收者与指针接收者 实现是不一样的
// 值接收者, 值类型和指针类型都可以赋值接口变量
// 指针接收者, 只能指针实例能赋值接口变量
func measure(g Geometry) {
	// 通过【运算后的】接口变量调用方法
	fmt.Println(g.Area())
	fmt.Println(g.Perimeter())
}

func main() {
	r := rect{
		width:  10,
		height: 10,
	}
	// 实现了接口就可以赋值给接口变量
	measure(r)

	c := circle{radius: 2}
	measure(c)
}
