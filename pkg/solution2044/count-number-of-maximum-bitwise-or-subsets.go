package solution2044

func countMaxOrSubsets(nums []int) (ans int) {
	maxOr := 0
	for i := 1; i < 1<<len(nums); i++ {
		or := 0
		for j, num := range nums {
			if i>>j&1 == 1 {
				or |= num
			}
		}
		if or > maxOr {
			maxOr = or
			ans = 1
		} else if or == maxOr {
			ans++
		}
	}
	return
}
