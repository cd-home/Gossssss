package main

import "fmt"

func main() {
	var arr = [...]string{"a", "b", "c", "d", "e", "f", "g"}
	s1 := arr[3:5]
	s2 := arr[4:5]
	// len 是当前有的元素, cap 是切片第一个元素到数组最后位置
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))

	// nil slice
	// nil 切片可以append, 但是无法进行其他操作
	var s []int
	// 切片只能和nil比较, 判断一个切片是否是空, 应该使用 len(s) == 0, 而不是 s == nil [有 s = []int{}, 非nil 切片, 但是是为空的]
	// nil 切片的 len = 0 cap = 0
	if len(s) == 0 {
		fmt.Println("s is empty")
	}

	// 除了特殊情况下, 一般情况下我们可以同等的对待长度为0的切片和nil切片

	// 实际上 s 是empty, 所以必须使用len 来判断是否是空
	s = []int{}
	if s == nil {
		fmt.Println("s is empty")
	} else {
		fmt.Println("s is not empty")
	}
}
