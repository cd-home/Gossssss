package main

import "fmt"

func main() {
	// 声明数组, 必须指定长度
	// 默认初始化为类型零值
	var arr [5]int
	fmt.Println(arr)

	// index
	arr[1] = 10
	fmt.Println(arr)

	// len
	fmt.Println(len(arr))

	// 字面量形式
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	// 多维数组
	var twoD [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
