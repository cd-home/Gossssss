package main

import "fmt"

func main() {
	/*
		常量
			编译期决定，无法修改
			通常是：字符串、布尔值、数字类型
			常量的操作返回的也是常量(算术运算、逻辑运算、比较运算、len等)
			常量可以指定类型, 如果没有指定, 会推断
			无类型常量: 可以延迟确定类型, 可以提高精度
		iota
			常量生成器,生成一组以相似规则初始化的常量
	*/
	const IPV4 = 4

	const (
		Sunday int = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	const (
		_ = 1 << (10 * iota)
		KB
		MB
		GB
		TB
	)
	fmt.Println(KB)
}
