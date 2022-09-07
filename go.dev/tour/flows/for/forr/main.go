package main

import (
	"fmt"
)

// ForRangeFlowList
// array, slice,
func ForRangeFlowList(slices []string) {
	// Note: index, value
	for i, v := range slices {
		fmt.Println(i, v)
	}
}

func ForRangeFlowMap[Key comparable](m map[Key]any) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func ForRangeFlowChannel() {
	ch := make(chan int, 3)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		// close(ch)
	}()
	v, ok := <-ch
	fmt.Println(v, ok)
	v, ok = <-ch
	fmt.Println(v, ok)
	v, ok = <-ch
	fmt.Println(v, ok)
	v, ok = <-ch
	fmt.Println(v, ok)
	v, ok = <-ch
	fmt.Println(v, ok)
	v, ok = <-ch
	fmt.Println(v, ok)
}

func main() {
	data := []string{"a", "b", "c", "d"}
	ForRangeFlowList(data)

	m := make(map[string]any)
	m["name"] = "yao"
	m["age"] = 28

	ForRangeFlowMap(m)
}
