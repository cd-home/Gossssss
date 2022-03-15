package algssss

import (
	"testing"
)

func TwoSum(arr []int, target int) []int {
	store := make(map[int]int)
	for kf, v := range arr {
		dif := target - v
		if ks, ok := store[dif]; ok {
			return []int{ks, kf}
		}
		store[v] = kf
	}
	return []int{-1, -1}
}

func TestTwoSum(t *testing.T) {
	t.Log(TwoSum([]int{2, 7, 12, 4}, 19))
	t.Log(TwoSum([]int{2, 7, 12, 4}, 9))
}