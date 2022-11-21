package main

func main() {
	// goroutine 是并发基本单元
	// channel 是 g 之间的通信机制

	// 并发程序实现 goroutine 与 channel, 值通过channel在goroutine之间传递
	// communicating sequential processes 顺序进程通信

	// 传统的并发模型：多线程共享变量

	// 对于开启的g, 需要知道是否需要退出, 能否正常退出
	// 泄漏的goroutines并不会被自动回收，因此确保每个不再需要的goroutine能正常退出是重要的
}
