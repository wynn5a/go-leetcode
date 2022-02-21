package solution838

import "testing"

func TestCase1(t *testing.T) {
	dominoes := ".L.R...LR..L.."
	expected := "LL.RR.LLRRLL.."
	result := pushDominoes(dominoes)
	if result != expected {
		t.Errorf("pushDominoes(%s)=%s, expected %s", dominoes, result, expected)
	}
}

func TestCase2(t *testing.T) {
	dominoes := "RR.L"
	expected := "RR.L"
	result := pushDominoes(dominoes)
	if result != expected {
		t.Errorf("pushDominoes(%s)=%s, expected %s", dominoes, result, expected)
	}
}

func TestCase3(t *testing.T) {
	dominoes := "..R.."
	expected := "..RRR"
	result := pushDominoes(dominoes)
	if result != expected {
		t.Errorf("pushDominoes(%s)=%s, expected %s", dominoes, result, expected)
	}
}
