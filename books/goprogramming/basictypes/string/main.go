package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	/*
		字符串
			字符串的值是不可变的：一个字符串包含的字节序列永远不会被改变
			索引访问
			len 获取长度
			支持 +
			特殊语意
				\n \t \f \r
		原始字符串: 反引号
			``
		ASCII: 128
		Unicode: 每个符号都分配一个唯一的Unicode码点，Unicode码点对应Go语言中的rune整数类型
			UTF-8：编码规则 1-4个字节来编码unicode
		Go语言源文件采用UTF-8编码

		字符串处理
			strings、bytes、
			strconv 字符串转换为数字
			unicode
	*/
	var s string = "hello world"
	fmt.Println(s[1])
	fmt.Println(s[2:6])
	fmt.Println(s[:])

	// unicode
	s = "Hello, 世界"
	fmt.Println(len(s))                    // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"
	// for-range 自动解码
	for _, char := range s {
		fmt.Println(string(char))
	}

	// string => bytes
	bs := []byte(s) // 分配一个新的字节数组用于保存字符串数据的拷贝
	fmt.Println(bs)
}
