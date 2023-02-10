package main

import "fmt"

func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		cur := arr[i]
		j := i - 1
		// 0 ~ j 的位置已经排好序
		for ; j >= 0 && arr[j] > cur; j-- {
			// 后移一位
			arr[j+1] = arr[j]
		}
		arr[j+1] = cur
	}
}

func main() {
	arr := []int{5, 2, 3, 1, 4}
	InsertionSort(arr)
	fmt.Println(arr)
}
