package minio

import (
	"testing"

	"gitlab.com/zharzhanov/mercury/internal/domain"
)

func TestNewStorageClient(t *testing.T) {

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
					MinioEndpoint:        "",
					MinioAccessKeyID:     "",
					MinioSecretAccessKey: "",
				},
			},
			expectError: true,
		},
		{
			name: "Fail: invalid endpoint",
			arguments: arguments{
				cfg: domain.AppConfig{
					MinioEndpoint:        "",
					MinioAccessKeyID:     "test",
					MinioSecretAccessKey: "test",
				},
			},
			expectError: true,
		},
		{
			name: "Fail: invlid access key id",
			arguments: arguments{
				cfg: domain.AppConfig{
					MinioEndpoint:        "test",
					MinioAccessKeyID:     "",
					MinioSecretAccessKey: "test",
				},
			},
			expectError: true,
		},
		{
			name: "Fail: invalid access key",
			arguments: arguments{
				cfg: domain.AppConfig{
					MinioEndpoint:        "test",
					MinioAccessKeyID:     "test",
					MinioSecretAccessKey: "",
				},
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.arguments

			_, err := NewStorageClient(args.cfg)
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
