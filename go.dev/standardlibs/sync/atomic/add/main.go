package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var num int64 = 0
var wg sync.WaitGroup

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&num, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(num)
}
