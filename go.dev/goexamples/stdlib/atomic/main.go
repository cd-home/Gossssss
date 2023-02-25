package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// atomic 并发下修改某个值的状态
	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				// 并发安全的, 提供低级别的原语来实现同步算法
				atomic.AddUint64(&ops, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Println(atomic.LoadUint64(&ops))
}
