package main

import "fmt"

func main() {
	var a int = 10
	fmt.Println(float64(a))

	var b float64 = 10.10
	fmt.Println(int64(b))

	var bs []byte = []byte("Hello world")
	fmt.Println(string(bs))

	// .... other type int int16 int...
	// .... other complex64 complex128
}
