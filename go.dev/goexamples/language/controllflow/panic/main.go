package main

import "os"

func main() {
	// panic 的一种常见用法是：当函数返回我们不知道如何处理（或不想处理）的错误值时，中止操作
	// 或者程序必要的组件出现问题，无法正常运行
	_, err := os.Open("no exist")
	if err != nil {
		panic(err)
	}
}
