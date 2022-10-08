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
		go func(i int) {
			atomic.StoreInt64(&num, int64(i))
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(atomic.LoadInt64(&num))
}
