package solution238

// productExceptSelf two loops, one gets the product of all numbers before i,
// the other gets the product of all numbers after i
// Time complexity: O(n)
// Space complexity: O(1)
func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	ans[0] = 1
	// 得到 i 之前的所有数字的乘积
	// ans[i] = nums[0] * nums[1] * ... * nums[i-1]
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}

	r := 1
	// 得到 i 之后的所有数字的乘积
	// ans[i] = ans[i] * nums[i+1] * ... * nums[n-1]
	for i := n - 1; i >= 0; i-- {
		ans[i] = ans[i] * r
		r *= nums[i]
	}
	return ans
}
