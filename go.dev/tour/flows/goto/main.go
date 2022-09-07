package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i == 10 {
			// 强制流程
			goto FLAG
		}
	}
FLAG:
	fmt.Println(111)
}
