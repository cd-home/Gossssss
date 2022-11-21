package main

import (
	"fmt"
	"time"
)

func main() {
	// go 开启并发的执行单元
	// 程序会有 main goroutine
	// 主函数退出, 所有的goroutine退出

	// 创建 goroutine, 不会阻塞
	go spinner(time.Millisecond * 100)

	fmt.Printf("\rFibonacci(%d) = %d\n", 30, fib(30))

}

func spinner(delay time.Duration) {
	for {
		for _, r := range `—\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		// 故意延时一下
		time.Sleep(time.Millisecond * 100)
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
