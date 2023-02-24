package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// go 中的字符串是 只读的字节序列
	// 默认是 UTF-8 编码的
	// go 中字符的概念是 rune == unicode code point
	s := "Hello World; 你好，世界"
	// len 返回的是字节数, 一个中文三个字节
	fmt.Println(len(s))
	// range 按照 rune 解码
	for _, char := range s {
		// char 是 unicode
		fmt.Println(char, string(char))
	}
	// 多少个rune
	fmt.Println(utf8.RuneCountInString(s))
}
