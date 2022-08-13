package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// 使用lock对数据的同步操作，安全的访问数据
	var state = make(map[int]int)
	var mutex = sync.Mutex{}

	var readOps uint64
	var writeOps uint64

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				// state跨越多个g，同步lock同步
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 100; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				value := rand.Intn(100)
				mutex.Lock()
				// state跨越多个g，同步lock同步
				state[key] = value
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second * 5)
	fmt.Println("readOps", atomic.LoadUint64(&readOps))
	fmt.Println("writeOps,", atomic.LoadUint64(&writeOps,))

	mutex.Lock()
	fmt.Println("state", state)
	mutex.Unlock()
}
