package redis_test

import (
	"time"

	"github.com/hamba/cache/redis"
)

func ExampleNew() {
	c, err := redis.New("redis://localhost:6379", redis.WithReadTimeout(10*time.Millisecond))
	if err != nil {
		// Handle error
	}

	i := c.Get("foobar")
	if i.Err != nil {
		// Handle error
	}

	i.Float64()
}
