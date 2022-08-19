package main

import "fmt"

// 同类型参数
func plus(a, b int) (c int) {
	c = a + b
	return
}

func SayHello(name string, year int) {
	fmt.Printf("Hello %s, %d\n", name, year)
}

func GetUser() (string, int) {
	return "God Yao", 25
}

func main() {
	fmt.Println(plus(1, 2))
	SayHello("God Yao", 2021)

	name, age := GetUser()
	fmt.Println(name)
	fmt.Println(age)

	// 空白标示符
	name2, _ := GetUser()
	fmt.Println(name2)

	fmt.Println(Sum(1, 2, 3, 4, 5))

	// 传递slice, 需要解开slice
	sli := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(Sum(sli...))


	nextSeq := IntSeq()
	fmt.Println(nextSeq())
	fmt.Println(nextSeq())
	fmt.Println(nextSeq())
	fmt.Println(nextSeq())

	fmt.Println("+++++++++++++++++")
	// 0 1 1 2 3 5 8 13
	fmt.Println(Fib(6))
}

func Sum(nums ...int) (r int) {
	for _, num := range nums {
		r += num
	}
	return
}

// 函数作为返回值 闭包
func IntSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
// 递归
func Fib(n int) int {
	// 退出条件
	if n < 2 {
		return n
	}
	// 递归公式
	return Fib(n-1) + Fib(n-2)
}