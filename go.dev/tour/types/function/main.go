package main

import (
	"fmt"
	"math"
)

// 函数是一种类型, 是值, 可以作为参数传递、返回值、可赋值给变量
// Function values
// Functions are values too. They can be passed around just like other values.
// Function values may be used as function arguments and return values.
func compute(fn func(float64, float64) float64, x, y float64) float64 {
	return fn(x, y)
}

// adder
// Go functions may be closures[闭包].
// A closure is a function value that references variables from outside its body.
// The function may access and assign to the referenced variables;
// in this sense the function is "bound" to the variables.
func adder() func(int) int {
	// outside variables
	sum := 0
	// closures bound above variable
	return func(x int) int {
		sum += x
		return sum
	}
}

// Fibonacci is a function that returns
// a function that returns an int.
func Fibonacci() func() int {
	x, y := -1, 1
	return func() int {
		x, y = y, x+y
		return y
	}
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot, 5, 12))
	fmt.Println(compute(math.Pow, 5, 12))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	f := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
