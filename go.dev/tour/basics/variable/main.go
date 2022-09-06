package main

import "fmt"

// 包内 全局变量

// DeclaresVariable
// Just Declares
var DeclaresVariable string

// SingleGlobalVar
// Package Scope Var
// Declaration And Initializers
// 类型推荐省略, 由自动推断完成
var SingleGlobalVar string = "Hello World"

// 声明并且初始化, 支持多个, 并且类型可不同
// Type Assets Meantime Diff Type
var a, b, c = 10, 20, true

// Group
var (
	x string = "Go"
	y string = "Python"
	z string = "JavaScript"
)

func main() {
	// Default Zero Value
	fmt.Println(DeclaresVariable)

	fmt.Println(SingleGlobalVar)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// 函数作用域下变量
	// Function Scope Var
	var name string = "Yao"

	fmt.Println(name)

	// 函数内 短变量模式
	// 声明加初始化
	// 类型推断
	// Short variable declarations Just In functions
	age := "18"
	fmt.Println(age)

	// 支持多个, 而且类型可不同
	length, sex := 18, "man"
	fmt.Println(length)
	fmt.Println(sex)
}
