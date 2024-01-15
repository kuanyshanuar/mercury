package repository

import (
	"context"
	"encoding/json"
	"time"

	"gitlab.com/zharzhanov/mercury/internal/domain"

	"github.com/go-redis/redis/v8"
)

type redisRepository struct {
	client *redis.Client
}

// NewProfileRedisRepository - creates a new redis repository
func NewProfileRedisRepository(
	client *redis.Client,
) domain.ProfileRedisRepository {
	return &redisRepository{
		client: client,
	}
}

func (r *redisRepository) Get(ctx context.Context, key string) (*domain.Profile, error) {
	bytes, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	value := &domain.Profile{}
	if err = json.Unmarshal(bytes, &value); err != nil {
		return nil, err
	}

	return value, nil
}

func (r *redisRepository) Set(ctx context.Context, key string, value *domain.Profile, seconds int) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if err = r.client.Set(ctx, key, bytes, time.Second*time.Duration(seconds)).Err(); err != nil {
		return err
	}

	return nil
}

func (r *redisRepository) Delete(ctx context.Context, key string) error {
	if err := r.client.Del(ctx, key).Err(); err != nil {
		return err
	}

	return nil
}
