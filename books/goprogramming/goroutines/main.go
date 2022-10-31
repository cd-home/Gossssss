package main

import (
	"fmt"
)

func main() {
	// close channel
	// 可以关闭双向，写 channel， 不可以关闭只读channel和已经关闭的channel
	// 双向可以转单向channel, 单向不可以转双向

	// 无缓冲的channel必须发送、接受同步准备
	// 有缓冲的channel可以异步, 队列满，发送阻塞，队列空，接受阻塞
	//ch := make(<-chan int)
	//close(ch)

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

	// ok 读取的方式, 可以ok模式
	v, ok := <-ch
	fmt.Println(v, ok)
}
