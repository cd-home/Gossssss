package main

import "fmt"

func main() {
	// channel 是连接并发goroutines 的管道, 用于goroutines 之间的通信、数据传输

	// 创建通道
	// 创建无缓存通道, 必须有两端准备好
	messages := make(chan string)

	go func() {
		// 写入操作
		messages <- "Hello"
	}()

	// 读取操作
	// 由于是无缓冲的通道, 会阻塞等待知道发送端、接收端准备好
	// 是一种靠通道的同步机制
	msg := <-messages
	fmt.Println(msg)

	// 缓冲通道
	bufferCh := make(chan string, 2)

	// 无需等接收方，可以先缓存
	bufferCh <- "Hello"
	bufferCh <- "World"

	fmt.Println(<-bufferCh)
	fmt.Println(<-bufferCh)
}
