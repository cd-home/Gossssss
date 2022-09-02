package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, r chan<- int)  {
	// 并发的worker都去jobs通道中获取任务
	for job := range jobs {
		fmt.Println("worker(id):", id)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", job)
		r <- job * 2
	}
}

func main() {
	const JobsNums = 5
	jobs := make(chan int, JobsNums)
	results := make(chan int, JobsNums)

	for w := 0; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for i := 0; i < JobsNums; i++ {
		jobs <- i
	}

	for i := 0; i < JobsNums; i++ {
		fmt.Println(<-results)
	}
}
