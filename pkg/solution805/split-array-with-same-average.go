package solution805

// splitArraySameAverage 数组 A 的平均值与数组 nums 的平均值相等,数组 nums 中的每个元素减去 nums 的平均值，这样数组 nums 的平均值则变为 0
// --> 从 nums 中找出若干个元素组成集合 A，使得 A 的元素之和为 0，剩下的元素组成集合 B，它们的和也同样为 0
func splitArraySameAverage(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return false
	}

	sum := 0
	for _, x := range nums {
		sum += x
	}
	for i := range nums {
		nums[i] = nums[i]*n - sum
	}

	m := n / 2
	left := map[int]bool{}
	for i := 1; i < 1<<m; i++ {
		tot := 0
		for j, x := range nums[:m] {
			if i>>j&1 > 0 {
				tot += x
			}
		}
		if tot == 0 {
			return true
		}
		left[tot] = true
	}

	rsum := 0
	for _, x := range nums[m:] {
		rsum += x
	}
	for i := 1; i < 1<<(n-m); i++ {
		tot := 0
		for j, x := range nums[m:] {
			if i>>j&1 > 0 {
				tot += x
			}
		}
		if tot == 0 || rsum != tot && left[-tot] {
			return true
		}
	}
	return false
}
