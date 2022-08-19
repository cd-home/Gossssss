package example_test

import (
	"fmt"
	"testing"
)

func GrowingStack(n int) {
	slices := make([]int, n)
	fmt.Println(slices)
}

func TestStackGrowth(t *testing.T) {
	ch := make(chan bool)
	input := 10
	go func(n int) {
		GrowingStack(n)
		ch <- true
	}(input)
	<-ch
}
