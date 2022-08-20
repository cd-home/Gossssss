package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		c1 <- "One"
	}()

	c := time.After(5 * time.Second)
	for {
		select {
		case v1 := <-c1:
			fmt.Println(v1)
			return
		case <-c:
			fmt.Println("timeout")
			return
		}
	}
}
