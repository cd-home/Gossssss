package main

import (
	"fmt"
	"reflect"
)

type TestSuite struct {
	a, b, c string
}

func main() {
	m := &TestSuite{"a", "b", "c"}
	n := &TestSuite{"a", "b", "c"}

	// 显示转 interface{}
	x := interface{}(m)
	y := interface{}(n)

	fmt.Println(m == n) // false
	fmt.Println(x == y) // false
	// 不转参数默认转
	fmt.Println(reflect.DeepEqual(m, n)) // true
	fmt.Println(reflect.DeepEqual(x, y)) // true
}
