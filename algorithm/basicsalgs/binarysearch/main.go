package main

import "fmt"

/*
	三个容易出错的地方：
		循环退出条件、mid 的取值，low 和 high 的更新。
	二分查找虽然性能比较优秀，但应用场景也比较有限。底层必须依赖数组，并且还要求数据是有序的。
	对于较小规模的数据查找，我们直接使用顺序遍历就可以了，二分查找的优势并不明显。
	二分查找更适合处理静态数据，也就是没有频繁的数据插入、删除操作
*/

// BinarySearch 非递归实现
func BinarySearch(arr []int, target int) int {
	n := len(arr)
	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// BinarySearchRecursive 递归实现
func BinarySearchRecursive(arr []int, target int) int {
	n := len(arr)
	return bs(arr, target, 0, n-1)
}

func bs(a []int, target int, low, high int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)>>1
	if a[mid] == target {
		return mid
	} else if a[mid] > target {
		return bs(a, target, low, mid-1)
	} else {
		return bs(a, target, mid+1, high)
	}
}

func main() {
	arr := []int{1, 2, 2, 3, 4, 5, 6, 7, 7, 9, 10, 20}
	fmt.Println(BinarySearch(arr, 7))
}
