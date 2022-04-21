package solution824

import (
	"testing"
)

func TestCase1(t *testing.T) {
	s := "I speak Goat Latin"
	actual := toGoatLatin(s)
	expect := "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
	if actual != expect {
		t.Errorf("expect %s, actual %s", expect, actual)
	}
}

func TestCase2(t *testing.T) {
	s := "The quick brown fox jumped over the lazy dog"
	actual := toGoatLatin(s)
	expect := "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"
	if actual != expect {
		t.Errorf("expect %s, actual %s", expect, actual)
	}
}
