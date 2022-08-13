package main

import (
	"fmt"
	"math"
)

// 几何接口， 方法的签名集合
type Geometry interface {
	Area() float64
	Perimeter() float64
}

type rect struct {
	width, height float64
}

// 非侵入式实现接口
func (r rect) Area() float64 {
	return r.width * r.height
}

func (r rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type circle struct {
	radius float64
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// 测了方法
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