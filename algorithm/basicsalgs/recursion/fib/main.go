package main

import "fmt"

func main() {
	// 0, 1, 1, 2, 3, 5, 8
	fmt.Println(Fib(5))
}

// Fib 递归求斐波那契
func Fib(n int) int {
	// 递归退出条件
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	// 递归公式
	return Fib(n-1) + Fib(n-2)
}
