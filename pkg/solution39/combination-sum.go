package solution39

// 回溯算法 + 剪枝 先画出树形图，然后编码实现
// combinationSum returns all unique combinations of candidates that sum up to target
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var dfs func(int, int, []int)

	// dfs is a recursive function that finds all combinations of candidates that sum up to target
	dfs = func(idx int, target int, path []int) {
		// if target is 0, we have found a valid combination and add it to the result
		if target == 0 {
			res = append(res, append([]int(nil), path...))
			return
		}
		// if we have exhausted all candidates, we return
		if idx == len(candidates) {
			return
		}
		// if target is greater than or equal to the current candidate, we include the current candidate in the path and recursively call dfs with the updated target and path
		if target >= candidates[idx] {
			dfs(idx, target-candidates[idx], append(path, candidates[idx]))
		}
		// we exclude the current candidate and recursively call dfs with the original target and path
		dfs(idx+1, target, path)
	}

	// we start dfs with the first candidate and an empty path
	dfs(0, target, []int{})
	return res
}
