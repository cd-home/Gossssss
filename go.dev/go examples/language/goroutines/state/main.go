package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key int
	val int
	resp chan bool
}

func main() {
	// Go 共享内存的思想是，通过通信使每个数据仅被单个协程所拥有，即通过通信实现共享内存
	var readOps uint64
	var writeOps uint64

	// 通过这两个通道来对state发布读写请求
	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		// state被单独的g拥有
		var state = make(map[int]int)
		for {
			select {
			// 读写请求
			case read := <- reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	fmt.Println("readOps", atomic.LoadUint64(&readOps))
	fmt.Println("writeOps,", atomic.LoadUint64(&writeOps,))
}
