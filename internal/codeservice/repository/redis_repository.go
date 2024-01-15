package repository

import (
	"context"
	"encoding/json"
	"time"

	"gitlab.com/zharzhanov/mercury/internal/domain"

	"github.com/go-redis/redis/v8"
)

type repository struct {
	client *redis.Client
}

// NewRepository - creates a new repository
func NewRepository(client *redis.Client) domain.CodeRedisRepository {
	return &repository{
		client: client,
	}
}

func (r *repository) Set(ctx context.Context, key string, value *domain.Code, seconds int) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if err = r.client.Set(ctx, key, bytes, time.Second*time.Duration(seconds)).Err(); err != nil {
		return err
	}

	return nil
}

func (r *repository) Get(ctx context.Context, key string) (*domain.Code, error) {
	bytes, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	value := &domain.Code{}
	if err = json.Unmarshal(bytes, &value); err != nil {
		return nil, err
	}

	return value, nil
}

func (r *repository) Delete(ctx context.Context, key string) error {
	if err := r.client.Del(ctx, key).Err(); err != nil {
		return err
	}

	return nil
}
