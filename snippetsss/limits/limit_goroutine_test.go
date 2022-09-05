package limits

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func SendTask(buf chan int, task int, wg *sync.WaitGroup) {
	wg.Add(1)
	buf <- task
}

func Work(buf chan int, wg *sync.WaitGroup) {
	for task := range buf {
		fmt.Println(task, runtime.NumGoroutine())
		wg.Done()
	}
}

// TestWorkPoolGoroutines will be better than TestLimitGoroutines
func TestWorkPoolGoroutines(t *testing.T) {
	t.Helper()
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

func doWork(limit chan struct{}, i int) {
	fmt.Println(i, runtime.NumGoroutine())
	time.Sleep(time.Second)
	<-limit
}

func TestLimitGoroutines(t *testing.T) {
	t.Helper()
	limit := make(chan struct{}, 3)
	for i := 0; i < 10; i++ {
		limit <- struct{}{}
		go doWork(limit, i)
	}
}
