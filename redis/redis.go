// Package redis implements a redis adapter for github.com/hamba/pkg/cache.
package redis

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/hamba/cache"
	"github.com/hamba/cache/internal/decoder"
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
func (c Redis) Get(ctx context.Context, key string) cache.Item {
	b, err := c.conn.Get(ctx, key).Bytes()
	if errors.Is(err, redis.Nil) {
		err = cache.ErrCacheMiss
	}

	return cache.NewItem(c.dec, b, err)
}

// GetMulti gets the items for the given keys.
func (c Redis) GetMulti(ctx context.Context, keys ...string) ([]cache.Item, error) {
	val, err := c.conn.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, err
	}

	i := []cache.Item{}
	for _, v := range val {
		valErr := cache.ErrCacheMiss
		var b []byte
		if v != nil {
			b = []byte(v.(string))
			valErr = nil
		}

		i = append(i, cache.NewItem(c.dec, b, valErr))
	}

	return i, nil
}

// Set sets the item in the cache.
func (c Redis) Set(ctx context.Context, key string, value interface{}, expire time.Duration) error {
	return c.conn.Set(ctx, key, value, expire).Err()
}

// Add sets the item in the cache, but only if the key does not already exist.
func (c Redis) Add(ctx context.Context, key string, value interface{}, expire time.Duration) error {
	if !c.conn.SetNX(ctx, key, value, expire).Val() {
		return cache.ErrNotStored
	}
	return nil
}

// Replace sets the item in the cache, but only if the key already exists.
func (c Redis) Replace(ctx context.Context, key string, value interface{}, expire time.Duration) error {
	if !c.conn.SetXX(ctx, key, value, expire).Val() {
		return cache.ErrNotStored
	}
	return nil
}

// Delete deletes the item with the given key.
func (c Redis) Delete(ctx context.Context, key string) error {
	return c.conn.Del(ctx, key).Err()
}

// Inc increments a key by the value.
func (c Redis) Inc(ctx context.Context, key string, value uint64) (int64, error) {
	return c.conn.IncrBy(ctx, key, int64(value)).Result()
}

// Dec decrements a key by the value.
func (c Redis) Dec(ctx context.Context, key string, value uint64) (int64, error) {
	return c.conn.DecrBy(ctx, key, int64(value)).Result()
}
