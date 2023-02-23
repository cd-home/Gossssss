package main

import "fmt"

func main() {
	r := Rect{
		Width:  10,
		Height: 10,
	}
	// 无论是值receiver 还是指针receiver， 值类型或者指针类型都可以调用
	// 会自动处理值和指针之间的转换
	fmt.Println(r.Area())

	fmt.Println(r.Perimeter())

	rp := &r

	fmt.Println(rp.Area())
	fmt.Println(rp.Perimeter())
}

type Rect struct {
	Width, Height int
}

// *Rect类型的receiver
// Method
func (r *Rect) Area() int {
	return r.Width * r.Height
}

// Rect类型的receiver
func (r Rect) Perimeter() int {
	return 2 * (r.Width + r.Height)
}