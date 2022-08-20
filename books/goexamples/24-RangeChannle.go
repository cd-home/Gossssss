package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)

	go func() {
		// range阻塞读取
		for v := range ch {
			fmt.Println(v)
		}
	}()

	for i := 0; i < 10; i++ {
		ch <- i
	}
	// 一个非空的通道也是可以关闭的， 并且，通道中剩下的值仍然可以被接收到
	close(ch)
}