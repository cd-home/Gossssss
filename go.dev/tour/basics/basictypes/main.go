package main

import (
	"fmt"
	"unsafe"
)

// Basic Type
// string
// bool
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

// groups variables
var (
	m = 1
	n = 2
	p = false
)

// zeroValue
// 类型零值 [基础类型零值, 复杂类型[map, channel, slice, pointer]零值nil]
func zeroValue() {
	// Variables declared without an explicit initial value are given their zero value.
	var name string
	fmt.Println(name)

	var age uint8
	fmt.Println(age)

	var flag bool
	fmt.Println(flag)
}

// typeConversions
// 类型转换
func typeConversions() {
	// The expression T(v) converts the value v to the type T.
	var age uint8
	age = 18
	var f float32 = float32(age)
	fmt.Println(f)
}

// typeInference
// 类型推断
func typeInference() {
	// When declaring a variable without specifying an explicit [明确的] type
	// (either by using the := syntax or var = expression syntax),
	// the variable's type is inferred from the value on the right hand side.
	// 类型推断
	var j = 10
	fmt.Printf("%T\n", j)

	k := j
	fmt.Printf("%T\n", k)

	l := 3.14
	fmt.Printf("%T\n", l)
}

func main() {
	// strings
	var s = "Hello World!"
	fmt.Println(s)

	// int, uint, uintptr 不同平台[32, 64] 不同位
	// When you need an integer value you should use int
	// unless you have a specific reason to use a sized or unsigned integer type
	// 大多数情况推荐直接使用int, 除非在对类型和大小特别明确的情况下
	var a int = 10
	fmt.Println(unsafe.Sizeof(a))

	var b int16 = 50
	fmt.Println(unsafe.Sizeof(b))

	// var groups
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(p)
}
