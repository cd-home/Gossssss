package snippetsss

import (
	"testing"
	"runtime"
	"fmt"
	"time"
)

func doWork(limit chan struct{}, i int)  {
	fmt.Println(i, runtime.NumGoroutine())
	time.Sleep(time.Second)
	<-limit
}

func TestLimitGoroutines(t *testing.T) {
	limit := make(chan struct{}, 3)
	for i := 0; i < 10; i++ {
		limit <- struct{}{}
		go doWork(limit, i)
	}
}