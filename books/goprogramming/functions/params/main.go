package main

import "fmt"

func main() {
	fmt.Println(Sum(1, 2, 3, 4, 5))

	nums := []int{1, 2, 3, 4, 5}
	// 解包传递
	fmt.Println(Sum(nums...))
}

func Add(a int, b int) int {
	return a + b
}

func Add2(a, b int) int {
	return a + b
}

func Sum(nums ...int) int {
	res := 0
	// 不定参数行为像是切片一样
	for _, num := range nums {
		res += num
	}
	return res
}
