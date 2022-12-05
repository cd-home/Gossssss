package main

import "fmt"

func main() {
	/*
			1. 指针: 变量的地址
				&T 取地址
				*  解引用
				指针的零值是 nil
			2. 变量是可以寻址的
				接口体字段可以寻址

			3. 函数内分配的变量并不一定在栈上, 有可能逃逸到堆;
		       编译器会自动选择在栈上还是在堆上分配局部变量的存储空间
	*/
	// 可以省略类型, 为了清晰暂不
	var x int = 10
	// 可以看待 &x 的类型 *int 指针类型
	var p *int = &x
	fmt.Println(x)
	// *p 解引用, 读取变量的值
	fmt.Println(*p)
	// 修改
	*p = 20
	fmt.Println(x)

	// new(T) 函数可以 创建返回一个初始化类型T的零值内存地址
	m := new(int)
	*m = 99
	fmt.Println(*m)
}
