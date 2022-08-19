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
}
