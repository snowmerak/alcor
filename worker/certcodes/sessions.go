package certcodes

import (
	"log"

	"github.com/bradfitz/gomemcache/memcache"
)

var store *memcache.Client

func init() {
	store = memcache.New("certcodes:11211")
	if store == nil {
		log.Fatal("cannot connect sessions")
	}
}

func Set(key string, value []byte) error {
	return store.Set(&memcache.Item{
		Key: key, Value: value,
	})
}

func Get(key string) ([]byte, error) {
	item, err := store.Get(key)
	if err != nil {
		return nil, err
	}
	return item.Value, nil
}

func Delete(key string) error {
	return store.Delete(key)
}
