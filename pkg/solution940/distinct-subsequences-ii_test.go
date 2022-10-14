package solution940

import "testing"

func TestCase1(t *testing.T) {
	s := "abc"
	expect := 7

	actual := distinctSubseqII(s)
	if actual != expect {
		t.Errorf("Test failed, actual: %d, expected: %d", actual, expect)
	}
}

func TestCase2(t *testing.T) {
	s := "aba"
	expect := 6

	actual := distinctSubseqII(s)
	if actual != expect {
		t.Errorf("Test failed, actual: %d, expected: %d", actual, expect)
	}
}

func TestCase3(t *testing.T) {
	s := "aaa"
	expect := 3

	actual := distinctSubseqII(s)
	if actual != expect {
		t.Errorf("Test failed, actual: %d, expected: %d", actual, expect)
	}
}
