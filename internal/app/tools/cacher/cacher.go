package cacher

type Cacher interface {
	Set(item *Item) error
	Get(key string) (*Item, error)
	Delete(key string) error
}

type Item struct {
	Key   string
	Value []byte
	// Expiration is the cache expiration time, in seconds: either a relative
	// time from now (up to 1 month), or an absolute Unix epoch time.
	// Zero means the Item has no expiration time.
	Expiration int32
}

func NewItem(key string, value []byte, expiration int32) *Item {
	return &Item{
		Key:        key,
		Value:      value,
		Expiration: expiration,
	}
}
