package pokecache

import (
	"reflect"
	"testing"
	"time"
)

func TestAddition(t *testing.T) {
	c := NewCache(time.Hour)
	val := []byte{123, 255}
	c.Add("key", val)
	gotten, _ := c.Get("key")
	if !reflect.DeepEqual(gotten, val) {
		t.Fatalf("Expected %v, got %v", val, gotten)
	}
}

func TestDNE(t *testing.T) {
	c := NewCache(time.Hour)
	c.Add("key", []byte{})
	_, exists := c.Get("asd")
	if exists {
		t.Fatal("key \"asd\" should not exist")
	}
}

func TestReap(t *testing.T) {
	c := NewCache(time.Millisecond * 500)
	c.Add("key", []byte{1, 2, 3})
	time.Sleep(time.Second)
	_, exists := c.Get("key")
	if exists {
		t.Fatal("Cache not cleared after 2* the given interval")
	}
}
