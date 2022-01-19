package solution219

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if val, ok := m[nums[i]]; ok {
			if i-val <= k {
				return true
			} else {
				m[nums[i]] = i
			}
		} else {
			m[nums[i]] = i
		}
	}
	return false
}
