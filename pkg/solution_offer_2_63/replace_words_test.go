package solution_offer_2_63

import "testing"

func TestCase1(t *testing.T) {
	s := "the cattle was rattled by the battery"
	d := []string{"cat", "bat", "rat"}

	expected := "the cat was rat by the bat"
	actual := replaceWords(d, s)

	if expected != actual {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	s := "aadsfasf absbs bbab cadsfafs"
	d := []string{"a", "b", "c"}

	expected := "a a b c"
	actual := replaceWords(d, s)

	if expected != actual {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	s := "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"
	d := []string{"a", "aa", "aaa", "aaaa"}

	expected := "a a a a a a a a bbb baba a"
	actual := replaceWords(d, s)

	if expected != actual {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}
func TestCase4(t *testing.T) {
	s := "the cattle was rattled by the battery"
	d := []string{"catt", "cat", "bat", "rat"}

	expected := "the cat was rat by the bat"
	actual := replaceWords(d, s)

	if expected != actual {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestCase5(t *testing.T) {
	s := "it is abnormal that this solution is accepted"
	d := []string{"ac", "ab"}

	expected := "it is ab that this solution is ac"
	actual := replaceWords(d, s)

	if expected != actual {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}
