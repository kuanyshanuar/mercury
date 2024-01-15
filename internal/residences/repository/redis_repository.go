package repository

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"

	"github.com/go-redis/redis/v8"
)

type redisRepository struct {
	client *redis.Client
}

// NewRedisRepository - create a new redis repository
func NewRedisRepository(client *redis.Client) domain.ResidenceRedisRepository {
	return &redisRepository{
		client: client,
	}
}

func (r redisRepository) Get(ctx context.Context, key string) (*domain.Residence, error) {
	//TODO implement me
	panic("implement me")
}

func (r redisRepository) Set(ctx context.Context, key string, seconds int, residence *domain.Residence) error {
	//TODO implement me
	panic("implement me")
}

func (r redisRepository) Delete(ctx context.Context, key string) error {
	//TODO implement me
	panic("implement me")
}
