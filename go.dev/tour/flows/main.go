package main

import (
	"fmt"
	"math"
	"time"
)

func ForLoopBaseControl() {
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		sum += i
	}
	fmt.Println(sum)
}

func ForOtherStyleControl() {
	//  Init And Incr Can "Out"
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

func ForEverLoopControl() {
	for {
		fmt.Println()
		time.Sleep(time.Second)
	}
}

func ForRangeFlowControl(slices []string) {
	// Note: index, value
	for i, v := range slices {
		fmt.Println(i, v)
	}
}

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
func IFControlPow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// IFControlPow2
// x^n
func IFControlPow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
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

// Sqrt
// NewTon Sqrt
func Sqrt(x float64) float64 {
	z := 1.0
	COUNT := 10
	for i := 1; i < COUNT; i++ {
		z = (z + x/z) / 2
	}
	return z
}

func main() {
	// For-Range
	data := []string{"a", "b", "c", "d"}
	ForRangeFlowControl(data)

	// For
	ForWhileControl()

	// If
	fmt.Println(IFControl(10))
	fmt.Println(IFControl(12))

	fmt.Println(IFControlSqrt(-16))
	fmt.Println(IFControlPow(2, 4, 15))

	fmt.Println(Sqrt(6))
	fmt.Println(math.Sqrt(5))
}
