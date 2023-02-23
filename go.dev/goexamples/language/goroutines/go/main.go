package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
// 并发的方式运行g
func main() {
	// sync
	f("sync")

	// async
	go f("async")


	// 匿名 async
	go func(msg string) {
		fmt.Println(msg)
	}("Golang")

	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}
