package main

import (
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func main() {
	fmt.Println("Hello World")

	p := Point{
		x: 1,
		y: 2,
	}

	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)

	fmt.Printf("%#v\n", p)

	fmt.Printf("%p\n", &p)

	fmt.Printf("%T\n", p)

	fmt.Printf("%t\n", true)

	fmt.Printf("%b\n", []byte("2"))

	fmt.Printf("%.2f\n", 22.677)

	fmt.Printf("%c\n", 97)

	fmt.Printf("%x\n", 123)

	fmt.Printf("%e\n", 12000000000.00)
	fmt.Printf("%E\n", 12000000000.00)

	fmt.Printf("%s\n", "Hello")

	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	_, _ = fmt.Fprintf(os.Stderr, "Error")
}
