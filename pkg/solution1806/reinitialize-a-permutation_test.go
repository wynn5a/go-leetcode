package solution1806

import "testing"

func TestCase1(t *testing.T) {
	step := reinitializePermutation(2)

	if step != 1 {
		t.Errorf("Test case 1 failed, expected 1, actual: %d \n", step)
	}
}

func TestCase2(t *testing.T) {
	step := reinitializePermutation(4)

	if step != 2 {
		t.Errorf("Test case 2 failed, expected 1, actual: %d \n", step)
	}
}

func TestCase3(t *testing.T) {
	step := reinitializePermutation(6)

	if step != 4 {
		t.Errorf("Test case 3 failed, expected 1, actual: %d \n", step)
	}
}
