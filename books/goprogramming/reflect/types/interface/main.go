package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Bar struct {
		Ex string
	}
	var i interface{} = &Bar{Ex: "Go"}
	rv := reflect.ValueOf(i)
	// 如果 变量是接口, 通过 Elem 来获取接口的动态值
	fmt.Println(rv.Elem())
}
