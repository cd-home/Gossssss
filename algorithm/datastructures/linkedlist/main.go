package main

import "fmt"

// LinkedList 单链表
type LinkedList[T any] struct {
	// 头节点，方便操作链表
	Head *Node[T]
}

// Node 节点
type Node[T any] struct {
	// 数据域
	Data T
	// 指针域
	Next *Node[T]
}

// NewNode 新建节点
func NewNode[T any](data T) *Node[T] {
	return &Node[T]{Data: data, Next: nil}
}

// NewLinkList 新建链表
func NewLinkList[T any](data T) *LinkedList[T] {
	return &LinkedList[T]{Head: NewNode(data)}
}

// IsEmpty 是否为空
func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.Head.Next == nil
}

// InsertHead 头部插入
func (ll *LinkedList[T]) InsertHead(data T) {
	node := NewNode[T](data)
	// node -> head后的
	node.Next = ll.Head.Next
	// head 指向 node
	ll.Head.Next = node
}

// InsertTail 尾部插入
func (ll *LinkedList[T]) InsertTail(data T) {
	node := NewNode[T](data)
	cur := ll.Head
	// 需要找到最后一个
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = node
}

// DeleteK 删除第k个元素
func (ll *LinkedList[T]) DeleteK(k int) {
	// 找到第 k-1个元素
	var pre *Node[T]
	cur := ll.Head
	for i := 1; i < k; i++ {
		pre = cur.Next
	}
	pre.Next = pre.Next.Next
}

// InsertK 从k位置后插入
func (ll *LinkedList[T]) InsertK(k int, data T) {
	cur := ll.Head
	for i := 1; i <= k; i++ {
		cur = cur.Next
	}
	node := NewNode(data)
	node.Next = cur.Next
	cur.Next = node
}

func (ll *LinkedList[T]) Show() {
	fmt.Printf("Head->")
	cur := ll.Head
	for cur.Next != nil {
		fmt.Print(cur.Data)
		fmt.Printf("->")
		cur = cur.Next
	}
	fmt.Println()
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

	ll.DeleteK(3)
	ll.Show()

	ll.InsertK(2, -2)
	ll.Show()
}
