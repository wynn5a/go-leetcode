package solution

import (
	"fmt"
	"math/rand"
	"time"
)

// RotateArray is working as:
// Given an array, rotate the array to the right by k steps, where k is non-negative.
// Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
// Could you do it in-place with O(1) extra space?
func rotateArray(array []int, k int) {
	//solution1(array, k)
	//solution2(array, k)
	solution3(array, k)
}

// Solution1 rotate array step by step, loop k times
// time: O(k*n) & space: O(1)
func solution1(array []int, k int) {
	i, l := 0, len(array)
	for ; i < k; i++ {
		temp := array[l-1]
		for j := 0; j < l; j++ {
			temp, array[j] = array[j], temp
		}
	}
}

// Solution2 rotate array using another array
// time: O(n) & space: O(n)
func solution2(array []int, k int) {
	l := len(array)
	a := make([]int, l)
	for i, e := range array {
		a[(i+k)%l] = e
	}
	copy(array, a)
}

// Solution3 rotate array by reversing it
// time: O(n) & space: O(1)
func solution3(array []int, k int) {
	k %= len(array)
	reverse(array, 0, len(array) - 1)
	reverse(array, 0, k - 1)
	reverse(array, k, len(array)-1)
}

func reverse(a []int, start int, end int) {
	for ; start < end; start, end = start+1, end-1 {
		a[start], a[end] = a[end], a[start]
	}
}

func Run189() {
	fmt.Println("--- LeetCode NO.189 begin")
	input := make([]int, 10)
	rand.Seed(time.Now().UTC().UnixNano())
	for i, _ := range input {
		input[i] = rand.Intn(20)
	}
	fmt.Println("input array      :", fmt.Sprint(input))
	k := 2
	rotateArray(input, k)
	fmt.Printf("answer when k = %v: %v\n", k, fmt.Sprint(input))
	fmt.Println("LeetCode NO.189 end")
}
