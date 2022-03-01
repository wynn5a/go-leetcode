package solution6

import "testing"

func TestCase1(t *testing.T) {
	s := "PAYPALISHIRING"
	numRows := 3
	expected := "PAHNAPLSIIGYIR"
	actual := convert(s, numRows)
	if actual != expected {
		t.Errorf("convert(%s, %d) = '%s', expected: '%s'", s, numRows, actual, expected)
	}
}

func TestCase2(t *testing.T) {
	s := "PAYPALISHIRING"
	numRows := 4
	expected := "PINALSIGYAHRPI"
	actual := convert(s, numRows)
	if actual != expected {
		t.Errorf("convert(%s, %d) = '%s', expected: '%s'", s, numRows, actual, expected)
	}
}
func TestCase3(t *testing.T) {
	s := "A"
	numRows := 1
	expected := "A"
	actual := convert(s, numRows)
	if actual != expected {
		t.Errorf("convert(%s, %d) = '%s', expected: '%s'", s, numRows, actual, expected)
	}
}
