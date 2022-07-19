package sync_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWaitForAllGroutines(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
		}(i)
	}
	wg.Wait()
}

func doSomething(wg sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("done")
}

func TestCopyWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	doSomething(wg)
	wg.Wait()
}
