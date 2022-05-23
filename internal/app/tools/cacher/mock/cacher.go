package mock

// import (
// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/cacher"
// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
// 	"github.com/stretchr/testify/mock"
// )

// type Memcacher struct {
// 	mock.Mock
// }

// func (m Memcacher) Set(item *cacher.Item) error {
// 	if item == nil {
// 		return servErrors.NewError(servErrors.CACH_ERROR, "")
// 	}
// 	return nil
// }

// func (m Memcacher) Get(key string) (*cacher.Item, error) {
// 	if key == "" {
// 		return nil, servErrors.NewError(servErrors.CACH_ERROR, "")
// 	}
// 	cacheItem := &cacher.Item{
// 		Key:        "string",
// 		Value:      []byte("1234"),
// 		Expiration: 1,
// 	}
// 	return cacheItem, nil
// }

// func (m Memcacher) Delete(key string) error {
// 	panic("implement me")
// }
