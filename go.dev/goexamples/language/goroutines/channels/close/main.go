package main

import "fmt"

func main() {
	done := make(chan bool)
	jobs := make(chan int, 5)

	go func() {
		for {
			// more是布尔值，表示通道是否有数据
			job, more := <-jobs
			// 关闭后的通道是可以读取的, 当通道关闭并且没有数据读取时: more==false
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
	// 发送完成 根据需要关闭通道, 假设后续不再往里面写入数据
	// 也是一种通知, 消息沟通
	close(jobs)
	<-done
}
