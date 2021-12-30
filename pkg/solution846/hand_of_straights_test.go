package solution846

import "testing"

func Test_case_1(t *testing.T) {
	var (
		hand      = []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
		groupSize = 3
	)
	actual := isNStraightHand(hand, groupSize)
	if !actual {
		t.Errorf("isNStraightHand(%v, %d) = %v; expected %v", hand, groupSize, actual, true)
	}
}
func Test_case_2(t *testing.T) {
	var (
		hand      = []int{1, 2, 3, 4, 5}
		groupSize = 4
	)
	actual := isNStraightHand(hand, groupSize)
	if actual {
		t.Errorf("isNStraightHand(%v, %d) = %v; expected %v", hand, groupSize, actual, false)
	}
}
func Test_case_3(t *testing.T) {
	var (
		hand      = []int{1, 2, 3, 4, 5, 7, 8, 9}
		groupSize = 4
	)
	actual := isNStraightHand(hand, groupSize)
	if actual {
		t.Errorf("isNStraightHand(%v, %d) = %v; expected %v", hand, groupSize, actual, false)
	}
}
