package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	requests := make(chan int, 5)

	limiter := time.NewTicker(1000 * time.Millisecond)

	wg.Add(1)
	go func() {
		defer wg.Done()
		// 控制消费效率
		for req := range requests {
			t := <-limiter.C
			fmt.Println("request", req, t)
		}
	}()

	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	fmt.Println("Send Over")
	wg.Wait()
}
