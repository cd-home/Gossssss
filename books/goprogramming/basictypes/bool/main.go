package main

import "fmt"

func main() {
	/*
		布尔类型
			true、false
			支持 && || !三种逻辑运算
	*/
	var b1 = true
	var b2 = false
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b1 && b2)
	fmt.Println(b1 || b2)
	fmt.Println(!b1)
}
