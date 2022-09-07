package x

import (
	"testing"
)

func TestForRangeCopyValue(t *testing.T) {
	arr := []string{"a1", "b", "c", "d", "e", "f"}
	for i, v := range arr {
		// v is copy of arr[i]
		// &v 已经是最后一个迭代的值，所有看似不同的&v 都是指向最后一个
		t.Logf("%v: %v: %v\n", i, &v, &arr[i])
	}
}

func TestForRangeCopyValueFixed(t *testing.T) {
	arr := []string{"a1", "b", "c", "d", "e", "f"}
	for i, v := range arr {
		k := v
		t.Logf("%v: %v: %v: %v\n", i, &v, &k, &arr[i])
	}
}

func TestTemporaryPointer(t *testing.T) {
	store := make(map[int]*int)
	for i := 0; i < 5; i++ {
		store[i] = &i
	}
	// v is a value copy
	for _, v := range store {
		t.Log(*v)
	}
}
