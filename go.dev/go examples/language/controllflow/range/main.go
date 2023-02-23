package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(nums)

	for i, num := range nums {
		fmt.Println(i, num)
	}

	m := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// unicode 第一个返回值是字符的起始字节位置，然后第二个是字符本身
	for i, c := range "中国" {
		fmt.Println(i, c)
	}

	for i, c := range "collection" {
		fmt.Println(i, c)
	}
}
