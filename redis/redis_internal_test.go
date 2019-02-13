package redis

import (
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func TestWithPoolSize(t *testing.T) {
	o := &redis.Options{}

	WithPoolSize(12)(o)

	assert.Equal(t, 12, o.PoolSize)
}

func TestWithPoolTimeout(t *testing.T) {
	o := &redis.Options{}

	WithPoolTimeout(time.Second)(o)

	assert.Equal(t, time.Second, o.PoolTimeout)
}

func TestWithReadTimeout(t *testing.T) {
	o := &redis.Options{}

	WithReadTimeout(time.Second)(o)

	assert.Equal(t, time.Second, o.ReadTimeout)
}

func TestWithWriteTimeout(t *testing.T) {
	o := &redis.Options{}

	WithWriteTimeout(time.Second)(o)

	assert.Equal(t, time.Second, o.WriteTimeout)
}

func TestNewRedis(t *testing.T) {
	c, err := New("redis://test", WithPoolSize(12))

	assert.NoError(t, err)
	assert.Equal(t, 12, c.client.Options().PoolSize)
}

func TestNewRedis_InvalidUri(t *testing.T) {
	_, err := New("test")
	assert.Error(t, err)
}
