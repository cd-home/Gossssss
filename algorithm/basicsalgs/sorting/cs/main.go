package main

import "fmt"

func main() {
	arr := []int{10, 7, 41, 9, 20, 5, 2, 4, 7, 3, 4}
	kvSort(arr)
}

func kvSort(arr []int) {
	// get max
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	// 设置计数的数组
	fg := make([]int, max+1)

	// 值大小，对于的index 位置计数
	for i := 0; i < len(arr); i++ {
		fg[arr[i]]++
	}

	for i, v := range fg {
		// 计数多少次输出多少次
		for j := 0; j < v; j++ {
			fmt.Println(i)
		}
	}
}
