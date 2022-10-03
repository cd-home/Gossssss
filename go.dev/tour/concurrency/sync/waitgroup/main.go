package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	testMap := make(map[int]int)

	// lock := &sync.Mutex{}
	// zero value can be used
	var lock sync.Mutex

	wg.Add(3)

	for i := 0; i < 3; i++ {
		// closure
		go func(j int) {
			defer wg.Done()
			lock.Lock()
			testMap[j] = j
			time.Sleep(time.Second * time.Duration(j))
			fmt.Println(testMap)
			lock.Unlock()
		}(i)
	}
	wg.Wait()
}
