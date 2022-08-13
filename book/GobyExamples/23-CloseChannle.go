package main

import "fmt"

func main() {
	done := make(chan bool)
	jobs := make(chan int, 5)

	go func() {
		for {
			// more是布尔值，表示通道是否有数据
			job, more := <-jobs
			if more {
				fmt.Println("recv: ", job)
			} else {
				fmt.Println("recv all")
				done <- true
				return
			}
		}
	}()

	for i := 0; i <= 3; i++ {
		jobs <- i
		fmt.Println("send job: ", i)
	}
	// 发送完成 关闭通道
	close(jobs)
	<-done
}
