package main

// 剑指 Offer 47. 礼物的最大价值

/*
	在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。
	你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。
	给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
	示例 1:
		输入:
		[
		 [1,3,1],
		 [1,5,1],
		 [4,2,1]
		]
		输出: 12
		解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物
*/

func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	var dp [][]int
	for i := 0; i < len(grid); i++ {
		row := make([]int, len(grid[0]))
		dp = append(dp, row)
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i-1][j]
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j] + grid[i][j]
				} else {
					dp[i][j] = dp[i][j-1] + grid[i][j]
				}
			}
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func main() {

}
