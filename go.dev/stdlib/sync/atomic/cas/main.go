package main

import (
	"fmt"
	"sync/atomic"
)

var num int64 = 10

func main() {
	old := atomic.LoadInt64(&num)
	var newVal int64 = 20
	if atomic.CompareAndSwapInt64(&num, old, newVal) {
		fmt.Println("Compare And Set")
	}
}
