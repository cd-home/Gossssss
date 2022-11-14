package main

import (
	"fmt"
	"os"
)

func main() {
	// 空白标识符号, 丢弃不需要的返回值
	f, _ := os.Open("path")
	buf := make([]byte, 1024)
	_, err := f.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
}
