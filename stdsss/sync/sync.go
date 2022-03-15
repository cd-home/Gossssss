package main

import (
	"fmt"
	"sync"
	"time"
)

// sync WaitForAllGroutines
func WaitForAllGroutines() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Second)
			fmt.Println(i)
		}(i)
	}

	// wait for all goroutines to finish
	wg.Wait()
}


func main() {
	WaitForAllGroutines()
}