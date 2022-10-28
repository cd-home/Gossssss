package main

import "fmt"

func main() {
	//m := make(map[string]string)

	// 知道map 大小的情况下，提前分配会更加友好
	//m := make(map[string]string, len)

	// 使用 map 必须先初始化, nil map 会 panic
	m := make(map[string]string)

	// set
	m["name"] = "yao"
	fmt.Println(m)
	// get
	name := m["name"]
	fmt.Println(name)
}
