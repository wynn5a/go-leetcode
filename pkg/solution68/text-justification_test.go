package solution68

import "testing"

func TestCase(t *testing.T) {
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	maxWidth := 16
	expected := []string{
		"This    is    an",
		"example  of text",
		"justification.  ",
	}
	actual := fullJustify(words, maxWidth)
	if len(actual) != len(expected) {
		t.Errorf("Expected %d lines, got %d", len(expected), len(actual))
	}
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected line %d to be '%s', got '%s'", i, expected[i], actual[i])
		}
	}
}

func TestCase2(t *testing.T) {
	words := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	maxWidth := 16
	expected := []string{
		"What   must   be",
		"acknowledgment  ",
		"shall be        ",
	}
	actual := fullJustify(words, maxWidth)
	if len(actual) != len(expected) {
		t.Errorf("Expected %d lines, got %d", len(expected), len(actual))
	}
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected line %d to be '%s', got '%s'", i, expected[i], actual[i])
		}
	}
}

func TestCase3(t *testing.T) {
	words := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain",
		"to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	maxWidth := 20
	expected := []string{
		"Science  is  what we",
		"understand      well",
		"enough to explain to",
		"a  computer.  Art is",
		"everything  else  we",
		"do                  ",
	}
	actual := fullJustify(words, maxWidth)
	if len(actual) != len(expected) {
		t.Errorf("Expected %d lines, got %d", len(expected), len(actual))
	}
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected line %d to be '%s', got '%s'", i, expected[i], actual[i])
		}
	}
}

func TestCase4(t *testing.T) {
	words := []string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country"}
	maxWidth := 16
	expected := []string{
		"ask   not   what", "your country can", "do  for  you ask", "what  you can do", "for your country",
	}
	actual := fullJustify(words, maxWidth)
	if len(actual) != len(expected) {
		t.Errorf("Expected %d lines, got %d", len(expected), len(actual))
	}
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Expected line %d to be '%s', got '%s'", i, expected[i], actual[i])
		}
	}
}
