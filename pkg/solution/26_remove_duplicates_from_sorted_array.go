package solution

import (
	"fmt"
	"math/rand"
	"sort"
)

// removeDuplicatesFromSortedArray is working as :
// Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.
// Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
func removeDuplicatesFromSortedArray(array []int) int {
	i, j, l := 0, 1, len(array)
	for ; j < l; j++ {
		if array[i] == array[j] {
			continue
		}
		i++
		array[i], array[j] = array[j], array[i]
	}
	return i + 1
}

func Run26() {
	fmt.Println("--- LeetCode NO.26 begin")
	sortedArray := make([]int, 10)
	for i := range sortedArray {
		sortedArray[i] = rand.Intn(10)
	}
	sort.Ints(sortedArray)
	fmt.Println("input array  :", fmt.Sprint(sortedArray))
	ans := removeDuplicatesFromSortedArray(sortedArray)
	fmt.Println("answer       :", ans)
	fmt.Println("after removal:", fmt.Sprint(sortedArray[0:ans]))
	fmt.Println("LeetCode NO.26 end")
}
