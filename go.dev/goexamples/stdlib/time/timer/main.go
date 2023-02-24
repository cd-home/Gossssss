package main

import (
	"fmt"
	"time"
)

func main() {
	// 定时器
	// timer 在将来某个时间点 回复你当前的时间
	timer1 := time.NewTimer(2 * time.Second)

	select {
	case t := <-timer1.C:
		fmt.Println("Timer 1 fired ", t)
	}

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	// 可以在某些情况下停止掉timer
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 Stopped")
	}
}
