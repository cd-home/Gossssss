package main

import "fmt"

type CircularQueue[T any] struct {
	queue   []T
	maxsize int
	front   int
	rear    int
}

func NewCircularQueue[T any](maxsize int) *CircularQueue[T] {
	return &CircularQueue[T]{
		queue:   make([]T, maxsize),
		maxsize: maxsize,
	}
}

func (cq *CircularQueue[T]) IsEmpty() bool {
	return cq.rear == cq.front
}

func (cq *CircularQueue[T]) IsFull() bool {
	// 栈满的条件，循环队列会空一格
	return (cq.rear+1)%cq.maxsize == cq.front
}

func (cq *CircularQueue[T]) EnQueue(item T) bool {
	if cq.IsFull() {
		return false
	}
	cq.queue[cq.rear] = item
	// 下一个位置，不能++，因为是环形的
	cq.rear = (cq.rear + 1) % cq.maxsize
	return true
}

func (cq *CircularQueue[T]) DeQueue() (item T) {
	if cq.IsEmpty() {
		return item
	}
	item = cq.queue[cq.front]
	// 设置零值
	var temp T
	cq.queue[cq.front] = temp
	cq.front = (cq.front + 1) % cq.maxsize
	return item
}

func main() {
	cq := NewCircularQueue[int](5)
	cq.EnQueue(1)
	cq.EnQueue(2)
	cq.EnQueue(3)
	cq.EnQueue(4)
	cq.EnQueue(5)
	fmt.Println(cq.EnQueue(6))
	fmt.Println(cq.DeQueue())
	fmt.Println(cq.EnQueue(6))
	fmt.Println(cq.queue)
}
