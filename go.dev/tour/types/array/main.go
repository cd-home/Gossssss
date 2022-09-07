package main

import "fmt"

// The type [n]T is an array of n values of type T.
// n 必须是明确的
// 实际上 数组不常用, 可使用更加灵活的切片

func main() {
	var arrays [5]int
	arrays[0] = 1
	arrays[1] = 2
	arrays[2] = 3
	arrays[3] = 4
	arrays[4] = 5
	fmt.Println(arrays)

	// by index
	arraysIndex := [3]int{1: 2, 2: 3}
	fmt.Println(arraysIndex)

	// len 可获取长度
	fmt.Println(len(arrays))
}
