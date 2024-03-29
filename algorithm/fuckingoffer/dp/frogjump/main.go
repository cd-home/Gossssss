package main

// 剑指 Offer 10- II. 青蛙跳台阶问题

/*
	一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。
	求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
	答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
*/

func jumpNumWays(n int) int {
	const mod = 1e9 + 7
	if n == 0 || n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {

}
