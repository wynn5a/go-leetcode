package solution398

import "testing"

func TestCase(t *testing.T) {
	nums := []int{1, 2, 3, 3, 3}
	obj := Constructor(nums)

	r := obj.Pick(3)
	if r == 0 || r == 1 {
		t.Errorf("Case %v, expect %v, but got %d", 1, []int{2, 3, 4}, r)
	}
}
