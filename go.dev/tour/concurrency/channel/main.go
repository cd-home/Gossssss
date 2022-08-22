package main

import (
	"fmt"
)

// BufferedChannel
// Sends to a buffered channel block only when the buffer is full.
// Receives block when the buffer is empty.
func BufferedChannel() {
	ch := make(chan int, 10)
	// Not blocked
	ch <- 1
	ch <- 2
	close(ch)

	// Read Until Close
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	sum := func(slices []int, res chan int) {
		sums := 0
		for _, v := range slices {
			sums += v
		}
		res <- sums
	}
	s := []int{7, 2, 8, -9, 4, 0}
	ch := make(chan int)
	go sum(s[:3], ch)
	go sum(s[3:], ch)
	x, y := <-ch, <-ch
	fmt.Println(x, y, x+y)

	// buffer channel
	BufferedChannel()

	// read channel
	cds := make(chan string, 3)
	cds <- "Hello World!"
	cds <- "Go"
	cds <- "Python"
	close(cds)
	for {
		v, ok := <-cds
		if !ok {
			break
		}
		fmt.Println(ok, v)
	}

	// close channel and for-range
	Fibonacci := func(n int) <-chan int {
		x, y := 0, 1
		// closure: do not expose ch
		// only expose read channel
		ch := make(chan int)
		go func() {
			// sender close ch
			// when is necessary to close a channel
			defer close(ch)
			for i := 0; i < n; i++ {
				ch <- x
				x, y = y, x+y
			}
		}()
		return ch
	}
	// for-range until close channel
	for v := range Fibonacci(10) {
		fmt.Println(v)
	}
}
