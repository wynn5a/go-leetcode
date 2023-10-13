package solution

import (
	"fmt"
	"math/rand"
)

// FindLongestIncreasingSubsequence will find the length of longest increasing subsequence
// from the giving unsorted array of integers.
func FindLongestIncreasingSubsequence(array []int32) int {
	return dpOnly(array)
}

// dpOnly solution will only use DP to solve this.
// dp[i] = max(dp[i], dp[j] + 1) for j in [0, i)
func dpOnly(array []int32) int {
	l := len(array)
	if l < 2 {
		return l
	}

	//create a array with 1 as default value to store the result for every number
	dp := make([]int, l)
	for i := range dp {
		dp[i] = 1
	}

	//make sure each of value in dp is the
	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if array[i] > array[j] && dp[i] <= dp[j] {
				dp[i] = dp[j] + 1
			}
		}
	}

	return MaxInArray(dp)
}

func Run300() {
	fmt.Println("--- LeetCode NO.300 begin")
	input := make([]int32, 10)
	for i := range input {
		input[i] = rand.Int31n(100)
	}
	fmt.Println("input array      :", fmt.Sprint(input))
	ans := FindLongestIncreasingSubsequence(input)
	fmt.Printf("answer           : %v\n", fmt.Sprint(ans))
	fmt.Println("LeetCode NO.300 end")
}
