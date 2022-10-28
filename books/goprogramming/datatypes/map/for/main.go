package main

import "fmt"

func main() {
	m := make(map[string]any)
	m["name"] = "yao"
	m["age"] = 28
	m["sex"] = 0

	// map 是无序的， 遍历亦是无序的
	for k, v := range m {
		fmt.Println(k, v)
	}
}
