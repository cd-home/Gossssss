package main

import "fmt"

// Vertex
// A struct is a collection of fields.
// 结构体就是多个字段的集合
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)

	// access by dot
	fmt.Println(v.Y)

	// modify
	v.X = 2
	fmt.Println(v)

	// pointer, 结构体指针可直接访问字段
	p := &v
	// do not need (*p).Y, 不需要显式的解引用去访问属性
	p.Y = 10
	fmt.Println(v)

	// diff init ways
	// 字面量形式, 推荐字段对应的赋值方式
	var (
		v1 = Vertex{1, 2} // has type Vertex
		v2 = Vertex{X: 1} // Y:0 is implicit
		v3 = Vertex{}     // X:0 and Y:0
		// & return pointer
		pointer = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(pointer)
}
