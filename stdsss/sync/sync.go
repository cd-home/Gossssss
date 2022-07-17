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

type SafeMap struct {
	sync.Mutex
	m map[string]string
}

func (m SafeMap) Set(key string, value string) {
	defer m.Unlock()
	m.Lock()
	m.m[key] = value
}

func main() {
	sm := &SafeMap{m: make(map[string]string)}
	for i := 0; i < 10; i++ {
		go func(i int) {
			sm.Set(string(i), string(i))
		}(i)
	}
	WaitForAllGroutines()
}
