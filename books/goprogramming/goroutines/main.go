package main

import (
	"time"
)

func main() {
	// goroutine 是并发基本单元
	// channel 是 g 之间的通信机制
	// 对于开启的g, 需要知道是否需要退出, 能否正常退出
	// 泄漏的goroutines并不会被自动回收，因此确保每个不再需要的goroutine能正常退出是重要的
	go func() {

	}()
	time.Sleep(time.Second)
}
