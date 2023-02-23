package main

import (
	"fmt"
	"sync/atomic"
)

var num int64 = 10

func main() {
	old := atomic.SwapInt64(&num, 20)
	fmt.Println(num)
	fmt.Println(old)
}
