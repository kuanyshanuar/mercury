package redis

import (
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
)

// SetupRedis - setup redis for testing
func SetupRedis() (*redis.Client, error) {
	mr, err := miniredis.Run()
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	return client, nil
}
