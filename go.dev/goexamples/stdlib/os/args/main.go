package main

import (
	"fmt"
	"os"
)

func main() {
	// 通常不用 args, 而是采用 parse 解析命令行参数
	// 包含执行文件
	argsAll := os.Args
	// 之后的参数
	argsLeft := os.Args[1:]
	fmt.Println(argsAll)
	fmt.Println(argsLeft)
}
