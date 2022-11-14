package main

import "fmt"

func main() {
	// 支持多返回值, 并且不需要的变量可以用空白标识符丢弃
	v, _ := de(10, 5)
	fmt.Println(v)
}

func de(a, b int) (int, error) {
	return a / b, nil
}
