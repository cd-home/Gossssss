package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 返回内存中字节的大小 uintptr
	fmt.Println(unsafe.Sizeof(float64(0)))

	// 只包含数据中固定的部分 pointer len cap
	var s []int
	fmt.Println(unsafe.Sizeof(s))

	// 对于聚合类型
	type Ops struct {
		A bool    // 1
		B float64 // 8
		C int16   // 2
	}
	// 所占字节
	fmt.Println(unsafe.Sizeof(Ops{}))
	// 对齐系数
	fmt.Println(unsafe.Alignof(Ops{}))

	// 字段相对OPS地址的偏移, 由此可以看出填充和对齐
	fmt.Println(unsafe.Offsetof(Ops{}.A))
	fmt.Println(unsafe.Offsetof(Ops{}.B))
	fmt.Println(unsafe.Offsetof(Ops{}.C))

	type Ops2 struct {
		A float64 // 8
		B int16   // 2
		C bool    // 1
	}
	fmt.Println(unsafe.Sizeof(Ops2{}))
	fmt.Println(unsafe.Alignof(Ops2{}))
	fmt.Println(unsafe.Offsetof(Ops{}.A))
	fmt.Println(unsafe.Offsetof(Ops{}.B))
	fmt.Println(unsafe.Offsetof(Ops{}.C))
}
