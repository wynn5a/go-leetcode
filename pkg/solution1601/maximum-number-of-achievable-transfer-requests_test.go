package solution1601

import "testing"

func TestCase1(t *testing.T) {
	n := 5
	requests := [][]int{{0, 1}, {1, 0}, {0, 1}, {1, 2}, {2, 0}, {3, 4}}
	expected := 5
	actual := maximumRequests(n, requests)
	if expected != actual {
		t.Errorf("maximumRequests(%d, %v) = %d: expected %d", n, requests, actual, expected)
	}
}
func TestCase2(t *testing.T) {
	n := 3
	requests := [][]int{{0, 0}, {1, 2}, {2, 1}}
	expected := 3
	actual := maximumRequests(n, requests)
	if expected != actual {
		t.Errorf("maximumRequests(%d, %v) = %d: expected %d", n, requests, actual, expected)
	}
}
