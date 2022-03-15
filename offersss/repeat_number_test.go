package offersss

import (
	"sort"
	"testing"
)

func FindRepeatNumber(arr []int) int {
	for i := 0; i < len(arr); i++ {
		for i != arr[i] {
			if arr[i] == arr[arr[i]] {
				return arr[i]
			}
			arr[i], arr[arr[i]] = arr[arr[i]], arr[i]
		}
	}
	return -1
}

func FindRepeatNumberBySort(arr []int) int {
	// slice sort
	sort.Ints(sort.IntSlice(arr))
	for i, _ := range arr {
		if arr[i] == arr[i+1] {
			return arr[i]
		}
	}
	return -1
}

func FindRepeatNumberByMap(arr []int) int {
	mp := make(map[int]bool)
	for _, v := range arr {
		if _, ok := mp[v]; ok {
			return v
		} else {
			mp[v] = true
		}
	}
	return -1
}

func TestFindRepeatNumber(t *testing.T) {
	arr := []int{2, 3, 1, 1, 2, 5, 3}
	t.Log(FindRepeatNumber(arr))

	t.Log(FindRepeatNumberBySort(arr))

	t.Log(FindRepeatNumberByMap(arr))
}
