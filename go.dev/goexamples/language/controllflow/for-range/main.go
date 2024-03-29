package main

import "fmt"

/*
	用于 序列(array slice string)、映射 遍历
	通道读取
*/

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(nums)

	// index, element
	for i, num := range nums {
		fmt.Println(i, num)
	}

	// key, value
	m := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// unicode 第一个返回值是字符的起始字节位置，然后第二个是字符本身
	for i, c := range "中国" {
		fmt.Println(i, c)
	}

	// index byte位置 rune==uint32
	for i, c := range "collection" {
		fmt.Println(i, c)
	}
}
