package main

import (
	_ "Gossssss/go.dev/tour/functions/init/from"
	"fmt"
)

// 在main之前运行
func init() {
	fmt.Println("main init")
}

func main() {
	fmt.Println("main run")
}
