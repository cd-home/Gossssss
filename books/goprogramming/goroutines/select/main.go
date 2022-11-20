package main

import (
	"fmt"
	"time"
)

func main() {
	// select 多路复用; 监听case channel 的数据通信;
	// 多个case同时就绪时, select会随机地选择一个执行,
	// 这样来保证每一个channel都有平等的被select的机会
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case t := <-ticker.C:
			fmt.Println(t)
		default:
			fmt.Println("wait...")
			time.Sleep(time.Second)
		}
	}
}
