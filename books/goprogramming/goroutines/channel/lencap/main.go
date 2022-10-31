package main

import "fmt"

func main() {
	// cap 获取通道的容量
	// len 获取当前通道的数据
	ch := make(chan int, 10)
	go func() {
		defer close(ch)
		for i := 0; i < 20; i++ {
			fmt.Println(cap(ch), len(ch))
			ch <- i
		}
	}()
	for v := range ch {
		fmt.Println(v)
	}
}
