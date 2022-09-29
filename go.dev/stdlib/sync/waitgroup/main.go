package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// A WaitGroup waits for a collection of goroutines to finish.
	// 不能被赋值, 如果要传递必须是指针
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			// 通常使用defer
			defer wg.Done()
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
		}(i)
	}
	wg.Wait()
}
