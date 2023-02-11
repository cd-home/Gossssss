package main

import "fmt"

func InsertionSort(arr []int) {
	n := len(arr)
	// 从1开始, 假定0位置有序
	for i := 1; i < n; i++ {
		cur := arr[i]
		j := i
		// 0 ~ j 的位置已经排好序
		for ; j >= 1 && arr[j-1] > cur; j-- {
			// 后移一位
			arr[j] = arr[j-1]
		}
		arr[j] = cur
	}
}

func main() {
	arr := []int{5, 2, 3, 1, 4, 20, 10, 19, 23, 11, 25, 98, 23, 3, 4, 123}
	InsertionSort(arr)
	fmt.Println(arr)
}
