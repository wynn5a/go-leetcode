package solution128

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 1. 创建一个map，key为nums中的值，value为true
	m := make(map[int]bool, len(nums))
	for _, num := range nums {
		m[num] = true
	}
	// 2. 遍历nums，如果当前值 n-1 不在map中，则说明当前值是一个连续序列的起始值
	// 3. 从起始值开始，不断+1，直到+1的值不在map中，此时得到一个连续序列的长度
	// 4. 比较当前连续序列的长度与最大长度，取最大值
	maxLen := 0
	for _, num := range nums {
		if _, ok := m[num-1]; !ok {
			curLen := 1
			for {
				if _, ok := m[num+1]; ok {
					curLen++
					num++
				} else {
					break
				}
			}
			if curLen > maxLen {
				maxLen = curLen
			}
		}
	}
	return maxLen
}
