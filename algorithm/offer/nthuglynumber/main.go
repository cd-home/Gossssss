package main

import "fmt"

func main() {
	fmt.Println(nthUglyNumber(10))
}

func nthUglyNumber(n int) int {
	// dp 存储丑数
	dp := make([]int, n+1)
	dp[0] = 1
	// 丑数指针
	a, b, c := 0, 0, 0
	for i := 1; i < n; i++ {
		// 下一个丑数, 必是上一个丑数*2 或*3 或者*5中最小的一个
		x2, x3, x5 := dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = min(min(x2, x3), x5)
		if dp[i] == x2 {
			a++
		}
		if dp[i] == x3 {
			b++
		}
		if dp[i] == x5 {
			c++
		}
	}
	return dp[n-1]
}

func min(m, n int) int {
	if m > n {
		return n
	}
	return m
}
