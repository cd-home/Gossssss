package main

import "fmt"

func main() {
	// 注册延迟函数，等待函数主体逻辑执行完成后执行
	defer fmt.Println("Done")

	fmt.Println("start")
}
