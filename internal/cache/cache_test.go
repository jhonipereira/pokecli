package cache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("value1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("value2"),
		},
	}

	for _, cas := range cases {
		cache.Add(cas.inputKey, []byte(cas.inputVal))
		actual, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Error(cas.inputKey + " not found")
			continue
		}
		if string(actual) != string(cas.inputVal) {
			t.Error(string(cas.inputVal) + " doesn't  match")
			continue
		}
	}

}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	k1 := "key1"

	cache.Add(k1, []byte("val1"))

	time.Sleep(interval / 2)

	_, ok := cache.Get(k1)
	if !ok {
		t.Errorf("%s should NOT have been reaped", k1)
	}
}
