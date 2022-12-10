package main

import "fmt"

func SelectionSort(arr []int) {
	n := len(arr)
	// n 个元素， n-1次循环即可
	for i := 0; i < n-1; i++ {
		min := i
		// 找最小
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

func main() {
	arr := []int{5, 2, 3, 1, 4}
	SelectionSort(arr)
	fmt.Println(arr)
}
