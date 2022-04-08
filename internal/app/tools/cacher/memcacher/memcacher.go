package memcacher

import (
	"strconv"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	cacher "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/cacher"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
)

type Memcacher struct {
	client *memcache.Client
}

func NewMemcacher(cfg *config.CachConfig) (*Memcacher, error) {
	client := memcache.New(cfg.Host + ":" + strconv.Itoa(cfg.Port))
	if err := client.Ping(); err != nil {
		return nil, err
	}
	return &Memcacher{
		client: client,
	}, nil
}

func (c *Memcacher) Set(item *cacher.Item) error {
	err := c.client.Set(&memcache.Item{
		Key:        item.Key,
		Value:      item.Value,
		Expiration: item.Expiration,
	})
	if err != nil {
		return servErrors.NewError(servErrors.CACH_ERROR, err.Error())
	}
	return nil
}

func (c *Memcacher) Get(key string) (*cacher.Item, error) {
	item, err := c.client.Get(key)
	if err == memcache.ErrCacheMiss {
		return nil, servErrors.NewError(servErrors.CACH_MISS_CODE, err.Error())
	}
	if err != nil {
		return nil, servErrors.NewError(servErrors.CACH_ERROR, err.Error())
	}
	return cacher.NewItem(item.Key, item.Value, item.Expiration), nil
}

func (c *Memcacher) Delete(key string) error {
	err := c.client.Delete(key)
	if err == memcache.ErrCacheMiss {
		return servErrors.NewError(servErrors.CACH_MISS_CODE, err.Error())
	}
	if err != nil {
		return servErrors.NewError(servErrors.CACH_ERROR, err.Error())
	}
	return nil
}

// func (c *Memcacher) Swap(item *cacher.Item) error {
// 	return c.client.CompareAndSwap(&memcache.Item{
// 		Key:        item.Key,
// 		Value:      item.Value,
// 		Expiration: item.Expiration,
// 	})
// }

// func (c *Memcacher) SetOrReplace(item *cacher.Item) error {
// 	_, err := c.client.Get(item.Key)
// 	if err != nil && err != memcache.ErrCacheMiss {
// 		return errors.Wrap(err, "memcache get item error")
// 	}

// 	if err == nil {
// 		return errors.Wrap(c.client.CompareAndSwap(&memcache.Item{
// 			Key:        item.Key,
// 			Value:      item.Value,
// 			Expiration: item.Expiration,
// 		}), "memcache set item error")
// 	}

// 	return errors.Wrap(c.client.Set(&memcache.Item{
// 		Key:        item.Key,
// 		Value:      item.Value,
// 		Expiration: item.Expiration,
// 	}), "memcache set item error")

// }
