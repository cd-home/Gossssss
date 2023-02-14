package main

// 剑指 Offer 53 - I. 在排序数组中查找数字 I

/*
	统计目标数字在排序数组中出现的次数。
	对于排序数组, 首先想到二分查找
*/

func search(nums []int, target int) int {
	p1 := 0
	p2 := len(nums) - 1
	for ; p1 < len(nums); p1++ {
		if nums[p1] == target {
			break
		}
	}
	for ; p2 >= p1; p2-- {
		if nums[p2] == target {
			break
		}
	}
	return p2 - p1 + 1
}

// search2
func search2(nums []int, target int) int {
	var helper = func(left, right int, nums []int, target int) int {
		for left <= right {
			mid := (left + right) >> 1
			if nums[mid] <= target {
				// 收缩左边
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return left
	}
	// 找 target的右边界, 找target-1的右边界
	return helper(0, len(nums)-1, nums, target) - helper(0, len(nums)-1, nums, target-1)
}
func main() {

}
