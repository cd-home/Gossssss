package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	done := make(chan struct{})
	go func() {
		// range阻塞读取, 直到通道被close
		// 如果没有close, 会一直阻塞住
		for v := range ch {
			fmt.Println(v)
		}
		fmt.Println("read all")
		done <- struct{}{}
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		ch <- i
	}
	// 发送完成根据需要的话(如果是一直发送数据), (否则)可以进行关闭
	// 一个非空的通道也是可以关闭的, 并且，通道中剩下的值仍然可以被接收到
	close(ch)
	<-done
}
