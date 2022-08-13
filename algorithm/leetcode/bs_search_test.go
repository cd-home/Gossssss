package algssss

/*
你需要着重掌握它的三个容易出错的地方：循环退出条件、mid 的取值，low 和 high 的更新。
二分查找虽然性能比较优秀，但应用场景也比较有限。底层必须依赖数组，并且还要求数据是有序的。
对于较小规模的数据查找，我们直接使用顺序遍历就可以了，
二分查找的优势并不明显。二分查找更适合处理静态数据，也就是没有频繁的数据插入、删除操作
*/

//非递归实现
func BinarySearch(a []int, v int) int {
	n := len(a)
	if n == 0 {
		return -1
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) >> 1
		if a[mid] == v {
			return mid
		} else if a[mid] > v {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

//递归实现
func BinarySearchRecursive(a []int, v int) int {
	n := len(a)
	if n == 0 {
		return -1
	}

	return bs(a, v, 0, n-1)
}

func bs(a []int, v int, low, high int) int {
	if low > high {
		return -1
	}
	mid := low + (high - low) >> 1
	if a[mid] == v {
		return mid
	} else if a[mid] > v {
		return bs(a, v, low, mid-1)
	} else {
		return bs(a, v, mid+1, high)
	}
}

//查找第一个等于给定值的元素
func BinarySearchFirst(arr []int, v int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high - low) >> 1
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

//查找最后一个值等于给定值的元素
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

//查找第一个大于等于给定值的元素
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

//查找最后一个小于等于给定值的元素
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