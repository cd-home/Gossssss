package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 10
	// 如果你需要修改, 那么必须是指针
	rva := reflect.ValueOf(&a)
	fmt.Println(rva.Kind())
	rva.Elem().SetInt(20)
	fmt.Println(a)

	var b bool = true
	rvb := reflect.ValueOf(&b)
	fmt.Println(rvb.Kind())
	rvb.Elem().SetBool(false)
	fmt.Println(b)

	var s string = "Hello"
	rvs := reflect.ValueOf(&s)
	rvs.Elem().SetString("World")
	fmt.Println(s)
	elem := rvs.Elem()
	for i := 0; i < elem.Len(); i++ {
		fmt.Println(elem.Index(i))
	}
}
