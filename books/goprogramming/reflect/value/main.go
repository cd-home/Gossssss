package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 10
	// 依赖 interface{} == any; 返回状态操作数动态值的 reflect.Value类型
	rv := reflect.ValueOf(i)
	fmt.Println(rv)
	fmt.Println(rv.String())

	// 转换为 reflect.Type类型
	rt := rv.Type()
	fmt.Println(rt.String())

	// 逆操作 reflect.Value => interface{}==any
	x := rv.Interface()
	fmt.Println(x.(int))

	/*
		reflect.Value与interface{}都能装载任意值;
		所不同的是: interface{} 隐藏了值的内部表示方法, 以及所有的方法
		要想知道具体类型的动态类型以及值只能使用断言.
		reflect.Value提供了很多API来检测内容.
	*/
}
