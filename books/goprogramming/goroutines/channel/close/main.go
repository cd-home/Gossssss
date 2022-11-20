package main

import "fmt"

func main() {
	// 通道是可以关闭的
	// 读写、写通道是可关闭； 只读通道无法关闭；已经关闭的通道无法关闭(panic)；

	// 可以关闭双向，写 channel， 不可以关闭只读channel和已经关闭的channel
	// 双向可以转单向channel, 单向不可以转双向

	// 通常要么是发送者自己关闭；要么该任务是生产-消费模型，不关闭
	ch := make(chan int)
	close(ch)
	_, ok := <-ch
	fmt.Println(ok)

	// 关闭后可以多次读取, 如果在多个goroutine中，可以起到一种广播的效果
	_, ok = <-ch
	fmt.Println(ok)
}
