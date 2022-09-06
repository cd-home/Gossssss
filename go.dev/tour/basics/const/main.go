package main

import (
	"Gossssss/go.dev/tour/basics/const/iotas"
	"fmt"
)

const PI = 3.14
const Version = "v1.0.0"
const Switch = false
const CHAR = 'a'
const num = 2
const PI2 = 2 * PI

func main() {
	// Constants
	fmt.Println(PI)
	fmt.Println(Version)
	fmt.Println(Switch)
	// CHAR ASCII
	fmt.Println(CHAR)

	fmt.Println(DoubleNumInt(num))
	fmt.Println(DoubleNumFloat(num))

	fmt.Println(PI2)

	fmt.Println(iotas.Python)
	fmt.Println(iotas.Javascript)
	fmt.Println(iotas.Golang)
	fmt.Println(iotas.Erlang)
}

func DoubleNumInt(x int) int {
	return x * x
}

func DoubleNumFloat(x float32) float32 {
	return x * x
}
