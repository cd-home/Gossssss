package main

import (
	"Gossssss/go.dev/tour/basics/const/iotas"
	"fmt"
)

// Constants can be character, string, boolean, or numeric values.
// 常量只能是 字符、字符串、布尔值、数值类型

const PI = 3.14
const Version = "v1.0.0"
const Switch = false
const CHAR = 'a'
const num = 2
const PI2 = 2 * PI

// const groups
const (
	name = "yao"
	age  = 28
)

// An untyped constant takes the type needed by its context.
const (
	// Big
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Small
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

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

	// Numeric constants are high-precision values.
	// An untyped constant takes the type needed by its context.
	// 未声明类型的数值常量获取由上下文决定的
	// 未声明类型的数值常量是高精度的256位
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func DoubleNumInt(x int) int {
	return x * x
}

func DoubleNumFloat(x float32) float32 {
	return x * x
}
