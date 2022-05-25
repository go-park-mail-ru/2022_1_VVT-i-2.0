package flashcall

import (
	conf "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	"reflect"
	"testing"
)

func TestFlashcaller_SendCode(t *testing.T) {
	type fields struct {
		apiKey string
		email  string
	}
	type args struct {
		phone string
		code  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Flashcaller{
				apiKey: tt.fields.apiKey,
				email:  tt.fields.email,
			}
			if err := f.SendCode(tt.args.phone, tt.args.code); (err != nil) != tt.wantErr {
				t.Errorf("SendCode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewFlashcaller(t *testing.T) {
	type args struct {
		cfg *conf.NotificatorConfig
	}
	tests := []struct {
		name string
		args args
		want *Flashcaller
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFlashcaller(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFlashcaller() = %v, want %v", got, tt.want)
			}
		})
	}
}
