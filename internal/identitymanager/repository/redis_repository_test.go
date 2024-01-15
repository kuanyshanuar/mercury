package repository

import (
	"context"
	"gitlab.com/zharzhanov/mercury/internal/helpers"
	"testing"

	redisv8 "github.com/go-redis/redis/v8"
	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/pkg/database/redis"
)

func TestRedisRepository_Set(t *testing.T) {

	// Setup redis
	//
	client, err := redis.SetupRedis()
	if err != nil {
		t.Fatal(err)
	}

	var (
		ctx = context.Background()
	)

	testRepository := NewIdentityManagerRedisRepository(client)

	// Define tests
	//
	tests := []struct {
		name        string
		key         string
		value       *domain.User
		expectError bool
	}{
		{
			name: "Success: data is stored",
			key:  "1234",
			value: &domain.User{
				RoleID:                  1,
				FirstName:               "test",
				LastName:                "test",
				Email:                   "test@mail.com",
				City:                    "test",
				Phone:                   "test",
				ConsultationPhoneNumber: "test",
				Image:                   "test",
				Password:                "test",
				IsVerified:              false,
				IsBanned:                helpers.PointerBool(false),
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := testRepository.Set(ctx, test.key, test.value, domain.RedisUserDuration)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestRedisRepository_Get(t *testing.T) {

	// Setup redis
	//
	client, err := redis.SetupRedis()
	if err != nil {
		t.Fatal(err)
	}

	var (
		ctx = context.Background()
	)

	testRepository := NewIdentityManagerRedisRepository(client)

	// Define tests
	//
	tests := []struct {
		name        string
		key         string
		value       *domain.User
		expectError bool
	}{
		{
			name: "Success: data is stored",
			key:  "1",
			value: &domain.User{
				RoleID:                  1,
				FirstName:               "test",
				LastName:                "test",
				Email:                   "test@mail.com",
				City:                    "test",
				Phone:                   "test",
				ConsultationPhoneNumber: "test",
				Image:                   "test",
				Password:                "test",
				IsVerified:              false,
				IsBanned:                helpers.PointerBool(false),
			},
			expectError: false,
		},
		{
			name: "Success: data is stored",
			key:  "2",
			value: &domain.User{
				RoleID:                  1,
				FirstName:               "test",
				LastName:                "test",
				Email:                   "test@mail.com",
				City:                    "test",
				Phone:                   "test",
				ConsultationPhoneNumber: "test",
				Image:                   "test",
				Password:                "test",
				IsVerified:              false,
				IsBanned:                helpers.PointerBool(false),
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := testRepository.Set(ctx, test.key, test.value, domain.RedisUserDuration)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}

				created, err := testRepository.Get(ctx, test.key)
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}
				if created == nil {
					t.Error("expected error: no user")
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}

func TestRedisRepository_Delete(t *testing.T) {

	// Setup redis
	//
	client, err := redis.SetupRedis()
	if err != nil {
		t.Fatal(err)
	}

	var (
		ctx = context.Background()
	)

	testRepository := NewIdentityManagerRedisRepository(client)

	// Define tests
	//
	tests := []struct {
		name        string
		key         string
		value       *domain.User
		expectError bool
	}{
		{
			name: "Success: data is stored",
			key:  "1",
			value: &domain.User{
				RoleID:                  1,
				FirstName:               "test",
				LastName:                "test",
				Email:                   "test@mail.com",
				City:                    "test",
				Phone:                   "test",
				ConsultationPhoneNumber: "test",
				Image:                   "test",
				Password:                "test",
				IsVerified:              false,
				IsBanned:                helpers.PointerBool(false),
			},
			expectError: false,
		},
		{
			name: "Success: data is stored",
			key:  "2",
			value: &domain.User{
				RoleID:                  1,
				FirstName:               "test",
				LastName:                "test",
				Email:                   "test@mail.com",
				City:                    "test",
				Phone:                   "test",
				ConsultationPhoneNumber: "test",
				Image:                   "test",
				Password:                "test",
				IsVerified:              false,
				IsBanned:                helpers.PointerBool(false),
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := testRepository.Set(ctx, test.key, test.value, domain.RedisUserDuration)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}

				err := testRepository.Delete(ctx, test.key)
				if err != nil {
					t.Errorf("unexpected error %s", err)
				}

				found, err := testRepository.Get(ctx, test.key)
				if err != redisv8.Nil {
					t.Errorf("unexpected error %s", err)
				}
				if found != nil {
					t.Error("expected error: no user")
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}
