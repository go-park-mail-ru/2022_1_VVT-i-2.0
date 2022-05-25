package memcacher

import (
	"github.com/bradfitz/gomemcache/memcache"
	conf "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/cacher"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewMemcacher(t *testing.T) {
	config := conf.NewAuthMicroserviceConfig()
	response, _ := NewMemcacher(&config.CacherConfig)
	if reflect.DeepEqual(response, nil) {
		t.Errorf("results not match, want %v, have %v", response, nil)
		return
	}
}

func TestNewMemcacherErr(t *testing.T) {
	config := conf.NewAuthMicroserviceConfig()
	_, err := NewMemcacher(&config.CacherConfig)
	assert.Error(t, err)
}

//func TestMemcacher_Set(t *testing.T) {
//	config := conf.NewAuthMicroserviceConfig()
//	memcacher, _ := NewMemcacher(&config.CacherConfig)
//	response := memcacher.Set(&cacher.Item{
//		Key:   "89166152595",
//		Value: []byte("1234"),
//		Expiration: int32(300),
//	})
//	if !reflect.DeepEqual(response, nil) {
//		t.Errorf("results not match, want %v, have %v", response, nil)
//		return
//	}
//}

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
