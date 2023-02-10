package main

import "fmt"

func SelectionSort(arr []int) {
	n := len(arr)
	// n 个 元素 只需要 确定 n-1 个元素大小
	for i := 0; i < n-1; i++ {
		min := i
		// i 与 i + 1 后面大小比较
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				// 一直找到最小的
				min = j
			}
		}
		// 交换 exchange
		arr[i], arr[min] = arr[min], arr[i]
	}
}

func main() {
	arr := []int{5, 2, 3, 1, 4}
	SelectionSort(arr)
	fmt.Println(arr)
}
