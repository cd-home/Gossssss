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
	// Elem() 获取原始对象
	// 表示获取原始值对应的反射对象，只有原始对象才能修改，当前反射对象是不能修改的
	// 如果要修改反射类型对象，其值必须是addressable
	// 对应的要传入的是指针，同时要通过Elem方法获取原始值对应的反射对象
	fmt.Println(rva.Elem().CanSet())
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
	// Interface returns v's current value as an interface{}.
	rvs = reflect.ValueOf(s)
	fmt.Println(rvs.Interface())

	// 更加的友好
	//reflect.Indirect()
}
