package solution_offer_47

// TODO: 空间优化，因为只和上面及前面的值有关，所以 DP 数组可以直接省略成两个变量
func maxValue(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			value := grid[i][j]
			if i < 1 && j < 1 {
				dp[i][j] = value
			}
			if i < 1 && j >= 1 {
				dp[i][j] = dp[i][j-1] + value
			}
			if i >= 1 && j < 1 {
				dp[i][j] = dp[i-1][j] + value
			}
			if i >= 1 && j >= 1 {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + value
			}
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func max(i int, j int) int {
	if i < j {
		return j
	}
	return i
}
