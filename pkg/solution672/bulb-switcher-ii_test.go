package solution672

import "testing"

func TestCase1(t *testing.T) {
	n := 1
	p := 1

	act := flipLights(n, p)

	if act != 2 {
		t.Errorf("Test case 1 failed, expected: 2, actual: %d", act)
	}
}

func TestCase2(t *testing.T) {
	n := 2
	p := 1

	act := flipLights(n, p)

	if act != 3 {
		t.Errorf("Test case 2 failed, expected: 3, actual: %d", act)
	}
}

func TestCase3(t *testing.T) {
	n := 3
	p := 1

	act := flipLights(n, p)

	if act != 4 {
		t.Errorf("Test case 3 failed, expected: 4, actual: %d", act)
	}
}
