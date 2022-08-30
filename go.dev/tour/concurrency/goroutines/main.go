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

func main() {
	// A goroutine is a lightweight thread managed by the Go runtime.
	// access to shared memory must be synchronized
	go Say("Hello")
	Say("World")
}
