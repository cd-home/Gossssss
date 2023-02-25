package main

import (
	"fmt"
	"os"
)

func main() {
	// 注册延迟函数，等待函数主体逻辑执行完成后执行
	defer fmt.Println("Done")
	fmt.Println("start")

	fd, err := os.Open("name.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()
}
