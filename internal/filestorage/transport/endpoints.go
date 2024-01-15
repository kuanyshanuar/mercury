package transport

import (
	"bytes"
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

// Endpoints collects all of the endpoints that compose an add service. It's meant to
// be used as a helper struct, to collect all of the endpoints into a single
// parameter.
type Endpoints struct {
	UploadFileEndpoint  endpoint.Endpoint
	IsFileExistEndpoint endpoint.Endpoint
}

// NewEndpoint returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func NewEndpoint(service domain.FileStorageService, serviceSecretKey domain.ServiceSecretKey, logger log.Logger) Endpoints {
	factory := func(creator func(domain.FileStorageService) endpoint.Endpoint, logKey string) endpoint.Endpoint {
		return helpers.SetupEndpoint(creator(service), serviceSecretKey, logger, "FileStorageEndpoint", logKey)
	}
	return Endpoints{
		UploadFileEndpoint:  factory(MakeUploadFileEndpoint, "UploadFile"),
		IsFileExistEndpoint: factory(MakeIsFileExistEndpoint, "IsFileExist"),
	}
}

// MakeUploadFileEndpoint Impl.
func MakeUploadFileEndpoint(service domain.FileStorageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(uploadFileRequest)

		// Call the service.
		uploadedURL, err := service.UploadFile(
			ctx,
			req.FileInfo,
			&req.Data,
		)
		return uploadFileResponse{
			FileURL: uploadedURL,
			Err:     err,
		}, nil
	}
}

type uploadFileRequest struct {
	FileInfo *domain.FileInfo
	Data     bytes.Buffer
}

type uploadFileResponse struct {
	FileURL domain.FileURL
	Err     error
}

// MakeIsFileExistEndpoint - Impl.
func MakeIsFileExistEndpoint(service domain.FileStorageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(isFileExistRequest)

		// Call the service.
		isExist, err := service.IsFileExist(
			ctx,
			req.Folder,
			req.FileName,
		)
		return isFileExistResponse{
			IsExist: isExist,
			Err:     err,
		}, nil
	}
}

type isFileExistRequest struct {
	Folder   string
	FileName string
}

type isFileExistResponse struct {
	IsExist bool
	Err     error
}

var (
	_ endpoint.Failer = uploadFileResponse{}
	_ endpoint.Failer = isFileExistResponse{}
)

func (r uploadFileResponse) Failed() error { return r.Err }

func (r isFileExistResponse) Failed() error { return r.Err }
