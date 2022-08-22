package main

import (
	"fmt"
	"time"
)

func main() {
	Fibonacci := func(exit chan struct{}) <-chan int {
		ch := make(chan int)
		x, y := 0, 1
		go func() {
			defer close(ch)
			for {
				// case must be a channel
				// It chooses one at random if multiple are ready.
				select {
				case <-exit:
					fmt.Println("exit.")
					return
				case ch <- x:
					x, y = y, x+y
				}
			}
		}()
		return ch
	}

	exit := make(chan struct{})
	ch := Fibonacci(exit)
	for i := 0; i < 10; i++ {
		v, ok := <-ch
		if ok {
			fmt.Println(v)
		}
	}
	// receiver 10 times and close
	close(exit)
	time.Sleep(time.Second)

	// select blocked
	cjs := make(chan int, 1)
	go func() {
		time.Sleep(time.Second * 10)
		cjs <- 1
		close(cjs)
	}()
	select {
	case v := <-cjs:
		fmt.Println(v)
	}
}
