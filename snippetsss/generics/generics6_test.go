package generics_test

import "testing"

func Reduce[T1 any, T2 any](slices []T1, start T2, f func(T2, T1) T2) T2 {
	r := start
	for _, item := range slices {
		r = f(r, item)
	}
	return r
}

func TestGenericsReduce(t *testing.T) {
	src := []int{1, 2, 3, 4, 5}
	t.Log(Reduce(src, 10, func(r, e int) int {
		return r + e
	}))
}
