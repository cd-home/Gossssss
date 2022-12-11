package main

import "fmt"

func main() {
	arr := []int{1, 1, 2, 3, 45, 67, 4}
	QuickSort(arr)
	fmt.Println(arr)
}

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
	mark := left
	pivot := arr[left]
	for i := left + 1; i <= right; i++ {
		if arr[i] < pivot {
			mark++
			arr[mark], arr[i] = arr[i], arr[mark]
		}
		// 实际上可以这样理解, 遇到大于 pivot i 就一直移动
		// 遇到小于pivot的就移动一个然后交换
	}
	arr[left], arr[mark] = arr[mark], arr[left]
	return mark
}
