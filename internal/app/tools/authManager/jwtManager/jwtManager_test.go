package jwtManager

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/authManager"
	"github.com/golang-jwt/jwt"
	"reflect"
	"testing"
	"time"
)

func TestJwtManager_CreateToken(t *testing.T) {
	type fields struct {
		key         []byte
		method      jwt.SigningMethod
		expDuration time.Duration
	}
	type args struct {
		payload *authManager.TokenPayload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manager := &JwtManager{
				key:         tt.fields.key,
				method:      tt.fields.method,
				expDuration: tt.fields.expDuration,
			}
			got, err := manager.CreateToken(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJwtManager_GetEpiryTime(t *testing.T) {
	type fields struct {
		key         []byte
		method      jwt.SigningMethod
		expDuration time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Duration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manager := &JwtManager{
				key:         tt.fields.key,
				method:      tt.fields.method,
				expDuration: tt.fields.expDuration,
			}
			if got := manager.GetEpiryTime(); got != tt.want {
				t.Errorf("GetEpiryTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJwtManager_ParseToken(t *testing.T) {
	type fields struct {
		key         []byte
		method      jwt.SigningMethod
		expDuration time.Duration
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *authManager.TokenPayload
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manager := &JwtManager{
				key:         tt.fields.key,
				method:      tt.fields.method,
				expDuration: tt.fields.expDuration,
			}
			got, err := manager.ParseToken(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewJwtManager(t *testing.T) {
	type args struct {
		cfg config.AuthManagerConfig
	}
	tests := []struct {
		name string
		args args
		want *JwtManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJwtManager(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJwtManager() = %v, want %v", got, tt.want)
			}
		})
	}
}
