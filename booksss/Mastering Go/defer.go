package main

import "fmt"

func main() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i)
	}
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i)
		}()
	}
	for i := 3; i > 0; i-- {
		defer func(i int) {
			fmt.Print(i)
		}(i)
	}
}
