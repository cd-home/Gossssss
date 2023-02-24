package main

import "fmt"

func main() {
	r := Rect{
		Width:  10,
		Height: 10,
	}
	// 无论是值receiver 还是指针receiver， 值类型或者指针类型实例都可以调用
	// Go 会自动处理值和指针之间的转换 以此来进行方法的调用
	fmt.Println(r.Area())

	fmt.Println(r.Perimeter())

	rp := &r

	fmt.Println(rp.Area())
	fmt.Println(rp.Perimeter())
}

type Rect struct {
	Width, Height int
}

// 通常推荐同一个 结构体采用相同的 receiver 类型

// Perimeter Rect类型的receiver
func (r Rect) Perimeter() int {
	return 2 * (r.Width + r.Height)
}

// Area *Rect类型的receiver
// 指针接收者可以减少接收者的拷贝消耗, 并且可以修改接收者
func (r *Rect) Area() int {
	return r.Width * r.Height
}
