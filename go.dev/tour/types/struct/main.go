package main

import "fmt"

// A struct is a collection of fields.

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)

	// access
	fmt.Println(v.Y)

	// modify
	v.X = 2
	fmt.Println(v)

	// pointer
	p := &v
	p.Y = 10
	fmt.Println(v)

	// diff init ways
	var (
		v1      = Vertex{1, 2}  // has type Vertex
		v2      = Vertex{X: 1}  // Y:0 is implicit
		v3      = Vertex{}      // X:0 and Y:0
		pointer = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(pointer)
}
