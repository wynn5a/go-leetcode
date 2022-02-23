package solution917

import "testing"

func TestCase1(t *testing.T) {
	s := "ab-cd"
	expected := "dc-ba"
	actual := reverseOnlyLetters(s)
	if expected != actual {
		t.Errorf("reverseOnlyLetters('%s')='%s', expected '%s'", s, actual, expected)
	}
}
func TestCase2(t *testing.T) {
	s := "a-bC-dEf-ghIj"
	expected := "j-Ih-gfE-dCba"
	actual := reverseOnlyLetters(s)
	if expected != actual {
		t.Errorf("reverseOnlyLetters('%s')='%s', expected '%s'", s, actual, expected)
	}
}
func TestCase3(t *testing.T) {
	s := "Test1ng-Leet=code-Q!"
	expected := "Qedo1ct-eeLg=ntse-T!"
	actual := reverseOnlyLetters(s)
	if expected != actual {
		t.Errorf("reverseOnlyLetters('%s')='%s', expected '%s'", s, actual, expected)
	}
}
