package solution32

func longestValidParentheses(s string) int {
	maxAns := 0
	dp := make([]int, len(s)) //dp[i] 以 s[i] 结尾的最长有效括号组合
	//当 s[i]=) 且 s[i−1]=(, dp[i]=dp[i-2]+2
	//当 s[i]=) 且 s[i-1]=)，如果 s[i-1] 是有效组合--记作sub, sub的长度为 dp[i-1]--的一部分，
	//那么 s[i] 如果有效，那么它对应的 （ 应该是在 s[i-dp[i-1]-1] 这个位置,
	//此时 dp[i] = dp[i-dp[i-1]-2] + [i-1] + 2
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			maxAns = max(maxAns, dp[i])
		}
	}
	return maxAns
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
