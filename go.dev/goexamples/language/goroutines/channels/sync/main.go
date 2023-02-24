package main

import (
	"fmt"
	"time"
)

func Work(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
	// 执行完成写入通道, 相当于一种通知的机制
	done <- true
}

func main() {
	done := make(chan bool)

	go Work(done)

	// 同步阻塞等待
	<-done
}
