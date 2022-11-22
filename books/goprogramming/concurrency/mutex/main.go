package main

import "sync"

func main() {
	// go 没有重入锁,
	// 加锁、解锁成对出现

	// 将变量限定在goroutine内部；如果是多个goroutine都需要访问的变量，使用互斥条件来访问
}

var (
	mu      sync.Mutex
	balance int
)

func Deposit(amount int) {
	// 写互斥
	mu.Lock()
	// 代码临界区
	// 当前 goroutine 才能读写
	balance = balance + amount
	mu.Unlock()
}

func Balance() int {
	// 读互斥
	mu.Lock()
	defer mu.Unlock()
	return balance
}
