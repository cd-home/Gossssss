package main

import "fmt"

func QuickSort(arr []int) {
	separateSort(arr, 0, len(arr)-1)
}

func separateSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	// 每次进入partition i 的位置元素就已经排好序
	i := partition(arr, left, right)
	separateSort(arr, left, i-1)
	separateSort(arr, i+1, right)
}

func partition(arr []int, left, right int) int {
	// 先把这个中间值保存起来
	// 如果你选左边基准, 那么必须从右边开始移动比较, 避免完全排序特殊情况下 会交换最大的过去
	pivot := arr[left]
	// 直到 left right 相遇
	for left < right {
		// 先移动右边, 如果大于pivot就移动
		for left < right && arr[right] >= pivot {
			right--
		}
		// 直到遇到小于的
		arr[left] = arr[right]
		// right 坑位留出来了

		// left 比 pivot 大的, 填right的坑
		for left < right && arr[left] <= pivot {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = pivot
	return left
}

func main() {
	arr := []int{2, 1, 2, 3, 45, 67, 10}
	QuickSort(arr)
	fmt.Println(arr)
}
