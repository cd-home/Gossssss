package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

func main() {
	// 定义切片
	s := []int{1, 2, 3}

	// 范型函数
	fmt.Println(slices.Index(s, 1))

	s2 := []string{"a", "b", "c"}
	fmt.Println(slices.Index(s2, "b"))
}
