package memcache_test

import (
	"time"

	"github.com/hamba/cache/memcache"
)

func ExampleNew() {
	c := memcache.New("localhost:11211", memcache.WithIdleConns(10), memcache.WithTimeout(10*time.Millisecond))

	i := c.Get("foobar")
	if i.Err != nil {
		// Handle error
	}

	_, _ = i.Float64()
}
