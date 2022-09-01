package main

import "fmt"

// List
// In addition to[除了] generic functions,
// Go also supports generic types.
// A type can be parameterized with a type parameter,
// which could be useful for implementing generic data structures
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	root := &List[int]{val: 0}
	root.next = &List[int]{val: 1}
	fmt.Println(root.next.val)

	boot := &List[string]{val: "a"}
	boot.next = &List[string]{val: "b"}
	fmt.Println(boot.next.val)
}
