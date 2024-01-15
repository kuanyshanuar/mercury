package repository

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"time"

	"gitlab.com/zharzhanov/mercury/internal/domain"

	"github.com/minio/minio-go/v7"
)

type minioStorage struct {
	minioClient *minio.Client
}

// NewFileStorageRepository - create a new storage repository
func NewFileStorageRepository(minioClient *minio.Client) domain.FileStorage {
	return &minioStorage{
		minioClient: minioClient,
	}
}

func (s *minioStorage) DownloadFile(
	ctx context.Context,
	params *domain.DownloadFileParams,
) (io.ReadCloser, error) {
	err := s.validateDownloadFileParams(params)
	if err != nil {
		return nil, fmt.Errorf("invalid download file params provided: %v", err)
	}

	reqCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	obj, err := s.minioClient.GetObject(reqCtx, params.BucketName, params.ObjectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to download object with id: %s from minio bucket %s. err: %w", params.ObjectName, params.BucketName, err)
	}

	return obj, nil
}

func (s *minioStorage) UploadFile(
	ctx context.Context,
	params *domain.UploadFileParams,
) error {
	err := s.validateUploadFileParams(params)
	if err != nil {
		return fmt.Errorf("invalid upload file params provided: %v", err)
	}

	reqCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	exists, err := s.minioClient.BucketExists(ctx, params.BucketName)
	if err != nil || !exists {
		fmt.Printf("no bucket %s. creating new one...", params.BucketName)
		err := s.minioClient.MakeBucket(ctx, params.BucketName, minio.MakeBucketOptions{})
		if err != nil {
			return fmt.Errorf("failed to create new bucket. err: %w", err)
		}
	}

	_, err = s.minioClient.PutObject(
		reqCtx, params.BucketName, params.FileName, params.Reader, params.FileSize,
		minio.PutObjectOptions{
			UserMetadata: map[string]string{
				"Name": params.FileName,
			},
			ContentType: "image/png",
		})
	if err != nil {
		return fmt.Errorf("failed to upload file. err: %w", err)
	}

	return nil
}

func (s *minioStorage) DeleteFile(
	ctx context.Context,
	params *domain.DeleteFileParams,
) error {
	err := s.validateDeleteFileParams(params)
	if err != nil {
		return fmt.Errorf("invalid delete file params provided: %v", err)
	}

	err = s.minioClient.RemoveObject(ctx, params.BucketName, params.ObjectName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete file. err: %w", err)
	}

	return nil
}

func (s *minioStorage) SignURL(
	ctx context.Context,
	params *domain.SignURLParams,
) (string, error) {
	err := s.validateSignURLParams(params)
	if err != nil {
		return "", fmt.Errorf("invalid sign URL params provided: %v", err)
	}

	reqParams := make(url.Values)
	fileURL, err := s.minioClient.PresignedGetObject(ctx, params.BucketName, params.ObjectName, time.Second*24*60*60, reqParams)
	if err != nil {
		return "", err
	}

	return fileURL.String(), nil
}

func (s *minioStorage) validateDownloadFileParams(params *domain.DownloadFileParams) error {
	if params == nil {
		return fmt.Errorf("nil download file params provided")
	}

	if params.BucketName == "" {
		return fmt.Errorf("empty bucket name provided")
	}

	if params.ObjectName == "" {
		return fmt.Errorf("empty object name provided")
	}

	return nil
}

func (s *minioStorage) validateUploadFileParams(params *domain.UploadFileParams) error {
	if params == nil {
		return fmt.Errorf("nil download file params provided")
	}

	if params.BucketName == "" {
		return fmt.Errorf("empty bucket name provided")
	}

	if params.FileName == "" {
		return fmt.Errorf("empty file name provided")
	}

	if params.FileSize <= 0 {
		return fmt.Errorf("invalid file size provided: %d", params.FileSize)
	}

	if params.Reader == nil {
		return fmt.Errorf("nil reader provided")
	}

	return nil
}

func (s *minioStorage) validateDeleteFileParams(params *domain.DeleteFileParams) error {
	if params == nil {
		return fmt.Errorf("nil download file params provided")
	}

	if params.BucketName == "" {
		return fmt.Errorf("empty bucket name provided")
	}

	if params.ObjectName == "" {
		return fmt.Errorf("empty object name provided")
	}

	return nil
}

func (s *minioStorage) validateSignURLParams(params *domain.SignURLParams) error {
	if params == nil {
		return fmt.Errorf("nil download file params provided")
	}

	if params.BucketName == "" {
		return fmt.Errorf("empty bucket name provided")
	}

	if params.ObjectName == "" {
		return fmt.Errorf("empty object name provided")
	}

	return nil
}
