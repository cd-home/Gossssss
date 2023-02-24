package main

import "fmt"

type describer interface {
	describe() string
}

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	// 内嵌匿名结构体, 也可以具名
	base
	str string
}

func main() {
	co := container{
		// 默认使用其类型名字
		base: base{num: 1},
		str:  "some name",
	}
	// 匿名结构体好处可以直接访问其字段
	fmt.Println(co.num, co.str)
	// 使用类型名字去方法亦可
	fmt.Println(co.base.num)
	// 内嵌结构体的方法也变成了容器的方法
	fmt.Println(co.describe())

	// 既然容器实现了其方法, 那么就能够赋值接口变量
	var d describer = co
	fmt.Println(d.describe())
}
