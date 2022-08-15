package main

import "fmt"

func ForRangeFlowControl(slices []string) {
	// Note: index, value
	for i, v := range slices {
		fmt.Println(i, v)
	}
}

func main() {
	data := []string{"a", "b", "c", "d"}
	ForRangeFlowControl(data)
}
