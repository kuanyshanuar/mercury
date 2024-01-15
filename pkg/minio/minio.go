package minio

import (
	"fmt"

	"gitlab.com/zharzhanov/mercury/internal/domain"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// NewStorageClient - creates connection to minio storage.
func NewStorageClient(cfg domain.AppConfig) (*minio.Client, error) {

	if len(cfg.MinioEndpoint) == 0 {
		return nil, fmt.Errorf("minio endpoint required")
	}
	if len(cfg.MinioAccessKeyID) == 0 {
		return nil, fmt.Errorf("minio access key id required")
	}
	if len(cfg.MinioSecretAccessKey) == 0 {
		return nil, fmt.Errorf("minio secret access key id required")
	}

	minioClient, err := minio.New(cfg.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinioAccessKeyID, cfg.MinioSecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, fmt.Errorf("could not create minio client. Err: %v", err)
	}

	return minioClient, nil
}
