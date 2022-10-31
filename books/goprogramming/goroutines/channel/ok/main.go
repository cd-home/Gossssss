package main

import "fmt"

func main() {
	// ok 模式读取通道
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	for {
		// 读取完成后，如果通道关闭, ok 返回false
		// ok 为 false 时，表示通道已经关闭，无值可接受
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v, ok)
	}
}
