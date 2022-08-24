package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mux sync.Mutex
	v   map[string]int
}

func (sc *SafeCounter) Inc(wg *sync.WaitGroup, key string) {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	defer wg.Done()
	sc.v[key]++
}

func (sc *SafeCounter) Value(key string) int {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	return sc.v[key]
}

func main() {
	var wg sync.WaitGroup
	sc := &SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go sc.Inc(&wg, "oneKey")
	}
	wg.Wait()
	fmt.Println(sc)
	fmt.Println(sc.Value("oneKey"))
}
