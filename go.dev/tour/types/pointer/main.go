package main

import "fmt"

// Go has pointers. A pointer holds the memory address of a value.
func main() {
	// The type *T is a pointer to a T value. Its zero value is nil.
	var p *int
	i := 66

	// The & operator generates a pointer to its operand.
	// operand: variable
	p = &i

	// The * operator denotes the pointer's underlying value.
	fmt.Println(*p)

	//
	*p = 99
	fmt.Println(i)
}
