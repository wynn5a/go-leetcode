package solution39

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7
	expected := [][]int{{2, 2, 3}, {7}}

	result := combinationSum(candidates, target)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCase2(t *testing.T) {
	candidates := []int{1, 2, 3}
	target := 3
	expected := [][]int{{1, 1, 1}, {1, 2}, {3}}

	result := combinationSum(candidates, target)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCase3(t *testing.T) {
	candidates := []int{2, 3, 5}
	target := 8
	expected := [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}

	result := combinationSum(candidates, target)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCase4(t *testing.T) {
	candidates := []int{8, 7, 4, 3}
	target := 11
	expected := [][]int{{8, 3}, {7, 4}, {4, 4, 3}}

	result := combinationSum(candidates, target)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
