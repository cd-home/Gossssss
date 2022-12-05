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
	//  通过reflect.Value
	// 方式1: Addr找到reflect.Value的指针, 然后调用Interface() 找到interface{}
	// 然后断言获取普通类型的指针
	y := 2
	d := reflect.ValueOf(&y).Elem()
	px := d.Addr().Interface().(*int)
	*px = 3
	fmt.Println(y)
	// 方式2: 直接通过提供的API
	fmt.Println(d.CanAddr(), d.CanSet())
	d.Set(reflect.ValueOf(4))
	fmt.Println(y)
}
