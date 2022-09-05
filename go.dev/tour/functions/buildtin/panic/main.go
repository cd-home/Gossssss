package main

import "fmt"

func main() {
	// 捕获到panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("force")
}
