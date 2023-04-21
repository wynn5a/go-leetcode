package solution80

func removeDuplicates(nums []int) int {
	// If the length of the array is less than or equal to 2, there are no duplicates to remove
	if len(nums) <= 2 {
		return len(nums)
	}
	// Initialize two pointers, slow and fast, to traverse the array
	slow, fast := 2, 2
	// Loop through the array with the fast pointer
	for fast < len(nums) {
		// If the current element is different from the element two positions before the current slow pointer,
		// it means we have found a new element to keep, so we copy it to the current slow position
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			// Increment the slow pointer to the next position
			slow++
		}
		// Increment the fast pointer to the next position
		fast++
	}
	// Return the new length of the array after removing duplicates
	return slow
}
