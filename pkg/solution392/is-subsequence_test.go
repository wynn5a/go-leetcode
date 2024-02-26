package solution392

import "testing"

func TestCase1(t *testing.T) {
	result := isSubsequence("abc", "ahbgdc")
	if result != true {
		t.Errorf("isSubsequence(abc, ahbgdc) = %t; expect %t", result, true)
	}
}

func TestCase2(t *testing.T) {
	result := isSubsequence("axc", "ahbgdc")
	if result != false {
		t.Errorf("isSubsequence(abc, ahbgdc) = %t; expect %t", result, false)
	}
}
