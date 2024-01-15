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

// NewIdentityManagerRedisRepository - creates a new repository.
func NewIdentityManagerRedisRepository(client *redis.Client) domain.IdentityManagerRedisRepository {
	return &redisRepository{
		client: client,
	}
}

func (r *redisRepository) Set(
	ctx context.Context,
	key string,
	value *domain.User,
	seconds int,
) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if err = r.client.Set(ctx, key, bytes, time.Second*time.Duration(seconds)).Err(); err != nil {
		return err
	}

	return nil
}

func (r *redisRepository) Get(
	ctx context.Context,
	key string,
) (*domain.User, error) {
	bytes, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	value := &domain.User{}
	if err = json.Unmarshal(bytes, &value); err != nil {
		return nil, err
	}

	return value, nil
}

func (r *redisRepository) Delete(
	ctx context.Context,
	key string,
) error {
	if err := r.client.Del(ctx, key).Err(); err != nil {
		return err
	}

	return nil
}
