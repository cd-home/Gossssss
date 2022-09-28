package main

import (
	"fmt"
	"reflect"
)

func main() {
	foo := func(a, b int) int {
		return a + b
	}
	rvf := reflect.ValueOf(foo)
	rtf := reflect.TypeOf(foo)

	v := rvf.Call([]reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)})

	fmt.Println(v[0])

	// 参数
	fmt.Println(rtf.NumIn())
	// 返回值
	fmt.Println(rtf.NumOut())
	// out type
	fmt.Println(rtf.Out(0))
	fmt.Println(rtf.In(0))
	fmt.Println(rtf.In(1))
}
