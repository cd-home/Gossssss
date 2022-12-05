package main

import "fmt"

// Foo
// 声明定义函数; 具体见函数;
// 函数名、参数列表、返回值列表、return
func Foo() {
	return
}

// Bar
// 自定义类型
type Bar struct {
	// 字段
	F1 string
	F2 int
}

func main() {
	// variable
	// 声明变量
	// 后置类型声明, 可以省略; 自动推导
	var a int = 10

	// 短变量
	b := 20
	fmt.Println(a)
	fmt.Println(b)

	// 支持多个
	i, j := 1, 2
	j, i = i, j
	fmt.Println(i, j)
	
	// constant
	// 常量
	const P = 3.15
}
