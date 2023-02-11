package main

import "fmt"

// ArrayQueue 数组构造队列 [受限数据结构]
type ArrayQueue[T any] struct {
	data []T
	max  int
	n    int
}

func NewArrayQueue[T any](max int) *ArrayQueue[T] {
	return &ArrayQueue[T]{
		data: make([]T, 0, max),
	}
}

func (queue *ArrayQueue[T]) IsEmpty() bool {
	return len(queue.data) == 0
}

// EnQueue 入队列
func (queue *ArrayQueue[T]) EnQueue(item T) bool {
	if len(queue.data) < cap(queue.data) {
		queue.data = append(queue.data, item)
		queue.n++
		return true
	}
	return false
}

// DeQueue 出队列
func (queue *ArrayQueue[T]) DeQueue() (item T) {
	if queue.IsEmpty() {
		return item
	}
	item = queue.data[0]
	queue.data = queue.data[1:]
	return item
}

func main() {
	queue := NewArrayQueue[int](5)
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	queue.EnQueue(5)
	fmt.Println(queue.EnQueue(6))
	fmt.Println(queue.DeQueue())
}
