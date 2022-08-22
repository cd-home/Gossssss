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
	go Say("Hello")
	Say("World")
}
