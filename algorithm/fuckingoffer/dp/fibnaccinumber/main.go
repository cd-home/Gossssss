package main

// 剑指 Offer 10- I. 斐波那契数列

/*
	写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。
	斐波那契数列的定义如下：
		F(0) = 0,F(1)= 1
		F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
*/

func fib(n int) int {
	const mod = 1e9 + 7
	if n < 2 {
		return n
	}
	// n是dp的下标, 元素个数为 n+1
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % mod
	}
	return dp[n]
}

func main() {

}
