package solution15

import "sort"

func threeSum(nums []int) [][]int {
	// Sort the slice
	sort.Ints(nums)

	// Initialize an empty slice for storing triplets
	var result [][]int

	// Iterate over each element except the last two
	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Set two pointers at both ends of the remaining slice
		left := i + 1
		right := len(nums) - 1

		// While there are still elements between left and right
		for left < right {
			// Calculate the sum of three elements
			s := nums[i] + nums[left] + nums[right]

			// If the sum is zero, we found a triplet
			if s == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicates on both sides
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// Move both pointers towards each other
				left++
				right--
				// If the sum is too small, move left pointer forward
			} else if s < 0 {
				left++

				// If the sum is too large, move right pointer backward
			} else {
				right--
			}
		}
	}

	return result
}
