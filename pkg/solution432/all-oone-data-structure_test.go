package solution432

import "testing"

func TestCase1(t *testing.T) {
	key := "key"
	obj := Constructor()
	obj.Inc(key)
	obj.Inc(key)
	param3 := obj.GetMaxKey()
	if param3 != key {
		t.Errorf("Case 1 failed - expected %v, got %v", key, param3)
	}
	param4 := obj.GetMinKey()
	if param4 != key {
		t.Errorf("Case 1 failed - expected %v, got %v", key, param4)
	}

	obj.Inc("key2")
	param5 := obj.GetMaxKey()
	if param5 != "key" {
		t.Errorf("Case 1 failed - expected %v, got %v", "key2", param5)
	}
	param6 := obj.GetMinKey()
	if param6 != "key2" {
		t.Errorf("Case 1 failed - expected %v, got %v", "key2", param6)
	}
}
