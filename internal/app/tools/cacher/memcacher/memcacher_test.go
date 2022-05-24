package memcacher

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/cacher"
	"reflect"
	"testing"
)

func TestMemcacher_Delete(t *testing.T) {
	type fields struct {
		client *memcache.Client
	}
	type args struct {
		key string
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
			c := &Memcacher{
				client: tt.fields.client,
			}
			if err := c.Delete(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMemcacher_Get(t *testing.T) {
	type fields struct {
		client *memcache.Client
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *cacher.Item
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Memcacher{
				client: tt.fields.client,
			}
			got, err := c.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemcacher_Set(t *testing.T) {
	type fields struct {
		client *memcache.Client
	}
	type args struct {
		item *cacher.Item
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
			c := &Memcacher{
				client: tt.fields.client,
			}
			if err := c.Set(tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewMemcacher(t *testing.T) {
	type args struct {
		cfg *config.CachConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *Memcacher
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMemcacher(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMemcacher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMemcacher() got = %v, want %v", got, tt.want)
			}
		})
	}
}
