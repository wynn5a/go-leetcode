package solution1220

//动态规划
func countVowelPermutation(n int) (ans int) {
	const mod int = 1e9 + 7
	dp := [5]int{1, 1, 1, 1, 1}
	for i := 1; i < n; i++ {
		dp = [5]int{
			(dp[1] + dp[2] + dp[4]) % mod, // a 前面可以为 e,u,i
			(dp[0] + dp[2]) % mod,         // e 前面可以为 a,i
			(dp[1] + dp[3]) % mod,         // i 前面可以为 e,o
			dp[2],                         // o 前面可以为 i
			(dp[2] + dp[3]) % mod,         // u 前面可以为 i,o
		}
	}
	for _, v := range dp {
		ans = (ans + v) % mod
	}
	return
}
