package solution239

import "container/list"

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if k == 1 {
		return nums
	}
	if k == len(nums) {
		return []int{m(nums)}
	}

	result := make([]int, len(nums)-k+1)
	deque := list.New()

	for i := 0; i < len(nums); i++ {
		// Remove elements out of window from the front of deque
		if deque.Len() > 0 && deque.Front().Value.(int) < i-k+1 {
			deque.Remove(deque.Front())
		}

		// Remove elements smaller than current from the back of deque
		for deque.Len() > 0 && nums[deque.Back().Value.(int)] < nums[i] {
			deque.Remove(deque.Back())
		}

		deque.PushBack(i)

		// The first element of deque is the maximum of the window
		if i >= k-1 {
			result[i-k+1] = nums[deque.Front().Value.(int)]
		}
	}

	return result
}

func m(a []int) int {
	r := 0
	for _, v := range a {
		if v > r {
			r = v
		}
	}
	return r
}
