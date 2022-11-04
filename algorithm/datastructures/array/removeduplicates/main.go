package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 3, 3, 3, 4, 5, 5}
	fmt.Println(removeDuplicates(arr))
}

// 给你一个 升序排列 的数组 nums ，
// 请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
// 元素的相对顺序应该保持 一致 。

// removeDuplicates
// 解题要点
// 1. 原地不能利用额外的空间
// 2. 数组是升序的
// 3. 双指针
func removeDuplicates(nums []int) int {
	p, q := 0, 1
	for q < len(nums) {
		if nums[p] == nums[q] {
			q++
		} else {
			// 覆盖p后面的那一位
			nums[p+1] = nums[q]
			p++
		}
	}
	return p + 1
}
