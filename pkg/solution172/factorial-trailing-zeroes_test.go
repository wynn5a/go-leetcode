package solution172

import "testing"

func TestCase(t *testing.T) {
	n := 5
	excepted := 1
	actual := trailingZeroes(n)
	if excepted != actual {
		t.Errorf("excepted: %d, but got: %d", excepted, actual)
	}
}
func TestCase1(t *testing.T) {
	n := 3
	excepted := 0
	actual := trailingZeroes(n)
	if excepted != actual {
		t.Errorf("excepted: %d, but got: %d", excepted, actual)
	}
}
func TestCase2(t *testing.T) {
	n := 0
	excepted := 0
	actual := trailingZeroes(n)
	if excepted != actual {
		t.Errorf("excepted: %d, but got: %d", excepted, actual)
	}
}
func TestCase3(t *testing.T) {
	n := 7
	excepted := 1
	actual := trailingZeroes(n)
	if excepted != actual {
		t.Errorf("excepted: %d, but got: %d", excepted, actual)
	}
}
func TestCase4(t *testing.T) {
	n := 30
	excepted := 7
	actual := trailingZeroes(n)
	if excepted != actual {
		t.Errorf("excepted: %d, but got: %d", excepted, actual)
	}
}
