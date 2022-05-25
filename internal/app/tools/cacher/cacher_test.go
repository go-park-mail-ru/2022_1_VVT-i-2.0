package cacher

import (
	"reflect"
	"testing"
)

func TestNewItem(t *testing.T) {
	response := NewItem("1", []byte("1"), 1)
	expect := &Item{
		Key: "1",
		Value: []byte("1"),
		Expiration: 1,
	}
	if !reflect.DeepEqual(response, expect) {
		t.Errorf("results not match, want %v, have %v", response, expect)
		return
	}
}
