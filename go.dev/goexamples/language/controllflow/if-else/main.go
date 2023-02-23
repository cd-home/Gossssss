package main

import "fmt"

/*
	if 条件不需要括号
	条件前面可以声明一个变量
	if {

	} else if {

	} else {

	}
	Go没有三目运算符 只能使用以上的模式
*/

func main() {
	// 单个if
	if true {
	}

	// if-else
	if 7%2 == 0 {
		fmt.Println("7 is Even")
	} else {
		fmt.Println("7 is Odd")
	}

	// 可以在if中声明一个变量, 该变量的作用域仅仅限于当前分支 与 后续分支
	if num := 9; num < 10 {
		fmt.Println("num < 10")
	}
}
