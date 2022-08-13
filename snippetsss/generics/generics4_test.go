package generics_test

import (
	"strings"
	"testing"
)

func Filter[T any](src []T, f func(T) bool) []T {
	var r []T
	for _, item := range src {
		if f(item) {
			r = append(r, item)
		}
	}
	return r
}

func TestGenericsFilter(t *testing.T) {
	src := []int{1, 2, 3, 4, 5, 6, 7}
	t.Log(Filter(src, func(e int) bool {
		return e%2 == 1
	}))

	src2 := []string{"nice", "good", "wonderful", "perfect"}
	t.Log(Filter(src2, func(e string) bool {
		return strings.Contains(e, "e")
	}))

}
