package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Abs
// method on type
// receiver: value type or pointer
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale
// receiver 是复制 v 的值, 当其是指针接收的时候, 就可以修改接收者指向的值
func (v *Vertex) Scale(s float64) {
	v.X = v.X * s
	v.Y = v.Y * s
}

// Scale
// 你可以用如下的函数形式来理解方法以及接受者
func Scale(v *Vertex, s float64) {
	v.X = v.X * s
	v.Y = v.Y * s
}

type MyInt int

// Abs
// method define on no-struct type
// 方法定义在类型上, 注意无法定义在内置类型上
// 注意: 类型与其方法必须在同一个包中
func (m MyInt) Abs() int {
	if m < 0 {
		return int(-m)
	}
	return int(m)
}

func main() {
	v := Vertex{X: 3, Y: 4}
	// pointer receiver:
	// if v == value (&v).Abs()
	// if v == pointer v.Abs()
	fmt.Println(v.Abs())

	// value receiver:
	// if v == value    v.Abs()
	// if v == pointer  (*v).Abs()

	// 总之: 在值调用方法上, go 会做自动处理
	// 那如何选者接受者类型呢, 指针？值？
	// 1. modify the value that its receiver points to
	// 如果你需要在在方法中修改接受者指向的值
	// 2. avoid copying the value on each method call
	// 避免复制接收者的值, 这对于大结构体是友好的

	//v.Scale(10)
	Scale(&v, 10)
	fmt.Println(v.Abs())

	var i MyInt = -10
	fmt.Println(i.Abs())
}
