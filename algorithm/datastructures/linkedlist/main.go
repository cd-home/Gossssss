package main

import "fmt"

// Node 节点
type Node[T any] struct {
	Data T
	Next *Node[T]
}

// LinkedList 链表
type LinkedList[T any] struct {
	// 头节点，方便操作链表
	head   *Node[T]
	length int
}

// NewNode 新建节点
func NewNode[T any](data T) *Node[T] {
	return &Node[T]{Data: data, Next: nil}
}

// NewLinkList 新建链表
func NewLinkList[T any](data T) *LinkedList[T] {
	return &LinkedList[T]{head: NewNode(data), length: 0}
}

// IsEmpty 是否为空
func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.head.Next == nil
}

// InsertHead 头部插入
func (ll *LinkedList[T]) InsertHead(data T) {
	node := NewNode[T](data)
	node.Next = ll.head.Next
	ll.head.Next = node
}

// InsertTail 尾部插入
func (ll *LinkedList[T]) InsertTail(data T) {
	node := NewNode[T](data)
	cur := ll.head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = node
}

func (ll *LinkedList[T]) Show() {
	fmt.Printf("Head->")
	cur := ll.head
	for cur.Next != nil {
		fmt.Print(cur.Data)
		fmt.Printf("->")
		cur = cur.Next
	}
}

func main() {
	ll := NewLinkList[int](0)
	ll.InsertTail(1)
	ll.InsertTail(2)
	ll.InsertTail(3)
	ll.InsertTail(4)
	ll.InsertHead(-1)
	ll.InsertHead(-2)
	ll.Show()
}
