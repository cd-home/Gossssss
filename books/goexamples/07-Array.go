package main

import "fmt"

func main() {
	var arr [5]int

	// index
	arr[1] = 10
	fmt.Println(arr)

	// len
	fmt.Println(len(arr))

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
