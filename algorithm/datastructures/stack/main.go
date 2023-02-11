package main

import "fmt"

// Stack 数组栈(亦可以使用链表实现, top 操作在head), 受限数据结构
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

// IsEmpty 是否为空
func (stack *Stack[T]) IsEmpty() bool {
	return stack.top == -1
}

// Push 入栈
func (stack *Stack[T]) Push(item T) {
	if len(stack.data) == cap(stack.data) {
		// 栈满
		return
	}
	stack.top += 1
	stack.data = append(stack.data, item)
}

// Pop 出栈
func (stack *Stack[T]) Pop() (item T) {
	// 1. 判断栈空, 返回零值
	if stack.IsEmpty() {
		return item
	}
	// 2. pop
	item = stack.data[stack.top]
	// 3. top--
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
