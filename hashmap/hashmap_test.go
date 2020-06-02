package hashmap

import "testing"

func TestHashMap(t *testing.T) {
	hashMap := NewHashMap(100)
	if hashMap.Get("name") != nil {
		t.Error("should be nil")
	}
	hashMap.Set("name", "peter")
	hashMap.Set("job", "dev")
	node := hashMap.Get("name")
	if node == nil || node.val != "peter" {
		t.Error("name should be petter")
	}
}
