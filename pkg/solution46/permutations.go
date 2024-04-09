package solution46

func permute(nums []int) [][]int {
	var res [][]int
	var backtrack func([]int, []int)
	backtrack = func(nums []int, path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for _, num := range nums {
			if contains(path, num) {
				continue
			}
			path = append(path, num)
			backtrack(nums, path)
			path = path[:len(path)-1]
		}
	}
	backtrack(nums, []int{})
	return res
}

func contains(path []int, num int) bool {
	for _, p := range path {
		if p == num {
			return true
		}
	}
	return false
}
