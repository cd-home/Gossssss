package main

import "fmt"

func main() {
	maxIntValue := max(0, 7, 6, 5, 4, 3, 2, 1) // 7 type int
	minIntValue := min(8, 7, 6, 5, 4, 3, 2, 1) // 1 type int
	fmt.Println(maxIntValue)
	fmt.Println(minIntValue)
}
