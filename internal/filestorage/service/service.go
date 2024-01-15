package service

import (
	"context"
	"fmt"
	"io"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"

	"github.com/go-kit/log"
	"github.com/google/uuid"
)

type service struct {
	storage    domain.FileStorage
	bucketName string
	baseURL    string
}

// NewService creates a new service
func NewService(
	storage domain.FileStorage,
	bucketName string,
	baseURL string,
	logger log.Logger,
) domain.FileStorageService {
	var service domain.FileStorageService
	{
		service = newBasicService(
			storage,
			bucketName,
			baseURL,
		)
		service = loggingServiceMiddleware(logger)(service)
	}
	return service
}

func newBasicService(
	storage domain.FileStorage,
	bucketName string,
	baseURL string,
) domain.FileStorageService {
	return &service{
		storage:    storage,
		bucketName: bucketName,
		baseURL:    baseURL,
	}
}

func (s *service) UploadFile(
	ctx context.Context,
	fileInfo *domain.FileInfo,
	data io.Reader,
) (domain.FileURL, error) {

	// Validate inputs.
	if fileInfo == nil {
		return "", errors.NewErrInvalidArgument("file info required")
	}
	if len(fileInfo.Folder) == 0 {
		return "", errors.NewErrInvalidArgument("folder required")
	}
	if len(fileInfo.FileName) == 0 {
		return "", errors.NewErrInvalidArgument("filename required")
	}

	randomFileName, err := s.assertGenerateRandomName()
	if err != nil {
		return "", errors.NewErrInternal(
			fmt.Sprintf("cannot generate name for image: %v", err),
		)
	}

	fullFileName := fmt.Sprintf("%s/%s", fileInfo.Folder, randomFileName)

	uploadFileParams := &domain.UploadFileParams{
		BucketName: s.bucketName,
		FileName:   fullFileName,
		Reader:     data,
		FileSize:   fileInfo.FileSize,
	}

	err = s.storage.UploadFile(ctx, uploadFileParams)
	if err != nil {
		return "", err
	}

	return domain.FileURL(fmt.Sprintf("%s/%s/%s", s.baseURL, s.bucketName, fullFileName)), nil
}

func (s *service) assertGenerateRandomName() (string, error) {
	randomName, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return randomName.String(), nil
}

func (s *service) IsFileExist(
	ctx context.Context,
	folder string,
	fileName string,
) (bool, error) {
	return false, nil
}
