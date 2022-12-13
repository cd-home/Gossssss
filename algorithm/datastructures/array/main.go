package main

import "fmt"

func main() {
	// 数组(切片), 线性存储结构
	// 支持索引、遍历, 注意越界
	s := make([]int, 5)
	s[0] = 0
	s[1] = 1
	s[2] = 2
	s[3] = 3
	s[4] = 4
	for index, value := range s {
		fmt.Println(index, value)
	}
}
