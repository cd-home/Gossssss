package main

import "fmt"

func main() {
	var num = 100

	fmt.Println(num)

	// 可以一次声明多个同类型变量
	var a, b = 1, 2
	fmt.Println(a, b)

	var f = true
	fmt.Println(f)

	// 默认是对应类型的0值
	var e int
	fmt.Println(e)

	// short method
	// 函数内部 声明、并且初始化
	s := "God Yao"
	fmt.Println(s)
}
