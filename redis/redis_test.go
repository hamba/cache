package redis_test

import (
	"context"
	"net"
	"testing"

	"github.com/hamba/cache"
	"github.com/hamba/cache/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testRedisServer = "localhost:6379"
	skipRedis       = false
)

func init() {
	c, err := net.Dial("tcp", testRedisServer)
	if err != nil {
		skipRedis = true
		return
	}
	_, _ = c.Write([]byte("SELECT 1\r\n"))
	_, _ = c.Write([]byte("FLUSHDB\r\n"))
	_ = c.Close()
}

func TestRedisCache(t *testing.T) {
	if skipRedis {
		t.Skipf("skipping test; no running server at %s", testRedisServer)
	}

	ctx := context.Background()

	c, err := redis.New("redis://" + testRedisServer + "/1")
	require.NoError(t, err)

	assert.Implements(t, (*cache.Cache)(nil), c)

	// Set
	err = c.Set(ctx, "test", "foobar", 0)
	require.NoError(t, err)

	// Get
	str, err := c.Get(ctx, "test").String()
	require.NoError(t, err)
	assert.Equal(t, "foobar", str)
	_, err = c.Get(ctx, "_").String()
	assert.EqualError(t, err, cache.ErrCacheMiss.Error())

	// Add
	err = c.Add(ctx, "test1", "foobar", 0)
	require.NoError(t, err)
	err = c.Add(ctx, "test1", "foobar", 0)
	assert.EqualError(t, err, cache.ErrNotStored.Error())

	// Replace
	err = c.Replace(ctx, "test1", "foobar", 0)
	require.NoError(t, err)
	err = c.Replace(ctx, "_", "foobar", 0)
	assert.EqualError(t, err, cache.ErrNotStored.Error())

	// GetMulti
	v, err := c.GetMulti(ctx, "test", "test1", "_")
	require.NoError(t, err)
	assert.Len(t, v, 3)
	assert.EqualError(t, v[2].Err, "cache: miss")

	// Delete
	err = c.Delete(ctx, "test1")
	require.NoError(t, err)
	_, err = c.Get(ctx, "test1").String()
	assert.Error(t, err)

	// Inc
	err = c.Set(ctx, "test2", 1, 0)
	require.NoError(t, err)
	i, err := c.Inc(ctx, "test2", 1)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), i)

	// Dec
	err = c.Set(ctx, "test2", 1, 0)
	require.NoError(t, err)
	i, err = c.Dec(ctx, "test2", 1)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), i)
}
