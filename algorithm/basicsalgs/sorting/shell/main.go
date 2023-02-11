package main

import "fmt"

// ShellSort 希尔排序, 优化插入排序
func ShellSort(arr []int) {
	n := len(arr)
	// increment 一直到1
	for increment := n / 2; increment > 0; increment /= 2 {
		// 插入排序
		for i := increment; i < n; i++ {
			tmp := arr[i]
			j := i
			// 比较移位
			for j = i; j >= increment && arr[j-increment] > tmp; j -= increment {
				arr[j] = arr[j-increment]
			}
			// 赋值
			arr[j] = tmp
		}
	}
}

func main() {
	arr := []int{5, 2, 3, 1, 4, 20, 10, 19, 23, 11, 25, 98, 23, 3, 4, 123}
	ShellSort(arr)
	fmt.Println(arr)
}
