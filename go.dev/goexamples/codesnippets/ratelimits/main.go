package main

import (
	"fmt"
	"time"
)

func main() {
	BurstsLimit := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		BurstsLimit <- time.Now()
	}
	
	go func() {
		for t := range time.NewTicker(time.Millisecond * 200).C {
			BurstsLimit <- t
		}
	}()

	// 任务、或者请求
	BurstsRequest := make(chan int, 5)
	go func() {
		for i := 0; i < 50; i++ {
			BurstsRequest <- i
		}
		close(BurstsRequest)
	}()

	for i := range BurstsRequest {
		// 通过这里来限制
		t := <-BurstsLimit
		fmt.Println("req: ", i, "t: ", t)
	}
}
