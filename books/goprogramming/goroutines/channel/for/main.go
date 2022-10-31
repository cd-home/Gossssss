package main

import "fmt"

func main() {
	// for-range 可读取通道数据
	// 关闭的channel，除了本身的数据读取完成，最后是读取零值
	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 20; i++ {
			ch <- i
			//time.Sleep(time.Second)
		}
		// 发送完成需要 close, 否则检测到无发送端(当前的协程执行完退出了), 出现死锁的情况
		close(ch)
	}()

	// for range 读取 channel, 阻塞；除非关闭channel
	// 读取关闭的结束
	for v := range ch {
		fmt.Println(v)
	}
}
