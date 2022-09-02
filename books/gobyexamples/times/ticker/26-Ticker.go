package main

import (
	"fmt"
	"time"
)

func main() {
	// 打点器， 重复执行
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
