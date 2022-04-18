package main

import "fmt"

var root = new(LinkNode)

type LinkNode struct {
	Value int
	Next  *LinkNode
}

func addNode(t *LinkNode, v int) int {
	if root == nil {
		t = &LinkNode{
			Value: v,
			Next:  nil,
		}
		root = t
		return 0
	}
	if v == t.Value {
		fmt.Println("Node already exists", v)
		return -1
	}
	if t.Next == nil {
		t.Next = &LinkNode{
			Value: v,
			Next:  nil,
		}
		return -2
	}
	return addNode(t.Next, v)
}

func traverseLinkedList(t *LinkNode) {
	if t == nil {
		fmt.Println("-> Empty List")
		return
	}
	for t != nil {
		fmt.Printf("%d ->", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func lookUpNode(t *LinkNode, v int) bool {
	if root == nil {
		t = &LinkNode{
			Value: v,
			Next:  nil,
		}
		root = t
		return false
	}
	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookUpNode(t.Next, v)
}

func size(t *LinkNode) int {
	if t == nil {
		fmt.Println("-> empty list")
		return 0
	}
	i := 0
	for t != nil {
		i += 1
		t = t.Next
	}
	return i
}

func main() {
	fmt.Println(root)
	root = nil
	traverseLinkedList(root)

	addNode(root, 1)
	addNode(root, -1)

	traverseLinkedList(root)

	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 45)
	addNode(root, 5)
	addNode(root, 5)

	traverseLinkedList(root)

	addNode(root, 100)

	traverseLinkedList(root)

	if lookUpNode(root, 100) {
		fmt.Println("100 Yes")
	} else {
		fmt.Println("100 No")
	}

	if lookUpNode(root, -100) {
		fmt.Println("-100 Yes")
	} else {
		fmt.Println("-100 No")
	}

	fmt.Println(size(root))
}
