package main

import (
	"fmt"
	"strconv"
)

func main() {
	// strconv 字符串转 基础类型
	// 第二个0  是自动判断进制
	i, _ := strconv.ParseInt("100", 0, 8)
	fmt.Println(i)

	f, _ := strconv.ParseFloat("99.6", 64)
	fmt.Println(f)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	b, _ := strconv.ParseBool("true")
	fmt.Println(b)
}
