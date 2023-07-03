package solution16

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	// Sort the input array
	sort.Ints(nums)

	n := len(nums)

	if n == 0 {
		return 0
	}

	closestSum := nums[0] + nums[1] + nums[2] // Initialize the closest sum

	for i := 0; i < n-2; i++ {

		// Skip duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1  // Pointer for the element to the right of nums[i]
		right := n - 1 // Pointer for the last element in the array
		sum := 0       // Sum of three elements

		for left < right {
			sum = nums[i] + nums[left] + nums[right]

			// If the sum is equal to the target, return it
			if sum == target {
				return sum
			}

			// If the current sum is closer to the target, update the closest sum
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-closestSum)) {
				closestSum = sum
			}

			// Adjust the pointers based on the sum compared to the target
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				// If the sum is equal to the target, return the sum
				return sum
			}
		}
	}

	return closestSum
}
