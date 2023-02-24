package main

import (
	"fmt"
	"time"
)

func main() {
	// 打点器， 重复执行; timer是未来某个时刻; ticker 是重复任务
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t, t.Unix())
			}
		}
	}()

	time.Sleep(3500 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker Stop")
}
