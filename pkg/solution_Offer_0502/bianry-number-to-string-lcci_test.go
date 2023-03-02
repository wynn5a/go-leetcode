package solution_Offer_0502

import "testing"

func TestCase1(t *testing.T) {
	d := 0.625
	expected := "0.101"
	actual := printBin(d)
	if expected != actual {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	d := 0.1
	expected := "ERROR"
	actual := printBin(d)
	if expected != actual {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}
