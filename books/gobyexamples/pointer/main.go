package main

import "fmt"
// 支持指针，允许通过引用传递 来 传递值和数据结构
// 函数都是值copy

func zeroVal(iVal int) {
	iVal = 0
}

// copy的是指针，所以解引用会修改真实地址的值
func zeroPtr(iPtr *int) {
	// 解引用
	*iPtr = 0
}

func main() {
	i := 1
	fmt.Println("init: ", i)

	zeroVal(i)
	fmt.Println("zeroValue: ", i)

	// 获取地址
	zeroPtr(&i)
	fmt.Println("zeroPtr: ", i)
}
