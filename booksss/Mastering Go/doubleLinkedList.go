package main

import "fmt"

type DoubleLinkedNode struct {
	Value int
	Pre   *DoubleLinkedNode
	Next  *DoubleLinkedNode
}
var head = new(DoubleLinkedNode)

func addDoubleLinkNode(t *DoubleLinkedNode, v int) int {
	if head == nil {
		t = &DoubleLinkedNode{
			Value: 0,
			Pre:   nil,
			Next:  nil,
		}
		head = t
		return 0
	}

	if v == t.Value {
		fmt.Println("Node already exists", v)
		return -1
	}

	if t.Next == nil {
		tmp := t
		t.Next = &DoubleLinkedNode{
			Value: v,
			Pre:   tmp,
			Next:  nil,
		}
		return -2
	}

	return addDoubleLinkNode(t.Next, v)
}

func traverseDoubleLink(t *DoubleLinkedNode)  {
	if t == nil {
		fmt.Println("-> Empty list")
		return
	}

	for t != nil {
		fmt.Printf("%d ->", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func reverse(t *DoubleLinkedNode) {
	if t == nil {
		fmt.Println("-> Empty list")
		return
	}
	tmp := t
	for t != nil {
		tmp = t
		t = t.Next
	}
	for tmp.Pre != nil {
		fmt.Printf("%d ->", tmp.Value)
		tmp = tmp.Pre
	}
	fmt.Printf("%d ->", tmp.Value)
	fmt.Println()
}

func DoubleSize(t *DoubleLinkedNode) int {
	if t == nil {
		fmt.Println("-> Empty")
		return 0
	}
	n := 0
	for t != nil {
		n++
		t = t.Next
	}
	return n
}

func lookUpDoubleNode(t *DoubleLinkedNode, v int) bool {
	if head == nil {
		return false
	}
	if v == t.Value {
		return true
	}
	if t.Next == nil {
		return false
	}
	return lookUpDoubleNode(t.Next, v)
}

func main() {
	fmt.Println(head)
	head = nil
	traverseDoubleLink(head)

	addDoubleLinkNode(head, 1)
	addDoubleLinkNode(head, 1)

	traverseDoubleLink(head)

	addDoubleLinkNode(head, 10)
	addDoubleLinkNode(head, 5)
	addDoubleLinkNode(head, 0)
	addDoubleLinkNode(head, 0)

	traverseDoubleLink(head)

	addDoubleLinkNode(head, 100)
	fmt.Println(DoubleSize(head))

	traverseDoubleLink(head)
	reverse(head)
}
