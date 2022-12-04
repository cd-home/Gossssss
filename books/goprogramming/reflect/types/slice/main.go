package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 对于切片而言: 获取index值、cap、len
	s := make([]int, 0, 10)
	s = append(s, 1, 2, 3, 4, 5)
	rv := reflect.ValueOf(s)
	// kind 类型
	if rv.Kind() == reflect.Slice {
		fmt.Println(rv.Len())
		fmt.Println(rv.Cap())
		for i := 0; i < rv.Len(); i++ {
			fmt.Println(rv.Index(i))
		}
	}
	// 要想修改值, 就需要传递指针
	rvp := reflect.ValueOf(&s)
	if rvp.Kind() == reflect.Pointer {
		// 返回实际指针指向的值, 才能进行操作
		elem := rvp.Elem()
		fmt.Println(elem.Len())
		fmt.Println(elem.Cap())
		for i := 0; i < elem.Len(); i++ {
			elem.Index(i).SetInt(int64(i + 2))
		}
		for i := 0; i < elem.Len(); i++ {
			fmt.Println(elem.Index(i))
		}
		// cap 不能小于len, 同时也不能大写cap
		elem.SetCap(8)
		fmt.Println(elem.Cap())
	}
}
