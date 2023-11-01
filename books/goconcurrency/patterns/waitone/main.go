package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(FirstResponse())
}

func FirstResponse() string {
	runner := 10
	// buffer 避免内存泄漏
	ch := make(chan string, runner)
	for i := 0; i < runner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func runTask(i int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("%d", i)
}
