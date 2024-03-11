package solution77

// combineHelper 是一个用于生成组合的辅助函数
func combineHelper(start, n, k int, current []int, result *[][]int) {
	// 如果达到所需长度，将当前组合添加到结果中
	if len(current) == k {
		// 注意：必须创建 current 的副本，因为切片是引用类型
		combination := make([]int, k)
		copy(combination, current)
		*result = append(*result, combination)
		return
	}

	// 回溯算法
	// 从当前数字开始，尝试每个数字
	for i := start; i <= n; i++ {
		current = append(current, i)              // 加入当前数字到组合中
		combineHelper(i+1, n, k, current, result) // 递归生成更长的组合
		current = current[:len(current)-1]        // 回溯，移除最后数字，尝试下一个
	}
}

// combine 是主函数，用于初始化并调用辅助函数
func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	combineHelper(1, n, k, []int{}, &result)
	return result
}
