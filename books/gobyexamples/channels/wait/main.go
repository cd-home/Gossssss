package main

import (
	"fmt"
	"time"
)

func Work(done chan bool)  {
	fmt.Println("Working...")
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
	done <- true
}

func main() {
	done := make(chan bool)

	go Work(done)

	// 阻塞等待
	<-done
}
