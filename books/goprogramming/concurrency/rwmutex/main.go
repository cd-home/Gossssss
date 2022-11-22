package main

import "sync"

func main() {
	// 	读写锁
	// 读可以并发, 写是互斥的
}

var (
	rw      sync.RWMutex
	balance int
)

func Balance() int {
	// 读可并发
	rw.Lock()
	defer rw.Unlock()
	return balance
}
