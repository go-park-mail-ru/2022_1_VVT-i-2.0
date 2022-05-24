package cacher

import (
	"reflect"
	"testing"
)

func TestNewItem(t *testing.T) {
	type args struct {
		key        string
		value      []byte
		expiration int32
	}
	tests := []struct {
		name string
		args args
		want *Item
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewItem(tt.args.key, tt.args.value, tt.args.expiration); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
