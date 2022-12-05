package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int = 10
	// pointer
	rv := reflect.ValueOf(&i)
	// 通过 Elem 获取 指针指向的变量
	elem := rv.Elem()
	elem.Set(reflect.ValueOf(20))
	fmt.Println(i)
}
