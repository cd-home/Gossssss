package main

import "fmt"

func main() {
	arr := []int{3, 2, 2, 3}
	fmt.Println(removeElement(arr, 3))
}

// 给你一个数组 num 和一个值 val，你需要 原地 移除所有数值等于val的元素，并返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

// removeElement
// 原地移除、考虑异常情况
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
