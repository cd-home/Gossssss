package main

func main() {

}

type CircularQueue struct {
	q    []interface{}
	cap  int
	head int
	tail int
}

func NewCircularQueue(n int) *CircularQueue {
	return &CircularQueue{
		q:    make([]interface{}, n),
		cap:  n,
		head: 0,
		tail: 0,
	}
}

func (cq *CircularQueue) IsEmpty() bool {
	return cq.head == cq.tail
}

func (cq *CircularQueue) IsFull() bool {
	// 栈满的条件，循环队列会空一格
	return (cq.tail+1)%cq.cap == cq.head
}

func (cq *CircularQueue) EnQueue(v interface{}) bool {
	if cq.IsFull() {
		return false
	}
	cq.q[cq.tail] = v
	// 下一个位置，不能++，因为是环形的
	cq.tail = (cq.tail + 1) % cq.cap
	return true
}

func (cq *CircularQueue) DeQueue() interface{} {
	if cq.IsEmpty() {
		return nil
	}
	v := cq.q[cq.head]
	cq.head = (cq.head + 1) % cq.cap
	return v
}