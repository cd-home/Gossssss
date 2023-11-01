package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	clear(s)
	fmt.Println(s)

	m := map[string]int{"a": 1}
	fmt.Println(m)
	clear(m)
	fmt.Println(m)
}
