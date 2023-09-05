package solution2605

import (
	"sort"
)

func minNumber(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	i := 0
	j := 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			return nums1[i]
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	m := nums1[0]
	n := nums2[0]
	if n < m {
		return n*10 + m
	} else {
		return m*10 + n
	}
}
