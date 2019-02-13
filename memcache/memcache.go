// Package redis implements a memcache adapter for github.com/hamba/pkg/cache.
package memcache

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/hamba/pkg/cache"
)

type decoder struct{}

func (d decoder) Bool(v []byte) (bool, error) {
	return string(v) == "1", nil
}

func (d decoder) Int64(v []byte) (int64, error) {
	return strconv.ParseInt(string(v), 10, 64)
}

func (d decoder) Uint64(v []byte) (uint64, error) {
	return strconv.ParseUint(string(v), 10, 64)
}

func (d decoder) Float64(v []byte) (float64, error) {
	return strconv.ParseFloat(string(v), 64)
}

// OptsFunc represents an configuration function for Memcache.
type OptsFunc func(*memcache.Client)

// WithIdleConns configures the Memcache max idle connections.
func WithIdleConns(size int) OptsFunc {
	return func(c *memcache.Client) {
		c.MaxIdleConns = size
	}
}

// WithTimeout configures the Memcache read and write timeout.
func WithTimeout(timeout time.Duration) OptsFunc {
	return func(c *memcache.Client) {
		c.Timeout = timeout
	}
}

// Memcache is a memcache adapter.
type Memcache struct {
	client *memcache.Client

	encoder func(v interface{}) ([]byte, error)
	decoder cache.Decoder
}

// New create a new Memcache instance.
func New(uri string, opts ...OptsFunc) *Memcache {
	c := memcache.New(uri)

	for _, opt := range opts {
		opt(c)
	}

	return &Memcache{
		client:  c,
		encoder: memcacheEncoder,
		decoder: decoder{},
	}
}

// Get gets the item for the given key.
func (c Memcache) Get(key string) *cache.Item {
	b := []byte(nil)
	v, err := c.client.Get(key)
	switch err {
	case memcache.ErrCacheMiss:
		err = cache.ErrCacheMiss
	case nil:
		b = v.Value
	}

	return &cache.Item{
		Decoder: c.decoder,
		Value:   b,
		Err:     err,
	}
}

// GetMulti gets the items for the given keys.
func (c Memcache) GetMulti(keys ...string) ([]*cache.Item, error) {
	val, err := c.client.GetMulti(keys)
	if err != nil {
		return nil, err
	}

	i := []*cache.Item{}
	for _, k := range keys {
		var err = cache.ErrCacheMiss
		var b []byte
		if v, ok := val[k]; ok {
			b = v.Value
			err = nil
		}

		i = append(i, &cache.Item{
			Decoder: c.decoder,
			Value:   b,
			Err:     err,
		})
	}

	return i, nil
}

// Set sets the item in the cache.
func (c Memcache) Set(key string, value interface{}, expire time.Duration) error {
	v, err := c.encoder(value)
	if err != nil {
		return err
	}

	return c.client.Set(&memcache.Item{
		Key:        key,
		Value:      v,
		Expiration: int32(expire.Seconds()),
	})
}

// Add sets the item in the cache, but only if the key does not already exist.
func (c Memcache) Add(key string, value interface{}, expire time.Duration) error {
	v, err := c.encoder(value)
	if err != nil {
		return err
	}

	err = c.client.Add(&memcache.Item{
		Key:        key,
		Value:      v,
		Expiration: int32(expire.Seconds()),
	})
	if err == memcache.ErrNotStored {
		return cache.ErrNotStored
	}
	return err
}

// Replace sets the item in the cache, but only if the key already exists.
func (c Memcache) Replace(key string, value interface{}, expire time.Duration) error {
	v, err := c.encoder(value)
	if err != nil {
		return err
	}

	err = c.client.Replace(&memcache.Item{
		Key:        key,
		Value:      v,
		Expiration: int32(expire.Seconds()),
	})

	if err == memcache.ErrNotStored {
		return cache.ErrNotStored
	}
	return err
}

// Delete deletes the item with the given key.
func (c Memcache) Delete(key string) error {
	return c.client.Delete(key)
}

// Inc increments a key by the value.
func (c Memcache) Inc(key string, value uint64) (int64, error) {
	v, err := c.client.Increment(key, value)
	return int64(v), err
}

// Dec decrements a key by the value.
func (c Memcache) Dec(key string, value uint64) (int64, error) {
	v, err := c.client.Decrement(key, value)
	return int64(v), err
}

func memcacheEncoder(v interface{}) ([]byte, error) {
	switch v.(type) {
	case bool:
		if v.(bool) {
			return []byte("1"), nil
		}
		return []byte("0"), nil
	case int, int8, int16, int32, int64:
		return []byte(fmt.Sprintf("%d", v)), nil
	case uint, uint8, uint16, uint32, uint64:
		return []byte(fmt.Sprintf("%d", v)), nil
	case float32, float64:
		return []byte(fmt.Sprintf("%f", v)), nil
	case string:
		return []byte(v.(string)), nil
	}

	return json.Marshal(v)
}
