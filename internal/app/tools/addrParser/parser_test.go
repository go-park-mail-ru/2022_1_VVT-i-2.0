package addrParser

import (
	"reflect"
	"testing"
)

func TestGetCity(t *testing.T) {
	response := GetCity("")
	if !reflect.DeepEqual(response, "") {
		t.Errorf("results not match, want %v, have %v", response, "")
		return
	}
}

func TestGetStreet(t *testing.T) {
	response := GetStreet("")
	expect := &StreetT{}
	if !reflect.DeepEqual(response, expect) {
		t.Errorf("results not match, want %v, have %v", response, expect)
		return
	}
}

func TestGetStreetErr(t *testing.T) {
	response := GetStreet("")
	if reflect.DeepEqual(response, "") {
		t.Errorf("results not match, want %v, have %v", response, "")
		return
	}
}

func TestGetHouse(t *testing.T) {
	response := GetHouse("")
	if !reflect.DeepEqual(response, "") {
		t.Errorf("results not match, want %v, have %v", response, "")
		return
	}
}

func TestConcatAddr(t *testing.T) {
	response := ConcatAddr("1", "2", "3")
	if !reflect.DeepEqual(response, "1, 2, 3") {
		t.Errorf("results not match, want %v, have %v", response, "1, 2, 3")
		return
	}
}

func TestConcatAddrErr(t *testing.T) {
	response := ConcatAddr("1", "2", "3")
	if reflect.DeepEqual(response, "1,2,3") {
		t.Errorf("results not match, want %v, have %v", response, "1,2,3")
		return
	}
}

func TestConcatAddrToComplete(t *testing.T) {
	response := ConcatAddrToComplete("1", "2", "3")
	if !reflect.DeepEqual(response, "1, 2, 3") {
		t.Errorf("results not match, want %v, have %v", response, "1, 2, 3")
		return
	}
}

func TestConcatAddrToComplete_WithOutHouse(t *testing.T) {
	response := ConcatAddrToComplete("1", "2", "")
	if !reflect.DeepEqual(response, "1, 2") {
		t.Errorf("results not match, want %v, have %v", response, "1, 2")
		return
	}
}

func TestConcatAddrToComplete_WithOutHouseErr(t *testing.T) {
	response := ConcatAddrToComplete("1", "2", "")
	if reflect.DeepEqual(response, "1, 2, 3") {
		t.Errorf("results not match, want %v, have %v", response, "1, 2, 3")
		return
	}
}
