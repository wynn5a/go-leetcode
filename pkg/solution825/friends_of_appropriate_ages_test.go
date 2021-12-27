package solution825

import "testing"

func Test_case_1(t *testing.T) {
	var (
		a        = []int{16, 16}
		expected = 2
	)
	actual := numFriendRequests(a)
	if actual != expected {
		t.Errorf("numFriendRequests(%v) = %d; expected %d", a, actual, expected)
	}
}

func Test_case_2(t *testing.T) {
	var (
		a        = []int{16, 17, 18}
		expected = 2
	)
	actual := numFriendRequests(a)
	if actual != expected {
		t.Errorf("numFriendRequests(%v) = %d; expected %d", a, actual, expected)
	}
}
func Test_case_3(t *testing.T) {
	var (
		a        = []int{20, 30, 100, 110, 120}
		expected = 3
	)
	actual := numFriendRequests(a)
	if actual != expected {
		t.Errorf("numFriendRequests(%v) = %d; expected %d", a, actual, expected)
	}
}
