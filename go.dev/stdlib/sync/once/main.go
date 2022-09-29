package main

import (
	"fmt"
	"sync"
)

func main() {
	// Once is an object that will perform exactly one action.
	// A Once must not be copied after first use.
	var once sync.Once
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(func() {
				fmt.Println("fmt Once")
			})
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}
