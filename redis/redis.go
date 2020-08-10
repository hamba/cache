// Package redis implements a redis adapter for github.com/hamba/pkg/cache.
package redis

import (
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/hamba/cache/internal/decoder"
	"github.com/hamba/pkg/cache"
)

// OptsFunc represents an configuration function for Redis.
type OptsFunc func(*redis.Options)

// WithPoolSize configures the Redis pool size.
func WithPoolSize(size int) OptsFunc {
	return func(o *redis.Options) {
		o.PoolSize = size
	}
}

// WithPoolTimeout configures the Redis pool timeout.
func WithPoolTimeout(timeout time.Duration) OptsFunc {
	return func(o *redis.Options) {
		o.PoolTimeout = timeout
	}
}

// WithReadTimeout configures the Redis read timeout.
func WithReadTimeout(timeout time.Duration) OptsFunc {
	return func(o *redis.Options) {
		o.ReadTimeout = timeout
	}
}

// WithWriteTimeout configures the Redis write timeout.
func WithWriteTimeout(timeout time.Duration) OptsFunc {
	return func(o *redis.Options) {
		o.WriteTimeout = timeout
	}
}

// Redis is a redis adapter.
type Redis struct {
	conn *redis.Client
	dec  cache.Decoder
}

// New create a new Redis instance.
func New(uri string, opts ...OptsFunc) (*Redis, error) {
	o, err := redis.ParseURL(uri)
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		opt(o)
	}

	c := redis.NewClient(o)

	return &Redis{
		conn: c,
		dec:  decoder.StringDecoder{},
	}, nil
}

// Get gets the item for the given key.
func (c Redis) Get(key string) cache.Item {
	b, err := c.conn.Get(key).Bytes()
	if err == redis.Nil {
		err = cache.ErrCacheMiss
	}

	return cache.NewItem(c.dec, b, err)
}

// GetMulti gets the items for the given keys.
func (c Redis) GetMulti(keys ...string) ([]cache.Item, error) {
	val, err := c.conn.MGet(keys...).Result()
	if err != nil {
		return nil, err
	}

	i := []cache.Item{}
	for _, v := range val {
		var err = cache.ErrCacheMiss
		var b []byte
		if v != nil {
			b = []byte(v.(string))
			err = nil
		}

		i = append(i, cache.NewItem(c.dec, b, err))
	}

	return i, nil
}

// Set sets the item in the cache.
func (c Redis) Set(key string, value interface{}, expire time.Duration) error {
	return c.conn.Set(key, value, expire).Err()
}

// Add sets the item in the cache, but only if the key does not already exist.
func (c Redis) Add(key string, value interface{}, expire time.Duration) error {
	if !c.conn.SetNX(key, value, expire).Val() {
		return cache.ErrNotStored
	}
	return nil
}

// Replace sets the item in the cache, but only if the key already exists.
func (c Redis) Replace(key string, value interface{}, expire time.Duration) error {
	if !c.conn.SetXX(key, value, expire).Val() {
		return cache.ErrNotStored
	}
	return nil
}

// Delete deletes the item with the given key.
func (c Redis) Delete(key string) error {
	return c.conn.Del(key).Err()
}

// Inc increments a key by the value.
func (c Redis) Inc(key string, value uint64) (int64, error) {
	return c.conn.IncrBy(key, int64(value)).Result()
}

// Dec decrements a key by the value.
func (c Redis) Dec(key string, value uint64) (int64, error) {
	return c.conn.DecrBy(key, int64(value)).Result()
}
