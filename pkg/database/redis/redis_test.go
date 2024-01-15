package redis

import (
	"gitlab.com/zharzhanov/mercury/internal/domain"
	"testing"
)

func TestNewRedisClient(t *testing.T) {

	type arguments struct {
		cfg domain.AppConfig
	}

	tests := []struct {
		name        string
		arguments   arguments
		expectError bool
	}{
		{
			name: "Fail: invalid cfg",
			arguments: arguments{
				cfg: domain.AppConfig{
					RedisHost:     "",
					RedisPort:     "",
					RedisPassword: "",
				},
			},
			expectError: true,
		},
		{
			name: "Fail: invalid endpoint",
			arguments: arguments{
				cfg: domain.AppConfig{
					RedisHost:     "",
					RedisPort:     "test",
					RedisPassword: "test",
				},
			},
			expectError: true,
		},
		{
			name: "Fail: invlid access key id",
			arguments: arguments{
				cfg: domain.AppConfig{
					RedisHost:     "test",
					RedisPort:     "",
					RedisPassword: "test",
				},
			},
			expectError: true,
		},
		{
			name: "Fail: invalid access key",
			arguments: arguments{
				cfg: domain.AppConfig{
					RedisHost:     "test",
					RedisPort:     "test",
					RedisPassword: "",
				},
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.arguments

			_, err := NewRedisClient(args.cfg)
			if !test.expectError {
				if err != nil {
					t.Errorf("unexpected error: %s", err)
				}
			} else {
				if err == nil {
					t.Error("expected error but got nothing")
				}
			}
		})
	}
}
