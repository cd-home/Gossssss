package main

import "fmt"

func main() {
	// 创建无缓存通道, 必须有两端准备好
	ch := make(chan int)

	go func() {
		// 往通道中写
		ch <- 100
	}()

	// 阻塞等待通道数据，读数据
	fmt.Println(<-ch)


	bufferCh := make(chan string, 2)

	// 无需等接收方，可以先缓存
	bufferCh <- "Hello"
	bufferCh <- "World"

	fmt.Println(<-bufferCh)
	fmt.Println(<-bufferCh)
}
