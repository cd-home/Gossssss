package main

import (
	"fmt"
	"math"
)

func ForLoopBaseControl() {
	sum := 0
	// The init statement will often be a short variable declaration,
	// and the variables declared there are visible only in the scope of the for statement.
	// 初始化语句采用短变量声明, 并且作用域仅限for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		sum += i
	}
	fmt.Println(sum)
}

func ForOtherStyleControl() {
	//  Init And Incr Can be "Out"
	// 初始化语句与增量语句是可选的
	i := 0
	for i < 10 {
		i++
		fmt.Println(i)
	}
}

func ForWhileControl() {
	i := 10
	for i > 0 {
		fmt.Println(i)
		i--
	}
}

// Sqrt
// NewTon Sqrt
// 牛顿迭代法
func Sqrt(x float64) float64 {
	z := 1.0
	COUNT := 10
	for i := 1; i < COUNT; i++ {
		z = (z + x/z) / 2
	}
	return z
}

func main() {
	ForLoopBaseControl()
	fmt.Println(Sqrt(6))
	fmt.Println(math.Sqrt(5))

	ForOtherStyleControl()

	ForWhileControl()
}
