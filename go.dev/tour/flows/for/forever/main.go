package main

import (
	"fmt"
	"time"
)

// ForEverLoopControl
// 通常不停的app, 如 TCP Accept
// 或者配合 select 读取 channel
func ForEverLoopControl() {
	// omit the loop condition it loops forever
	// 忽略循环的三个部分, 即是无限循环
	for {
		fmt.Println()
		time.Sleep(time.Second)
	}
}

func ForSelectChannel() {
	ch := make(chan int)
	// 写
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second * 1)
		}
		close(ch) // 发送完成关闭channel
	}()
	// 读
	for {
		// 监听channel数据
		select {
		case data, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(data)
		default:
			fmt.Println("wait data")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ForSelectChannel()
}
