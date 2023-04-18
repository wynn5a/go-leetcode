package solution45

func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	jumps := 1         // start with one jump from the first index
	curEnd := nums[0]  // current farthest index that can be reached in one jump
	nextEnd := nums[0] // farthest index that can be reached in the next jump

	for i := 1; i < n-1; i++ {
		nextEnd = max(nextEnd, i+nums[i]) // update the farthest index that can be reached in the next jump
		if i == curEnd {                  // if the current farthest index is reached, a new jump is needed
			jumps++          // increment the number of jumps
			curEnd = nextEnd // update the current farthest index to the farthest index that can be reached in the next jump
		}
	}
	return jumps
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}
