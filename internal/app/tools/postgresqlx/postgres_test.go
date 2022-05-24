package postgresqlx

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	"github.com/jmoiron/sqlx"
	"reflect"
	"testing"
)

func TestNewPostgresqlX(t *testing.T) {
	type args struct {
		cfg *config.DatabaseConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *sqlx.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPostgresqlX(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPostgresqlX() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPostgresqlX() got = %v, want %v", got, tt.want)
			}
		})
	}
}
