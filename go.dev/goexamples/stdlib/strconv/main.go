package main

import (
	"fmt"
	"strconv"
)

func main() {
	// strconv 字符串转 基础类型
	// 第二个0  是自动判断字符串的进制  第三个参数指定 "100" bit位数
	i, _ := strconv.ParseInt("100", 10, 64)
	fmt.Println(i)

	bs, _ := strconv.ParseInt("11111000", 2, 8)
	fmt.Println(bs)

	f, _ := strconv.ParseFloat("99.6", 64)
	fmt.Println(f)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	b, _ := strconv.ParseBool("true")
	fmt.Println(b)

}
