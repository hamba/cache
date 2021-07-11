package redis_test

import (
	"context"
	"time"

	"github.com/hamba/cache/redis"
)

func ExampleNew() {
	c, err := redis.New("redis://localhost:6379", redis.WithReadTimeout(10*time.Millisecond))
	if err != nil {
		// Handle error
	}

	i := c.Get(context.Background(), "foobar")
	if i.Err != nil {
		// Handle error
	}

	_, _ = i.Float64()
}
