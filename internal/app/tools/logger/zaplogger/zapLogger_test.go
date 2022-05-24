package zaplogger

import (
	conf "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	"go.uber.org/zap"
	"reflect"
	"testing"
)

func TestNewZapLogger(t *testing.T) {
	type args struct {
		cfg *conf.LogConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *zap.SugaredLogger
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewZapLogger(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewZapLogger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewZapLogger() got = %v, want %v", got, tt.want)
			}
		})
	}
}
