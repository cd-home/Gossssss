package main

import "fmt"

// BubbleSort
// 冒泡排序
func BubbleSort(arr []int) {
	n := len(arr)
	// 轮数
	for i := 0; i < n; i++ {
		// 每一轮进来定义一个是否有数据交换的标志
		flag := false
		// 每一轮比较交换, 这一轮能得到一个最值
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		// 如果没有交换说明已经达到有序的状态
		if !flag {
			break
		}
	}
}

func main() {
	s := []int{5, 2, 3, 1, 4}
	BubbleSort(s)
	fmt.Println(s)
}
