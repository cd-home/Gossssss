package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 对于普通类型的(简单类型) 反射修改值都有其对应的api
	var i int = 10
	rvi := reflect.ValueOf(&i)
	rvi.Elem().SetInt(20)
	fmt.Println(i)

	var b bool = true
	rvb := reflect.ValueOf(&b)
	rvb.Elem().SetBool(false)
	fmt.Println(b)

	var s string = "Hello World"
	rvs := reflect.ValueOf(&s)
	rvs.Elem().SetString(s + "!")
	fmt.Println(s)
}
