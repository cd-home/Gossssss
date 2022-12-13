package main

import "fmt"

func main() {
	// 映射 Map
	// 键值存储 k-v
	m := make(map[int]string, 5)
	m[1] = "one"
	m[2] = "two"
	m[3] = "three"
	for k, v := range m {
		fmt.Println(k, v)
	}
}
