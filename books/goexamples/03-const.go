package main

import "fmt"

func main() {
	const PI = 3.14
	fmt.Println(PI)

	// 常量地址不可变
	// 可以声明确定的表达式
	const Value = 1 << 2

	fmt.Println(Value)

	// n并没有确定数值的具体类型，可以根据需要自动确定类型
	const n = 1000
	fmt.Println(n)
}
