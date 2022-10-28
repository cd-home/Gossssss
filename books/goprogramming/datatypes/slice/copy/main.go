package main

import "fmt"

func main() {
	src := []int{4, 5, 6}
	// 分配目标内存
	dst := make([]int, len(src))

	// 赋值元素到新的内存空间
	copy(dst, src)
	fmt.Println(src)
	fmt.Println(dst)

	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	copy(s[3:], s[4:])
	fmt.Println(s[:len(s)-1])
	// 第一个子切片 s[3:] 4 5 6 7 8
	// 第二个子切片 s[4:] 5 6 7 8    // 复制上去, 干掉4, 但是最后会多一个, 所以需要 len(s)-1
}
