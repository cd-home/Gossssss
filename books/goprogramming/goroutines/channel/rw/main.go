package main

import "fmt"

func main() {
	// read-only
	gen := func() <-chan int {
		// 读写channel
		rw := make(chan int)
		go func() {
			//defer close(rw)
			for i := 0; i < 10; i++ {
				rw <- i
			}
			close(rw)
		}()
		// 返回只读的chan
		return rw
	}
	r := gen()
	for v := range r {
		fmt.Println(v)
	}

	// 当一个channel作为一个函数参数时，它一般总是被专门用于只发送或者只接收
}
