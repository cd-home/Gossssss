package main

import (
	"fmt"
	"strconv"
)

func main() {
	// blocked channel
	ch := make(chan string)
	go func() {
		ch <- "yao"
	}()
	fmt.Println(<-ch)

	// buffered channel
	bufferedCh := make(chan string, 10)
	go func() {
		defer close(bufferedCh)
		for i := 0; i < 10; i++ {
			bufferedCh <- strconv.Itoa(i)
		}
	}()
	for s := range bufferedCh {
		fmt.Println(s)
	}
}
