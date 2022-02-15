package solution1984

import (
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	if len(nums) < 2 {
		return 0
	}
	if k < 2 {
		return 0
	}
	sort.Ints(nums)
	res := math.MaxInt32
	for i, num := range nums[:len(nums)-k+1] {
		res = min(res, nums[i+k-1]-num)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
