package generics_test

import (
	"strings"
	"testing"
)

func Map[T1 any, T2 any](slices []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(slices))
	for i, item := range slices {
		r[i] = f(item)
	}
	return r
}

func TestGenericsMap(t *testing.T) {
	src := []int{1, 2, 3, 4, 5}
	t.Log(Map(src, func(e int) int {
		return e * 2
	}))

	src2 := []string{"a", "b", "c", "d"}
	t.Log(Map(src2, func(e string) string {
		return strings.ToUpper(e)
	}))
}
