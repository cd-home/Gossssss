package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Pointer
	/*
		*T 为指针, 指向变量的指针
		unsafe.Pointer 是一种特殊的指针,
		Pointer represents a pointer to an arbitrary type
		//   - A pointer value of any type can be converted to a Pointer.
		//   - A Pointer can be converted to a pointer value of any type.
		//   - A uintptr can be converted to a Pointer.
		//   - A Pointer can be converted to a uintptr.
	*/
	var x struct {
		a bool
		b int16
		c []int
	}

	// 和 pb := &x.b 等价
	pb := (*int16)(unsafe.Pointer(
		uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b) // "42"

	// 注意不要将 uintptr 赋值给一个变量
}
