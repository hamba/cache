// Package redis implements a redis adapter for github.com/hamba/pkg/cache.
package redis

import (
	"strconv"
	"time"

	"github.com/go-redis/redis"
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
	client  *redis.Client
	decoder cache.Decoder
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
		client:  c,
		decoder: decoder{},
	}, nil
}

// Get gets the item for the given key.
func (c Redis) Get(key string) *cache.Item {
	b, err := c.client.Get(key).Bytes()
	if err == redis.Nil {
		err = cache.ErrCacheMiss
	}

	return &cache.Item{
		Decoder: c.decoder,
		Value:   b,
		Err:     err,
	}
}

// GetMulti gets the items for the given keys.
func (c Redis) GetMulti(keys ...string) ([]*cache.Item, error) {
	val, err := c.client.MGet(keys...).Result()
	if err != nil {
		return nil, err
	}

	i := []*cache.Item{}
	for _, v := range val {
		var err = cache.ErrCacheMiss
		var b []byte
		if v != nil {
			b = []byte(v.(string))
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
func (c Redis) Set(key string, value interface{}, expire time.Duration) error {
	return c.client.Set(key, value, expire).Err()
}

// Add sets the item in the cache, but only if the key does not already exist.
func (c Redis) Add(key string, value interface{}, expire time.Duration) error {
	if !c.client.SetNX(key, value, expire).Val() {
		return cache.ErrNotStored
	}
	return nil
}

// Replace sets the item in the cache, but only if the key already exists.
func (c Redis) Replace(key string, value interface{}, expire time.Duration) error {
	if !c.client.SetXX(key, value, expire).Val() {
		return cache.ErrNotStored
	}
	return nil
}

// Delete deletes the item with the given key.
func (c Redis) Delete(key string) error {
	return c.client.Del(key).Err()
}

// Inc increments a key by the value.
func (c Redis) Inc(key string, value uint64) (int64, error) {
	return c.client.IncrBy(key, int64(value)).Result()
}

// Dec decrements a key by the value.
func (c Redis) Dec(key string, value uint64) (int64, error) {
	return c.client.DecrBy(key, int64(value)).Result()
}
