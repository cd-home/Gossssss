package main

// BinarySearchFirst
// 查找第一个等于给定值的元素
func BinarySearchFirst(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	// 结果
	index := -1
	for left <= right {
		mid := left + (right-left)>>1
		if target == nums[mid] {
			// 记录备胎，收缩右边界
			index = mid
			right = mid - 1
		} else if target < nums[mid] {
			// 收缩右边界
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return index
}

// BinarySearchLast
// 查找最后一个值等于给定值的元素
func BinarySearchLast(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	index := -1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			// 记录备胎, 收缩左边界
			index = mid
			left = mid + 1
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return index
}

// BinarySearchFirstGT
// 查找第一个大于等于给定值的元素
// TODO
func BinarySearchFirstGT(arr []int, v int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] >= v {
			if mid == 0 || arr[mid-1] < v {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}

	return -1
}

// BinarySearchLastLT
// 查找最后一个小于等于给定值的元素
// TODO
func BinarySearchLastLT(arr []int, v int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] > v {
			high = mid - 1
		} else {
			if mid == n-1 || arr[mid+1] > v {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

func main() {

}
