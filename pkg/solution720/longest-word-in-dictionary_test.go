package solution720

import "testing"

func TestCase1(t *testing.T) {
	output := longestWord([]string{"w", "wo", "wor", "worl", "world"})
	expected := "world"
	if output != expected {
		t.Errorf("Expected %s but got %s", expected, output)
	}
}
func TestCase2(t *testing.T) {
	output := longestWord([]string{"w", "a", "wo", "wor", "worl", "worle", "world"})
	expected := "world"
	if output != expected {
		t.Errorf("Expected %s but got %s", expected, output)
	}
}

func TestCase3(t *testing.T) {
	output := longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"})
	expected := "apple"
	if output != expected {
		t.Errorf("Expected %s but got %s", expected, output)
	}
}
func TestCase4(t *testing.T) {
	output := longestWord([]string{"b", "br", "bre", "brea", "break", "breakf", "breakfa", "breakfas", "breakfast", "l", "lu", "lun", "lunc", "lunch", "d", "di", "din", "dinn", "dinne", "dinner"})
	expected := "breakfast"
	if output != expected {
		t.Errorf("Expected %s but got %s", expected, output)
	}
}
func TestCase5(t *testing.T) {
	output := longestWord([]string{"yo", "ew", "fc", "zrc", "yodn", "fcm", "qm", "qmo", "fcmz", "z", "ewq", "yod", "ewqz", "y"})
	expected := "yodn"
	if output != expected {
		t.Errorf("Expected %s but got %s", expected, output)
	}
}
