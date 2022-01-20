package solution2029

import "testing"

func Test_case_1(t *testing.T) {
	var (
		stones = []int{2, 1}
	)
	actual := stoneGameIX(stones)
	if !actual {
		t.Errorf("stoneGameIX(%v) = %v; expected %v", stones, actual, true)
	}
}
func Test_case_2(t *testing.T) {
	var (
		stones = []int{2}
	)
	actual := stoneGameIX(stones)
	if actual {
		t.Errorf("stoneGameIX(%v) = %v; expected %v", stones, actual, false)
	}
}
func Test_case_3(t *testing.T) {
	var (
		stones = []int{5, 1, 2, 4, 3}
	)
	actual := stoneGameIX(stones)
	if actual {
		t.Errorf("stoneGameIX(%v) = %v; expected %v", stones, actual, false)
	}
}
