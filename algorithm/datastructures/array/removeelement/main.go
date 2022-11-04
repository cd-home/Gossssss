package main

import "fmt"

func main() {
	arr := []int{3, 2, 2, 3}
	fmt.Println(removeElement(arr, 3))
}

func removeElement(nums []int, val int) int {
	// 数组为空的情况
	if len(nums) == 0 {
		return 0
	}
	p, q := 0, len(nums)-1
	// 数组只有一个元素, 并且该值==val
	if q == 0 && nums[q] == val {
		return 0
	}
	for p < q {
		if nums[p] == val {
			nums[p] = nums[q]
			q--
		} else {
			p++
		}
	}
	// 解决所有值都是val, 导致最后nums[p] == val的情况
	if nums[q] == val {
		return 0
	}
	// 返回的是 length
	return q + 1
}
