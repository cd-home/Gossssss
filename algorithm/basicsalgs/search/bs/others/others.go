package others

// BinarySearchFirst
// 查找第一个等于给定值的元素
func BinarySearchFirst(arr []int, v int) int {
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
		} else if arr[mid] < v {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid-1] != v {
				return mid
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}

// BinarySearchLast
// 查找最后一个值等于给定值的元素
func BinarySearchLast(a []int, v int) int {
	n := len(a)
	if n == 0 {
		return -1
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) >> 1
		if a[mid] > v {
			high = mid - 1
		} else if a[mid] < v {
			low = mid + 1
		} else {
			if mid == n-1 || a[mid+1] != v {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

// BinarySearchFirstGT
// 查找第一个大于等于给定值的元素
func BinarySearchFirstGT(a []int, v int) int {
	n := len(a)
	if n == 0 {
		return -1
	}

	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high-low)>>1
		if a[mid] >= v {
			if mid == 0 || a[mid-1] < v {
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
func BinarySearchLastLT(a []int, v int) int {
	n := len(a)
	if n == 0 {
		return -1
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high-low)>>1
		if a[mid] > v {
			high = mid - 1
		} else {
			if mid == n-1 || a[mid+1] > v {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}
