package solution75

import "testing"

func Test(t *testing.T) {
	// Test Case 1: Merging two equal-length strings
	result := mergeAlternately("abc", "123")
	expected := "a1b2c3"
	if result != expected {
		t.Errorf("Expected: %s, but got: %s", expected, result)
	}

	result = mergeAlternately("hello", "world")
	expected = "hweolrllod"
	if result != expected {
		t.Errorf("Expected: %s, but got: %s", expected, result)
	}

	result = mergeAlternately("1", "xyz")
	expected = "1xyz"
	if result != expected {
		t.Errorf("Expected: %s, but got: %s", expected, result)
	}

	result = mergeAlternately("abcdefgh", "123")
	expected = "a1b2c3defgh"
	if result != expected {
		t.Errorf("Expected: %s, but got: %s", expected, result)
	}
}
