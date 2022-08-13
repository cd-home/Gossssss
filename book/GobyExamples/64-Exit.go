package main

import (
	"fmt"
	"os"
)

func main() {

	// 不会被执行
	defer fmt.Println("can not exec")

	os.Exit(1)
}
