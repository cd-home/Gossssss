package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var num int64 = 0

func main() {
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&num, 1)
			time.Sleep(time.Second)
		}()
	}
	for i := 0; i < 10; i++ {
		fmt.Println(atomic.LoadInt64(&num))
		time.Sleep(time.Second)
	}
}
