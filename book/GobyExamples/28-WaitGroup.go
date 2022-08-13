package main

import (
	"fmt"
	"sync"
	"time"
)

// wg 必须通过指针传递，A WaitGroup must not be copied after first use.
func wk(id int, wg *sync.WaitGroup)  {
	// 完成一个就 Add(-1)
	defer wg.Done()
	fmt.Printf("worker %d staring\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		// 计数器
		wg.Add(1)
		go wk(i, wg)
	}
	// 阻塞等待所有的g执行完成
	wg.Wait()
}
