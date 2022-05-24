package servErrors

import (
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
	type fields struct {
		Description string
		Code        int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Error{
				Description: tt.fields.Description,
				Code:        tt.fields.Code,
			}
			if err := e.Cause(); (err != nil) != tt.wantErr {
				t.Errorf("Cause() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestError_Error(t *testing.T) {
	type fields struct {
		Description string
		Code        int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Error{
				Description: tt.fields.Description,
				Code:        tt.fields.Code,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewError(t *testing.T) {
	type args struct {
		eCode  int
		eDescr string
	}
	tests := []struct {
		name string
		args args
		want Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewError(tt.args.eCode, tt.args.eDescr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewError() = %v, want %v", got, tt.want)
			}
		})
	}
}
