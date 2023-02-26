package main

import "fmt"

func main() {
	arr := []int{1, 1, 2, 3, 45, 67, 4}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func QuickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	// 每次进入partition i 的位置元素就已经排好序
	i := partition(arr, left, right)
	QuickSort(arr, left, i-1)
	QuickSort(arr, i+1, right)
}

func partition(arr []int, left, right int) int {
	mark := left
	// 基准
	pivot := arr[left]
	for i := left + 1; i <= right; i++ {
		if arr[i] < pivot {
			// 把mark指向的大于pivot的值与 i 指向的小于pivot的值交换
			mark++
			arr[mark], arr[i] = arr[i], arr[mark]
		}
		// 实际上可以这样理解, 遇到大于 pivot i 就一直移动, 但是mask没有动
		// 遇到小于pivot的就移动一个然后交换
	}
	// 将基准值pivot 与mark对应的值交换
	arr[left], arr[mark] = arr[mark], arr[left]
	return mark
}
