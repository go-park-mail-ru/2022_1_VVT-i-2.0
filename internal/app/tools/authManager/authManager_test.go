package authManager

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"reflect"
	"testing"
)

func TestMapToTokenPayload(t *testing.T) {
	type args struct {
		payloadMap map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want *TokenPayload
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapToTokenPayload(tt.args.payloadMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapToTokenPayload() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTokenPayload(t *testing.T) {
	type args struct {
		id models.UserId
	}
	tests := []struct {
		name string
		args args
		want *TokenPayload
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTokenPayload(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTokenPayload() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenPayloadToMap(t *testing.T) {
	type args struct {
		payload TokenPayload
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TokenPayloadToMap(tt.args.payload); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TokenPayloadToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
