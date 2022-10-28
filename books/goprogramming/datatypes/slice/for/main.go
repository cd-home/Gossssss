package main

import "fmt"

func main() {
	var s = []string{"a", "b", "c", "d", "e", "f", "g"}

	// 增量迭代循环
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	// range 循环
	// index, item
	for i, v := range s {
		fmt.Println(i, v)
	}
}
