package example_test

import (
	"testing"
)

type user struct {
	name  string
	email string
}

// escape analysis
// 逃逸分析

func stayOnstack() user {
	u := user{
		name:  "GodYao",
		email: "example.com",
	}
	return u
}

func excapeToHeap() *user {
	u := user{
		name:  "GodYao",
		email: "example.com",
	}
	return &u
}

func TestStayOnStack(t *testing.T) {
	stayOnstack()
}

func TestExcapeToHeap(t *testing.T) {
	excapeToHeap()
}
