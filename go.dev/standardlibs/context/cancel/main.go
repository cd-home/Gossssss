package main

import (
	"context"
	"fmt"
)

func main() {
	gen := func(ctx context.Context) chan int {
		n := 0
		out := make(chan int)
		go func() {
			// 关闭chan
			defer close(out)
			for {
				select {
				case <-ctx.Done():
					// return 避免泄漏
					return
				case out <- n:
					n++
				}
			}
		}()
		return out
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for d := range gen(ctx) {
		fmt.Println(d)
		if d == 5 {
			break
		}
	}
}
