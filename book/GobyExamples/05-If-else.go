package main

import "fmt"

func main() {
	if 7 % 2 == 0 {
		fmt.Println("7 is Even")
	} else {
		fmt.Println("7 is Odd")
	}

	// 可以在if中声明一个变量
	if num := 9; num < 10 {
		fmt.Println("num < 10")
	}

	// Go没有三目运算符
}
