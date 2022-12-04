package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 10
	// 返回动态类型 reflect.Type
	rt := reflect.TypeOf(i)
	// 反射依赖接口: interface{} == any, 值 => 接口的隐式接口转换操作
	// 创建一个接口值：操作数的动态类型 和 操作数的值
	fmt.Println(rt.String())
}
