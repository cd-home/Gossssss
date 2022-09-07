package main

import (
	"fmt"
	"math"
)

// IFControlSqrt
// x = v*v
func IFControlSqrt(x float64) string {
	if x < 0 {
		return IFControlSqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// IFControlPow
// x^n
// if 支持变量声明语句
func IFControlPow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// IFControlPow2
// x^n
func IFControlPow2(x, n, lim float64) float64 {
	// Go 只有 if - else
	if v := math.Pow(x, n); v < lim {
		return v
		// else
	} else {
		fmt.Printf("%f > %f", v, lim)
	}
	return lim
}

func IFControl(x int) int {
	if x < 10 {
		return x
	}
	return x % 10
}

func main() {
	// If
	fmt.Println(IFControl(10))
	fmt.Println(IFControl(12))

	fmt.Println(IFControlSqrt(-16))
	fmt.Println(IFControlPow(2, 4, 15))
}
