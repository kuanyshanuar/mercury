package domain

import (
	"context"
	"io"
)

// FileURL - file location.
type FileURL string

// FileInfo - describes a file.
type FileInfo struct {
	Folder   string `json:"folder"`
	FileName string `json:"filename"`
	Override bool   `json:"override"`
	FileSize int64  `json:"file_size"`
}

// FileStorageService - provides access to a business logic.
type FileStorageService interface {

	// UploadFile - uploads file to the storage.
	UploadFile(
		ctx context.Context,
		fileInfo *FileInfo,
		data io.Reader,
	) (FileURL, error)

	// IsFileExist - checks if file exists in the storage
	IsFileExist(
		ctx context.Context,
		folder string,
		fileName string,
	) (bool, error)
}
