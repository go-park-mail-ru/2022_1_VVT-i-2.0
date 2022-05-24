package logger

import (
	"reflect"
	"testing"
	"time"
)

func TestNewServLogger(t *testing.T) {
	type args struct {
		logger Logger
	}
	tests := []struct {
		name string
		args args
		want *ServLogger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServLogger(tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServLogger_Access(t *testing.T) {
	type fields struct {
		Logger Logger
	}
	type args struct {
		requestId     uint64
		method        string
		remoteAddr    string
		url           string
		procesingTime time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = ServLogger{
				Logger: tt.fields.Logger,
			}
		})
	}
}

func TestServLogger_Error(t *testing.T) {
	type fields struct {
		Logger Logger
	}
	type args struct {
		reqId    uint64
		errorMsg string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = ServLogger{
				Logger: tt.fields.Logger,
			}
		})
	}
}
