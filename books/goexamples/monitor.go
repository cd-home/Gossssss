package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	go monitor()
	n := 10
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Set(rand.Intn(10 * n))
		}()
	}
	wg.Wait()
	fmt.Printf("%d\n", read())
}

var readValue = make(chan int)
var writeValue = make(chan int)

func Set(newValue int) {
	writeValue <- newValue
}

func read() int {
	return <-readValue
}

func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeValue:
			value = newValue
			fmt.Println(value)
		case readValue <- value:
		}
	}
}