package redis

import (
	"fmt"
	"time"

	"gitlab.com/zharzhanov/mercury/internal/domain"

	"github.com/go-redis/redis/v8"
)

// NewRedisClient - Returns new redis client
func NewRedisClient(cfg domain.AppConfig) (*redis.Client, error) {

	if len(cfg.RedisHost) == 0 {
		return nil, fmt.Errorf("redis host required")
	}
	if len(cfg.RedisPassword) == 0 {
		return nil, fmt.Errorf("redis password required")
	}
	if len(cfg.RedisPort) == 0 {
		return nil, fmt.Errorf("redis port required")
	}

	redisHost := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)

	client := redis.NewClient(&redis.Options{
		Addr:         redisHost,
		MinIdleConns: 200,
		PoolSize:     cfg.RedisPoolSize,
		PoolTimeout:  time.Duration(240) * time.Second,
		Password:     cfg.RedisPassword,
		DB:           0,
	})

	return client, nil
}
