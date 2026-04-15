package pokecache

import (
	"testing"
	"time"
)

// Test file for pokecache package

func TestGet(t *testing.T) {
	newCache := NewCache(3 * time.Second)
	newCache.Add("test123", []byte("testdata"))
	val, present := newCache.Get("test123")
	if !present || string(val) != "testdata" {
		t.Errorf("Added test123 entry with \"testdata\" value, but received %v", val)
	}
}

func TestReap(t *testing.T) {
	newCache := NewCache(3 * time.Second)
	newCache.Add("test123", nil)
	time.Sleep(4 * time.Second)
	_, present := newCache.Get("test123")
	if present {
		t.Errorf("Entry should be deleted, but is present.")
	}
}
