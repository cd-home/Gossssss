package main

import "fmt"

func main() {
	// map 是无序的键值对, map是hash表的引用
	// k - v k的类型必须是可比较的,并且相同的, v 可以是相同的任意类型

	// nil map 无法直接使用, 必须初始化
	// map 的零值就是 nil
	// map 之间无法比较, map只能与nil比较
	var m map[string]string

	// 必须初始化, 字面量/make, 通常情况下使用make
	m = map[string]string{}

	// set
	m["name"] = "yao"

	fmt.Println(m)

	m1 := map[string]any{
		"name": "yao",
		"age":  "28",
	}
	fmt.Println(m1)

	age := m1["age"]
	fmt.Println(age)

	// 如果key不存在，那么将得到value对应类型的零值
	x := m1["no exist key"]
	fmt.Println(x)

	// 注意 value 是不可寻址的, 不是变量无法寻址
	// 禁止对map元素取址的原因是map可能随着元素数量的增长而重新分配更大的内存空间
	// 从而可能导致之前的地址无效
	//p := &m1["name"]
}
