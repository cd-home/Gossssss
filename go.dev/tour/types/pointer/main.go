package main

import "fmt"

// Go has pointers. A pointer holds the memory address of a value.
func main() {
	// The type *T is a pointer to a T value. Its zero value is nil.
	var p *int
	i := 66

	// The & operator generates a pointer to its operand.
	// operand: variable
	// 获取地址
	p = &i

	// The * operator denotes the pointer's underlying value [底层值].
	// 解引用、间接引用
	fmt.Println(*p)

	// 修改值
	*p = 99
	fmt.Println(i)
}
