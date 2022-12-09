package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// 任意二进制到文本字符串的编码方法，常用于在URL、Cookie、网页中传输少量二进制数据
	data := "I need base64"
	s := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(s)

	bytes, _ := base64.StdEncoding.DecodeString(s)
	fmt.Println(string(bytes))
}
