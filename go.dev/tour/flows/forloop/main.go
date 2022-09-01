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

	fmt.Println(Sqrt(6))
	fmt.Println(math.Sqrt(5))
}
