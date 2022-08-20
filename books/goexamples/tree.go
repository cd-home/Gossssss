package main

import (
	"log"
	"math/rand"
	"time"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func traverse(t *Tree) {
	if t == nil {
		return
	}
	traverse(t.Left)
	log.Println(t.Value)
	traverse(t.Right)
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{
			Left:  nil,
			Value: v,
			Right: nil,
		}
	}
	if v == t.Value {
		return t
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

func Create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++ {
		tmp := rand.Intn(n * 2)
		t = insert(t, tmp)
	}
	return t
}

func main() {
	tree := Create(10)
	traverse(tree)

	tree = insert(tree, -10)
	tree = insert(tree, -2)

	traverse(tree)
}