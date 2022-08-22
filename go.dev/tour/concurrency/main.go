package main

import (
	"fmt"
	"time"
)

func Say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Second)
	}
}

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
	//go Say("Hello")
	//Say("World")

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

	BufferedChannel()
}
