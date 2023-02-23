package generics_test

import "testing"

// Reverse
// any
func Reverse[T any](slices []T) []T {
	l := len(slices)
	for i := 0; i < l/2; i++ {
		pivot := l - i - 1
		slices[i], slices[pivot] = slices[pivot], slices[i]
	}
	return slices
}

func TestGenericsReverse(t *testing.T) {
	strings := []string{"a", "b", "c", "d"}
	t.Log(Reverse(strings))

	int64s := []int{1, 2, 3, 4}
	t.Log(Reverse(int64s))
}
