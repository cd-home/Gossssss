package main

import "fmt"

func main() {
	// 函数可以作为变量赋值，可以作为函数参数、可以作为返回值
	// 函数的零值是nil
	f := Add
	v := f(1, 2)
	fmt.Println(v)

	// 函数作为参数
	fmt.Println(funcValues(Add, Del, 1, 2, 3, 4))

	fmt.Println("next.......")

	ff := next(0)
	fmt.Println(ff())
	fmt.Println(ff())
	fmt.Println(ff())
	fmt.Println(ff())
}

func Add(a, b int) int {
	return a + b
}

func Del(c, d int) int {
	return c - d
}

func funcValues(add func(int, int) int, del func(int, int) int, a, b, c, d int) int {
	return add(a, b) + del(c, d)
}

func next(x int) func() int {
	// 闭包, 函数作为返回值
	return func() int {
		// 匿名函数访问外部变量
		x++
		return x
	}
}
