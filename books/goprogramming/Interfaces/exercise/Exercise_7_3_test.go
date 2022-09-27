package exercise_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type tree struct {
	value int
	left  *tree
	right *tree
}

func Sort(values []int) {
	var root *tree
	for _, value := range values {
		root = add(root, value)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	var s string
	var visit func(tr *tree)
	visit = func(tr *tree) {
		if tr.left != nil {
			visit(tr.left)
		}
		s = fmt.Sprintf("%s %d", s, tr.value)
		if tr.right != nil {
			visit(tr.right)
		}
	}
	visit(t)
	return s
}

func TestExercise73(t *testing.T) {
	rand.Seed(time.Now().Unix())
	data := make([]int, 50)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Int() % 50
	}
	var root *tree
	for _, v := range data {
		root = add(root, v)
	}
	fmt.Println(root)
}
