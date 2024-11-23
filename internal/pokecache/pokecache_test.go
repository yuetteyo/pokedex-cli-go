package pokecache

import "testing"

func TestCreateCashe(t *testing.T) {
	cache := NewCache()
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCashe(t *testing.T) {
	cache := NewCache()

	cache.Add("key1", []byte("val1"))
	actual, ok := cache.Get("key1")
	if !ok {
		t.Errorf("key1 not found")
	}
	if string(actual) != "val1" {
		t.Errorf("value doesnt't match")
	}
}
