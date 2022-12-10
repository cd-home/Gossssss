package main

import "fmt"

type Stack[T any] struct {
	data []T
	top  int
}

func NewStack[T any](size int) *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0, size),
		top:  -1,
	}
}

func (stack *Stack[T]) IsEmpty() bool {
	return stack.top < 0
}

func (stack *Stack[T]) Push(item T) {
	stack.top += 1
	stack.data = append(stack.data, item)
}

func (stack *Stack[T]) Pop() (item T) {
	if stack.IsEmpty() {
		return item
	}
	item = stack.data[stack.top]
	stack.top--
	stack.data = stack.data[:stack.top]
	return item
}

func main() {
	stack := NewStack[int](10)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	fmt.Println(stack.top)
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.Pop())
}
