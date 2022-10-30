package main

import "fmt"

type User struct {
	_    struct{}
	Name string
	Age  int
}

func main() {
	// 如果需要显示命名赋值
	u := User{Age: 28, Name: "yao"}
	fmt.Println(u)
}
