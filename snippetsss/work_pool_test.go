package snippetsss

import (
	"fmt"
	"sync"
	"testing"
	"runtime"
)

func SendTask(buf chan int,  task int, wg *sync.WaitGroup) {
	wg.Add(1)
	buf <- task
}

func Work(buf chan int, wg *sync.WaitGroup) {
	for task := range buf {
		fmt.Println(task, runtime.NumGoroutine())
		wg.Done()
	}
}

func TestWorkPoolGoroutines(t *testing.T) {

	// watch out: wg must be *sync.WaitGroup (pointer)
	wg := &sync.WaitGroup{}

	buf := make(chan int)

	// Wokers: 消费
	worker := 3
	for i := 0; i < worker; i++ {
		go Work(buf, wg)
	}

	// 生产
	tasks := 100
	for i := 0; i < tasks; i++ {
		SendTask(buf, i, wg)
	}

	wg.Wait()
}