package main

func main() {
	/*
		1. 数据竞争会在两个以上的goroutine并发访问相同的变量且至少其中一个为写操作时发生
		2. 避免数据竞争
				避免写变量
				避免并发访问
				互斥
		3. 条件静态检测
			go run -race
			go test -race
		4. 线程
			2MB Thread 栈
			2KM goroutine 动态伸缩 最大1GB

			线程 调度 硬件计时器
				挂起 寄存器 => 内存; 运行 内存=> 寄存器;
			协程 Runtime
				N:M, 调度goroutine 不需要进入内核的上下文
			GOMAXPROCS
				决定多个核心执行Go代码 即是 N
			Goroutine
				没有ID
	*/
}
