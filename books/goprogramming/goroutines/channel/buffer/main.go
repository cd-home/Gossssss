package main

func main() {
	// make 创建通道, cap 无缓冲与有缓冲
	// 通道的零值是 nil
	// 通道是可比较的，类型相同的通道是可比较的
	//ch := make(chan int)
	//bufCh := make(chan int, 5)

	// 无缓冲的channel必须发送、接受同步准备
	// 有缓冲的channel可以异步, 队列满，发送阻塞，队列空，接受阻塞

	// 建立了通道, 通常就是要进行 读写, 也就是需要和 goroutine 相关

	// 关于无缓存或带缓存channels之间的选择,或者是带缓存channels的容量大小的选择,都可能影响程序的正确性。
	// 1. 无缓存channel更强地保证了每个发送操作与相应的同步接收操作;
	// 2. 带缓存channel, 这些操作是解耦的, 但是并不一定能提高性能, 需要根据实际情况调试大小
}
