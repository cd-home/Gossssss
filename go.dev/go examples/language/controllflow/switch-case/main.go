package main

import (
	"fmt"
	"time"
)

func main() {
	// 基础的switch
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// case 支持多个表达式值
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	// switch省略表达式值 case 支持逻辑表达式, 相当于if/else
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("before noon")
	default:
		fmt.Println("after noon")
	}

	// switch 类型选择
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Printf("Do not Know: %T", t)
		}
	}
	whatAmI(true)
	whatAmI(123)
	whatAmI("name")

	// switch 支持变量声明
	f := func() int {
		return 10
	}
	switch i := f(); i {
		
	}
}
