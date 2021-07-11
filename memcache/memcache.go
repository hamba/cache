// Package memcache implements a memcache adapter for github.com/hamba/pkg/cache.
package memcache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/hamba/cache"
	"github.com/hamba/cache/internal/decoder"
)

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

	enc func(v interface{}) ([]byte, error)
	dec cache.Decoder
}

// New create a new Memcache instance.
func New(uri string, opts ...OptsFunc) *Memcache {
	c := memcache.New(uri)

	for _, opt := range opts {
		opt(c)
	}

	return &Memcache{
		client: c,
		enc:    memcacheEncoder,
		dec:    decoder.StringDecoder{},
	}
}

// Get gets the item for the given key.
func (c Memcache) Get(_ context.Context, key string) cache.Item {
	b := []byte(nil)
	v, err := c.client.Get(key)
	switch {
	case errors.Is(err, memcache.ErrCacheMiss):
		err = cache.ErrCacheMiss
	case err == nil:
		b = v.Value
	}

	return cache.NewItem(c.dec, b, err)
}

// GetMulti gets the items for the given keys.
func (c Memcache) GetMulti(_ context.Context, keys ...string) ([]cache.Item, error) {
	val, err := c.client.GetMulti(keys)
	if err != nil {
		return nil, err
	}

	i := make([]cache.Item, 0, len(keys))
	for _, k := range keys {
		valErr := cache.ErrCacheMiss
		var b []byte
		if v, ok := val[k]; ok {
			b = v.Value
			valErr = nil
		}

		i = append(i, cache.NewItem(c.dec, b, valErr))
	}

	return i, nil
}

// Set sets the item in the cache.
func (c Memcache) Set(_ context.Context, key string, value interface{}, expire time.Duration) error {
	v, err := c.enc(value)
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
func (c Memcache) Add(_ context.Context, key string, value interface{}, expire time.Duration) error {
	v, err := c.enc(value)
	if err != nil {
		return err
	}

	err = c.client.Add(&memcache.Item{
		Key:        key,
		Value:      v,
		Expiration: int32(expire.Seconds()),
	})
	if errors.Is(err, memcache.ErrNotStored) {
		return cache.ErrNotStored
	}
	return err
}

// Replace sets the item in the cache, but only if the key already exists.
func (c Memcache) Replace(_ context.Context, key string, value interface{}, expire time.Duration) error {
	v, err := c.enc(value)
	if err != nil {
		return err
	}

	err = c.client.Replace(&memcache.Item{
		Key:        key,
		Value:      v,
		Expiration: int32(expire.Seconds()),
	})

	if errors.Is(err, memcache.ErrNotStored) {
		return cache.ErrNotStored
	}
	return err
}

// Delete deletes the item with the given key.
func (c Memcache) Delete(_ context.Context, key string) error {
	return c.client.Delete(key)
}

// Inc increments a key by the value.
func (c Memcache) Inc(_ context.Context, key string, value uint64) (int64, error) {
	v, err := c.client.Increment(key, value)
	return int64(v), err
}

// Dec decrements a key by the value.
func (c Memcache) Dec(_ context.Context, key string, value uint64) (int64, error) {
	v, err := c.client.Decrement(key, value)
	return int64(v), err
}

func memcacheEncoder(v interface{}) ([]byte, error) {
	switch val := v.(type) {
	case bool:
		if val {
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
		return []byte(val), nil
	}

	return json.Marshal(v)
}
