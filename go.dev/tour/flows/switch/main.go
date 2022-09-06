package main

import "fmt"

func main() {
	a, b, c, d := 1, 2, 3, 4
	number := 4
	// switch-case no need to break
	switch number {
	case a:
		fmt.Println("a == number")
	case b:
		fmt.Println("b == number")
	case c:
		fmt.Println("c == number")
	case d:
		fmt.Println("d == number")
	}
}
