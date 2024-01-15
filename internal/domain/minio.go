package domain

import (
	"context"
	"io"
)

// DownloadFileParams - download file params
type DownloadFileParams struct {
	BucketName string
	ObjectName string
}

// UploadFileParams - upload file params
type UploadFileParams struct {
	BucketName string
	FileName   string
	FileSize   int64
	Reader     io.Reader
}

// DeleteFileParams - delete file params
type DeleteFileParams struct {
	BucketName string
	ObjectName string
}

// SignURLParams - sign url params
type SignURLParams struct {
	BucketName string
	ObjectName string
}

// FileStorage - provides access to a storage.
type FileStorage interface {
	DownloadFile(ctx context.Context, params *DownloadFileParams) (io.ReadCloser, error)
	UploadFile(ctx context.Context, params *UploadFileParams) error
	DeleteFile(ctx context.Context, params *DeleteFileParams) error
	SignURL(ctx context.Context, params *SignURLParams) (string, error)
}
