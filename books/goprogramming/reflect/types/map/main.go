package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := make(map[string]string, 10)
	m["name"] = "yao"
	m["age"] = "27"
	m["sex"] = "man"

	rv := reflect.ValueOf(&m)
	elem := rv.Elem()
	// 获取所有的keys
	fmt.Println(elem.MapKeys())

	// 通过key 获取值
	fmt.Println(elem.MapIndex(reflect.ValueOf("name")))

	// 迭代
	iter := elem.MapRange()
	for iter.Next() {
		fmt.Println(iter.Key())
		fmt.Println(iter.Value())
	}

	// 设置
	elem.SetMapIndex(reflect.ValueOf("name"), reflect.ValueOf("yoo"))
	fmt.Println(rv)
}
