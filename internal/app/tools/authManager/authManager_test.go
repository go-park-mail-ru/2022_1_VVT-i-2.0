package authManager

import (
	"reflect"
	"testing"
)

func TestNewTokenPayload(t *testing.T) {
	response := NewTokenPayload(1)
	expect := &TokenPayload{Id: 1}
	if !reflect.DeepEqual(response, expect) {
		t.Errorf("results not match, want %v, have %v", response, expect)
		return
	}
}

func TestNewTokenPayloadErr(t *testing.T) {
	response := NewTokenPayload(1)
	expect := &TokenPayload{}
	if reflect.DeepEqual(response, expect) {
		t.Errorf("results not match, want %v, have %v", response, expect)
		return
	}
}

func TestTokenPayloadToMap(t *testing.T) {
	dataTest := NewTokenPayload(1)
	response := TokenPayloadToMap(*dataTest)
	expectTest := NewTokenPayload(1)
	expect := map[string]interface{}{
		idTitle:      expectTest.Id,
		expiresTitle: expectTest.Exp,
	}
	if !reflect.DeepEqual(response, expect) {
		t.Errorf("results not match, want %v, have %v", response, expect)
		return
	}
}

func TestTokenPayloadToMapErr(t *testing.T) {
	dataTest := NewTokenPayload(1)
	response := TokenPayloadToMap(*dataTest)
	expectTest := NewTokenPayload(2)
	expect := map[string]interface{}{
		idTitle:      expectTest.Id,
		expiresTitle: expectTest.Exp,
	}
	if reflect.DeepEqual(response, expect) {
		t.Errorf("results not match, want %v, have %v", response, expect)
		return
	}
}

func TestMapToTokenPayloadErr(t *testing.T) {
	dataTest := NewTokenPayload(1)
	data := map[string]interface{}{
		idTitle:      dataTest.Id,
		expiresTitle: dataTest.Exp,
	}
	response := MapToTokenPayload(data)
	expect := &TokenPayload{Id: 1}
	if reflect.DeepEqual(response, expect) {
		t.Errorf("results not match, want %v, have %v", response, expect)
		return
	}
}
