package main

import (
	"fmt"
	"time"
)

func main() {
	BurstsLimit := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		BurstsLimit <- time.Now()
	}
	go func() {
		for t := range time.NewTicker(time.Millisecond * 200).C {
			BurstsLimit <- t
		}
	}()

	BurstsRequest := make(chan int, 5)
	for i := 0; i < 5; i++ {
		BurstsRequest <- i
	}
	close(BurstsRequest)

	for i  := range BurstsRequest {
		t := <-BurstsLimit
		fmt.Println("req: ", i, "t: ", t)
	}
}
