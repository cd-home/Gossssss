package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		c1 <- "One"
	}()

	go func() {
		time.Sleep(6 * time.Second)
		c2 <- "Two"
	}()

	// 注意如果time.After放在select里面，每次select都会重置
	c := time.After(4 * time.Second)

	// select 可以同时等待多个channel的操作
	// 通常是读、也可以写
	for {
		select {
		case v1 := <-c1:
			fmt.Println(v1)
		case v2 := <-c2:
			fmt.Println(v2)
		case <-c:
			fmt.Println("timeout")
			return
		}
	}
	// select 如果有多个case可以执行, 那么会随机选择一个
	// select 如果没有可执行的
	// 如果有default走default, 没有就会阻塞住

}

func NonBlocking() {
	messages := make(chan string)
	// 通常在某个时机做一次性检测
	select {
	case msg := <-messages:
		fmt.Println(msg)
	default:
		fmt.Println("Non-Blocking")
	}
}
