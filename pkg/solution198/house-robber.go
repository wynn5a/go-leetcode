package solution198

// rob dynamic programming solution
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// dp[i] represents the maximum amount of money that can be robbed up to house i
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	if len(nums) > 1 {
		dp[1] = max(nums[0], nums[1])
	}

	for i := 2; i < len(nums); i++ {
		// dp[i] = max(dp[i-1], dp[i-2] + nums[i])
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}
