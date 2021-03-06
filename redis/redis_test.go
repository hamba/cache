package redis_test

import (
	"net"
	"testing"

	"github.com/hamba/cache/redis"
	"github.com/hamba/pkg/cache"
	"github.com/stretchr/testify/assert"
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

	c, err := redis.New("redis://" + testRedisServer + "/1")
	assert.NoError(t, err)

	// Set
	err = c.Set("test", "foobar", 0)
	assert.NoError(t, err)

	// Get
	str, err := c.Get("test").String()
	assert.NoError(t, err)
	assert.Equal(t, "foobar", str)
	_, err = c.Get("_").String()
	assert.EqualError(t, err, cache.ErrCacheMiss.Error())

	// Add
	err = c.Add("test1", "foobar", 0)
	assert.NoError(t, err)
	err = c.Add("test1", "foobar", 0)
	assert.EqualError(t, err, cache.ErrNotStored.Error())

	// Replace
	err = c.Replace("test1", "foobar", 0)
	assert.NoError(t, err)
	err = c.Replace("_", "foobar", 0)
	assert.EqualError(t, err, cache.ErrNotStored.Error())

	// GetMulti
	v, err := c.GetMulti("test", "test1", "_")
	assert.NoError(t, err)
	assert.Len(t, v, 3)
	assert.EqualError(t, v[2].Err, "cache: miss")

	// Delete
	err = c.Delete("test1")
	assert.NoError(t, err)
	_, err = c.Get("test1").String()
	assert.Error(t, err)

	// Inc
	err = c.Set("test2", 1, 0)
	assert.NoError(t, err)
	i, err := c.Inc("test2", 1)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), i)

	// Dec
	err = c.Set("test2", 1, 0)
	assert.NoError(t, err)
	i, err = c.Dec("test2", 1)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), i)
}
