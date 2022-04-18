package main

import (
	"fmt"
	"log"
)

func main() {
	queue = nil
	Push(queue, 10)
	log.Println("QSize", QSize)

	QTraverse(queue)

	v, b := Pop(queue)
	if b {
		log.Println("Pop", v)
	}
	for i := 0; i < 5; i++ {
		Push(queue, i)
	}

	QTraverse(queue)

	log.Println("QSize", QSize)
}

type QNode struct {
	Value int
	Next *QNode
}

var QSize = 0
var queue = new(QNode)

func Push(t *QNode, v int) bool {
	if queue == nil {
		queue = &QNode{
			Value: v,
			Next:  nil,
		}
		QSize++
		return true
	}
	t = &QNode{
		Value: v,
		Next:  nil,
	}
	t.Next = queue
	queue = t
	QSize++
	return true
}

func Pop(t *QNode) (int, bool) {
	if QSize == 0 {
		return 0, false
	}
	if QSize == 1 {
		queue = nil
		QSize--
		return t.Value, true
	}
	tmp := t
	for t.Next != nil {
		tmp = t
		t = t.Next
	}
	v := (tmp.Next).Value
	tmp.Next = nil

	QSize--
	return v, true
}

func QTraverse(t *QNode) {
	if QSize == 0 {
		log.Println("Empty Q")
	}
	for t != nil {
		fmt.Printf("%d->", t.Value)
		t = t.Next
	}
	fmt.Println()
}
