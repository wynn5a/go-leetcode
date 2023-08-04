package solution560

// subarraySum 函数返回整数数组 nums 中，和为 k 的连续子数组的个数。
func subarraySum(nums []int, k int) int {
	// count 用于存放符合条件的子数组的个数，pre 用于存放到当前位置的前缀和
	count, pre := 0, 0
	// mp 是一个哈希表，键是前缀和，值是前缀和出现的次数
	mp := map[int]int{0: 1} // 初始化和为 0 的情况，因为单独一个数字就可以等于 k

	// 遍历整个 nums 数组
	for i := 0; i < len(nums); i++ {
		// 计算到当前位置 i 的前缀和
		pre += nums[i]
		// 如果哈希表中存在键值为 pre-k 的键，那么就找到了一个和为 k 的连续子数组
		if _, ok := mp[pre-k]; ok {
			// 对应的值即为和为 k 的连续子数组个数，加到 count 上
			count += mp[pre-k]
		}
		// 若哈希表中没有键值为 pre 的键，就添加进去，值为 1
		// 若已存在键值为 pre 的键，那就把对应的值加 1，表示再次出现了这一前缀和
		mp[pre]++
	}
	// 返回和为 k 的连续子数组的个数
	return count
}
