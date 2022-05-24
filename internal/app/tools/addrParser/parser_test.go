package addrParser

import (
	"reflect"
	"testing"
)

func TestConcatAddr(t *testing.T) {
	type args struct {
		city   string
		street string
		house  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConcatAddr(tt.args.city, tt.args.street, tt.args.house); got != tt.want {
				t.Errorf("ConcatAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcatAddrToComplete(t *testing.T) {
	type args struct {
		city   string
		street string
		house  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConcatAddrToComplete(tt.args.city, tt.args.street, tt.args.house); got != tt.want {
				t.Errorf("ConcatAddrToComplete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCity(t *testing.T) {
	type args struct {
		city string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCity(tt.args.city); got != tt.want {
				t.Errorf("GetCity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetHouse(t *testing.T) {
	type args struct {
		house string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHouse(tt.args.house); got != tt.want {
				t.Errorf("GetHouse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStreet(t *testing.T) {
	type args struct {
		streetStr string
	}
	tests := []struct {
		name string
		args args
		want *StreetT
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStreet(tt.args.streetStr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStreet() = %v, want %v", got, tt.want)
			}
		})
	}
}
