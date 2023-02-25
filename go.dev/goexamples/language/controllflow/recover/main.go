package main

import (
	"fmt"
	"os"
)

func main() {
	// 捕获panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	_, err := os.Open("name.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("do not run this")
}
