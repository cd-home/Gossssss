package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// NewPerson ( 称为构造函数、或者constructs
func NewPerson(name string, age int) *Person {
	// 可以安全的返回指针
	return &Person{
		Name: name,
		Age:  age,
	}
}

func main() {
	// 结构体是组织字段集合
	p := Person{
		Name: "GodYao",
		Age:  25,
	}
	fmt.Println(p)

	// dot . 语法访问、设置字段值
	p.Name = "Jack"

	fmt.Println(p)

	// 指针可以使用 dot . 自动解引用
	pp := &p
	fmt.Println(pp.Age)
	pp.Age = 27

	fmt.Println(*pp)

	// 忽略字段, 默认是类型的零值
	p2 := Person{
		Name: "jack ma",
	}
	fmt.Println(p2)
}
