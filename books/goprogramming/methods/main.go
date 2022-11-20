package main

import (
	"fmt"
	"math"
)

func main() {
	p := Point{
		X: 2,
		Y: 5,
	}
	q := Point{
		X: 5,
		Y: 6,
	}
	// p. 选择器
	// 编译器会根据方法的名字以及接收器来决定具体调用的是哪一个函数
	fmt.Println(p.Distance(q))
}

type Point struct {
	X, Y float64
}

// Distance
// Point Receiver 接收器 [简短、一致] 通常采用类型的某几个字母
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Path
// 可以给包内任意命名类型定义方法, 要注意这个命名的底层类型不是指针或者interface
type Path []Point

// Distance 求线段长度
// 不同的类型可以有相同名字的方法
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		sum += path[i-1].Distance(path[i])
	}
	return sum
}
