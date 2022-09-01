package main

import "fmt"

// Index
// Go functions can be written to work on multiple types using type parameters
// comparable is a useful constraint that
// makes it possible to use the == and != operators on values of the type
func Index[T comparable](s []T, x T) int {
	for idx, value := range s {
		// T can comparable
		if x == value {
			return idx
		}
	}
	return -1
}

func main() {
	s1 := []int64{1, 3, 4, 2, 5}
	fmt.Println(Index(s1, 2))
	fmt.Println(Index[int64](s1, 2))
}
