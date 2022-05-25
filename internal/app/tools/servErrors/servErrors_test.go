package servErrors

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestErrorAs(t *testing.T) {
	type args struct {
		e error
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrorAs(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrorAs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Cause(t *testing.T) {
	errI := &Error{"",1}
	response := errI.Cause()
	assert.Error(t, response)
}

func TestError_Error(t *testing.T) {
	errI := &Error{"",1}
	response := errI.Error()
	expect := "error with code 1 description: "
	if !reflect.DeepEqual(response, expect) {
		t.Errorf("results not match, want %v, have %v", response, expect)
		return
	}
}

func TestNewError(t *testing.T) {
	response := NewError(1,"")
	expect := Error{
		Code:        1,
		Description: "",
	}
	if !reflect.DeepEqual(response, expect) {
		t.Errorf("results not match, want %v, have %v", response, expect)
		return
	}
}
