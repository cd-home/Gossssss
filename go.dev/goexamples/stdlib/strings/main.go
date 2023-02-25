package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello World"
	// 大量字符串处理函数都在strings包
	fmt.Println(strings.HasPrefix(s, "He"))
	fmt.Println(strings.Split(s, "o"))
	fmt.Println(strings.Fields(s))
	fmt.Println(strings.Cut("127.0.0.1:8080", ":"))
	fmt.Println(strings.TrimPrefix(s, "H"))
}
