package main

import (
	"fmt"
	"unsafe"
)

// Basic Type
// string bool
// int    int8   int16   int32   int64
// uint   uint8  uint16  uint32  uint64
// uintptr

// byte
// alias for uint8

// rune
// alias for uint32

// float32
// float64

// complex64
// complex128

const PI = 3.14
const Version = "v1.0.0"
const Switch = false
const CHAR = 'a'

const num = 2

func main() {
	fmt.Println("Hello World!")
	// 不同平台 不同位
	var a int = 10
	fmt.Println(unsafe.Sizeof(a))

	var b int16 = 50
	fmt.Println(unsafe.Sizeof(b))

	// Variables declared without an explicit initial value are given their zero value.
	var name string
	fmt.Println(name)

	var age uint8
	fmt.Println(age)

	var flag bool
	fmt.Println(flag)

	// The expression T(v) converts the value v to the type T.
	age = 18
	var f float32 = float32(age)
	fmt.Println(f)

	// When declaring a variable without specifying an explicit type
	// (either by using the := syntax or var = expression syntax),
	// the variable's type is inferred from the value on the right hand side.
	var j = 10
	fmt.Printf("%T\n", j)

	k := j
	fmt.Printf("%T\n", k)

	l := 3.14
	fmt.Printf("%T\n", l)

	// Constants
	fmt.Println(PI)
	fmt.Println(Version)
	fmt.Println(Switch)
	// CHAR ASCII
	fmt.Println(CHAR)

	//
	fmt.Println(DoubleNumInt(num))
	fmt.Println(DoubleNumFloat(num))
}

func DoubleNumInt(x int) int {
	return x * x
}

func DoubleNumFloat(x float32) float32 {
	return x * x
}