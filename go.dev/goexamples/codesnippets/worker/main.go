package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, r chan<- int) {
	// 注意这里就使用了只读的 通道 jobs, r 只写
	// 并发的worker都去jobs通道中获取任务
	for job := range jobs {
		fmt.Println("worker(id):", id)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", job)
		r <- job * 2
	}
}

func main() {
	const JobsNums = 1000
	jobs := make(chan int, JobsNums)
	results := make(chan int, JobsNums)

	for w := 0; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for i := 0; i < JobsNums; i++ {
		jobs <- i
		// 写完关闭;
		// 一个适用的原则是不要从接收端关闭channel，也不要在多个并发发送端中关闭channel
		// 并且禁止关闭只读通道
		// 发送端关闭
	}
	close(jobs)

	for i := 0; i < JobsNums; i++ {
		fmt.Println(<-results)
	}
	// 读取完成
	close(results)
}
