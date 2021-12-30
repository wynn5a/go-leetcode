package solution846

import (
	"fmt"
	"sort"
)

//Alice 手中有一把牌，她想要重新排列这些牌，分成若干组，使每一组的牌数都是 groupSize ，并且由 groupSize 张连续的牌组成。
//给你一个整数数组 hand 其中 hand[i] 是写在第 i 张牌，和一个整数 groupSize 。如果她可能重新排列这些牌，返回 true ；否则，返回 false

//记录每个数字出现的次数，提前结束判断
func isNStraightHand(hand []int, groupSize int) bool {
	l := len(hand)
	if l < groupSize {
		return false
	}
	if l%groupSize != 0 {
		return false
	}

	if groupSize == 1 {
		return true
	}

	sort.Ints(hand)
	count := make(map[int]int, len(hand))
	for _, h := range hand {
		count[h]++
	}

	for _, h := range hand {
		//被挑完的，不能作为每组的第一个
		if count[h] == 0 {
			continue
		}
		for num := h; num < h+groupSize; num++ {
			//没有发现必须的数字
			if count[num] == 0 {
				return false
			}
			count[num]--
		}
	}
	return true

}

//模拟，每次抽 SIZE 张出来
//func isNStraightHand(hand []int, groupSize int) bool {
//	l := len(hand)
//	if l < groupSize {
//		return false
//	}
//	if l%groupSize != 0 {
//		return false
//	}
//
//	if groupSize == 1 {
//		return true
//	}
//	sort.Ints(hand)
//	fmt.Printf("sorted: %v \n", hand)
//
//	hand = group(hand, groupSize)
//
//	return len(hand) == 0
//}

//每次都抽出 size 张牌，直到没牌，或者牌不够
func group(hand []int, size int) []int {
	l := len(hand)
	if l == 0 || l < size {
		return hand
	}
	var handNew []int
	n := hand[0]
	s := 1
	for i := 1; i < l; i++ {
		if s < size && hand[i] == n+1 {
			s++
			n++
		} else {
			handNew = append(handNew, hand[i])
		}
	}
	fmt.Printf("hand new: %v \n", handNew)
	if s != size {
		return handNew
	}

	hand = group(handNew, size)
	return hand
}
