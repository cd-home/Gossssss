package tree

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}

func Walk(tree *Tree) <-chan int {
	ch := make(chan int)
	var walk func(t *Tree)
	walk = func(t *Tree) {
		if t == nil {
			return
		}
		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}
	go func(t *Tree) {
		defer close(ch)
		walk(t)
	}(tree)
	return ch
}

func Same(t1, t2 *Tree) bool {
	ch1 := Walk(t1)
	ch2 := Walk(t2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		// no more
		if !ok1 && !ok2 {
			break
		}
		// diff number
		if ok1 != ok2 {
			return false
		}

		if v1 != v2 {
			return false
		}
	}
	return true
}
