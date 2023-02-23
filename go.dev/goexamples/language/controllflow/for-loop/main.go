package main

import "fmt"

/*
	go 唯一的循环控制结构
*/

func main() {
	// 类似 while
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// 经典的 init condition after
	for j := 0; j <= 9; j++ {
		fmt.Println(j)
	}

	// for-ever 死循环
	for {
		fmt.Println("loop")
		break
	}

	// break continue 用于跳出循环
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// 多个变量
	for i, j := 0, 10; i < j; i, j = i+1, j-1 {

	}
}
