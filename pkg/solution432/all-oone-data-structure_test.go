package solution432

import "testing"

func TestCase1(t *testing.T) {
	key := "key"
	obj := Constructor()
	obj.Inc(key)
	obj.Inc(key)
	param_3 := obj.GetMaxKey()
	if param_3 != key {
		t.Errorf("Case 1 failed - expected %v, got %v", key, param_3)
	}
	param_4 := obj.GetMinKey()
	if param_4 != key {
		t.Errorf("Case 1 failed - expected %v, got %v", key, param_4)
	}

	obj.Inc("key2")
	param_5 := obj.GetMaxKey()
	if param_5 != "key" {
		t.Errorf("Case 1 failed - expected %v, got %v", "key2", param_5)
	}
	param_6 := obj.GetMinKey()
	if param_6 != "key2" {
		t.Errorf("Case 1 failed - expected %v, got %v", "key2", param_6)
	}
}
