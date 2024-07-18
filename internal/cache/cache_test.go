package cache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Second)
	if cache.Data == nil {
		t.Error("cache is nil")
	}
}

func TestCacheFunc(t *testing.T) {
	cache := NewCache(time.Second)

	cache.Add("k", []byte("ayaya"))
	actual, ok := cache.Get("k")
	if !ok {
		t.Error("key is not found")
	}
	if string(actual) != "ayaya" {
		t.Error("key is not found")
	}

}
