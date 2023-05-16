package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	GoroutineLeak()
}

func GoroutineLeak() {
	doWork := func(strings <-chan string) <-chan struct{} {
		completed := make(chan struct{})
		go func() {
			defer fmt.Println("doWork done.")
			defer close(completed)
			// nil chan will make goroutine mem leak
			for s := range strings {
				fmt.Println(s)
			}
		}()
		return completed
	}
	doWork(nil)
	time.Sleep(time.Second * 2)
	fmt.Println("Done")
}

func NotifyGoroutineExitAvoidLeak() {
	doWork := func(done <-chan struct{}, strings <-chan string) <-chan struct{} {
		terminated := make(chan struct{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)
			for {
				select {
				case <-done:
					return
				case s := <-strings:
					fmt.Println(s)
					time.Sleep(time.Second)
				}
			}
		}()
		return terminated
	}

	// send to doWork signal
	done := make(chan struct{})

	// first param done notify signal; doWork return terminated signal
	// who create goroutine who in charge
	terminated := doWork(done, nil)

	// if parent goroutine cancel son goroutine
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("cancel doWork goroutine.")
		close(done)
	}()

	// wait
	<-terminated
	fmt.Println("done.")
}

func GoroutineBlocked() {
	newRandStream := func(done <-chan struct{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			// 由于消费者停止消费, 导致 goroutine 阻塞在这
			for {
				select {
				case <-done:
					return
				case randStream <- rand.Int():
				}
			}
		}()
		return randStream
	}
	done := make(chan struct{})

	randStream := newRandStream(done)

	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)
	time.Sleep(time.Second * 2)
	fmt.Println("done.")
}
