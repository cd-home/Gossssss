package main

import "fmt"

func main() {
	p := Person{
		Name: "GodYao",
		Age:  25,
	}
	fmt.Println(p)

	p.Name = "Jack"

	fmt.Println(p)

	// 自动解引用
	pp := &p
	fmt.Println(pp.Age)
	pp.Age = 27

	fmt.Println(*pp)
}

type Person struct {
	Name string
	Age  int
}
