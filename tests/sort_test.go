package testsss

import (
	"sort"
	"testing"
)

type TestSortSuite struct {
	v int
}

func TestSorter(t *testing.T) {
	data := []TestSortSuite{{2}, {1}, {3}, {4}, {5}, {6}, {7}}

	sort.Slice(data, func(i, j int) bool {
		return data[i].v >= data[j].v
	})
	t.Log(data)

	sort.Slice(data, func(i, j int) bool {
		return data[i].v < data[j].v
	})
	t.Log(data)
}
