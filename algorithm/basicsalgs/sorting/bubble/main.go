package main

import "fmt"

// BubbleSort
// 冒泡排序
func BubbleSort(arr []int) {
	l := len(arr)
	// 检查入参
	if l <= 1 {
		return
	}
	// 轮数
	for i := 0; i < l; i++ {
		// 每一轮进来定义一个是否有数据交换的标志
		flag := false
		// 每一轮做什么, 一个数和剩下的比较
		for j := 0; j < l-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
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
