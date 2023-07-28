package solution88

func merge(nums1 []int, m int, nums2 []int, n int) {

	//use builtin function sort.Ints
	//copy(nums1[m:], nums2[:n])
	//sort.Ints(nums1[:m+n])

	// without builtin function
	// Move elements of num1[] in the end
	j := m - 1
	for i := m + n - 1; i >= n; i-- {
		nums1[i] = nums1[j]
		j--
	}

	// Merge arrays
	i := n
	j = 0
	k := 0

	for i < m+n && j < n {
		if nums1[i] <= nums2[j] {
			nums1[k] = nums1[i]
			k++
			i++
		} else {
			nums1[k] = nums2[j]
			j++
			k++
		}
	}

	// Copy remaining elements of nums2[] if any
	for j < n {
		nums1[k] = nums2[j]
		j++
		k++
	}
}
