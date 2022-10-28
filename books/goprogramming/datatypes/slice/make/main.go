package main

import "fmt"

func main() {
	// make 创建匿名的底层数组, 返回切片引用该数组
	//s := make([]int, len)
	//s := make([]int, len, cap)

	// s1 就是完全的底层数组的 view
	s1 := make([]int, 5)
	fmt.Println(s1, len(s1), cap(s1))

	// s2 包含前5个元素, 但是容量是整个底层的数组, 预留未来元素增长的
	// 通常在编码的时候, 如果可以预想到未来进行内存分配, 可以采用预留的方式, 避免多次内存分配
	s2 := make([]int, 5, 10)
	fmt.Println(s2, len(s2), cap(s2))
}
