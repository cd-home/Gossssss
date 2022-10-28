package main

import "fmt"

func main() {
	m := make(map[string]any)
	m["name"] = "yao"
	m["age"] = 28
	m["sex"] = 0

	fmt.Println(m["name"])
	// 获取操作也是安全的
	fmt.Println(m["no exist key"])

	// ok 模式通常来判断获取v, 如果key不存在，那么将得到value对应类型的零值
	if v, ok := m["no exist key"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("no exist key")
		fmt.Println(v)
	}
}
