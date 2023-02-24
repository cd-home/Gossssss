package main

import (
	"fmt"
	"sync"
	"time"
)

// wg 如果要传递必须通过指针传递，A WaitGroup must not be copied after first use.
// 通常 wg 只是借用给子goroutine, 所以尽量不要传递
func worker(id int) {
	// 完成一个就 Add(-1)
	defer wg.Done()
	fmt.Printf("worker %d staring\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

var wg = sync.WaitGroup{}

func main() {

	for i := 0; i < 5; i++ {
		// 计数器
		wg.Add(1)
		go worker(i)
	}
	// 阻塞等待所有的g执行完成
	wg.Wait()
}
