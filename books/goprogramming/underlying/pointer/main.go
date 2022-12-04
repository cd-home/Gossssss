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
	// pT := uintptr(unsafe.Pointer(new(T))) // 提示: 错误!
	// 没有指针引用new的变量

	// 将所有包含变量地址的uintptr类型变量当作BUG处理
	// 同时减少不必要的unsafe.Pointer类型到uintptr类型的转换
	// 在第一个例子中，有三个转换——字段偏移量到uintptr的转换和转回unsafe.Pointer类型的操作——所有的转换全在一个表达式完成

	// 最后, 你可能在大多数情况下没有必要使用 unsafe;
}
