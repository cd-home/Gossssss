package fan

import (
	"sync"
	"time"
)

func Producer(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			out <- i
		}
	}()
	return out
}

func Square(inCh <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range inCh {
			out <- n * n
			time.Sleep(time.Second)
		}
	}()
	return out
}

func Merge(res ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(res))

	for _, r := range res {
		go func(r <-chan int) {
			defer wg.Done()
			for n := range r {
				out <- n
			}
		}(r)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
