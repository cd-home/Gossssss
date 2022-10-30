package main

import "fmt"

type Employee struct {
	// 大写是导出的
	ID        int
	Name      string
	Address   string
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	mike := Employee{
		ID:        1001,
		Name:      "yao",
		Address:   "guangzhou",
		Position:  "1023",
		Salary:    13000,
		ManagerID: 10,
	}
	// dot get
	fmt.Println(mike.Address)

	// dot set
	mike.Salary = 14000

	// get address
	manager := &mike.ManagerID
	*manager = 9
	fmt.Println(mike)

	// pointer, 指针可以直接访问字段属性的
	p := &mike
	p.Salary = 15000 // 实际上 (*p).Salary = 15000
	fmt.Println(mike)

	// 字面量, 不用名字映射，顺序即可，对于小结构体较友好, 但一般使用命名的
	yao := Employee{10002, "yao", "", "", 15000, 90}
	fmt.Println(yao)

	// 如果结构体成员是可比较的, 那么结构体就是可比较的
	// 结构体是可以内嵌 其他结构体、 或者匿名结构体的(匿名成员)
	// 匿名的方式, 可简化获取匿名成员的属性 与 方法
}
