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
	// sync 同步执行
	f("sync")

	// async 异步执行
	// goroutine 是 轻量级的"线程"
	go f("async")

	// 匿名 async
	go func(msg string) {
		fmt.Println(msg)
	}("Golang")

	// 必须等到其执行, 更好的办法是采用 WaitGroup
	// goroutines 之间是同时执行的 (单核并发、多核可能是并行的)
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}
