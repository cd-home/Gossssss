package main

// 剑指 Offer 42. 连续子数组的最大和

/*
	输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
	要求时间复杂度为O(n)。
	示例1:
		输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
		输出: 6
		解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
*/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	max := dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
			// dp[i-1] < 0 时, 再往后加就是负作用了, 当前连续子数组断掉 dp[i] = nums[i]
		} else {
			dp[i] = nums[i]
		}
		// 记录最大值
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func main() {

}
