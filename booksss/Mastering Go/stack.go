package main

import "fmt"

func main() {
	stack = nil
	pop, b := SPop(stack)
	if b {
		fmt.Print(pop, "")
	} else {
		fmt.Println("SPop Failed")
	}

	SPush(100)
	STraverse(stack)

	SPush(200)
	STraverse(stack)

	for i := 0; i < 10; i++ {
		SPush(i)
	}

	for i := 0; i < 15; i++ {
		sPop, b2 := SPop(stack)
		if b2 {
			fmt.Print(sPop, " ")
		} else {
			break
		}
	}
	fmt.Println()
	STraverse(stack)
}

type SNode struct {
	Value int
	Next *SNode
}

var SSize = 0
var stack = new(SNode)

func SPush(v int) bool {
	if stack == nil {
		stack = &SNode{
			Value: v,
			Next:  nil,
		}
		SSize++
		return true
	}
	tmp := &SNode{
		Value: v,
		Next:  nil,
	}
	tmp.Next = stack
	stack = tmp
	SSize++
	return true
}

func SPop(t *SNode) (int, bool) {
	if SSize == 0 {
		return 0, false
	}

	if SSize == 1 {
		SSize = 0
		stack = nil
		return t.Value, true
	}
	stack = stack.Next
	SSize--
	return t.Value, true
}

func STraverse(t *SNode) {
	if SSize == 0 {
		fmt.Println("Empty Stack")
		return
	}

	for t != nil {
		fmt.Printf("%d->", t.Value)
		t = t.Next
	}
	fmt.Println()
}