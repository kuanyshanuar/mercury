package transport

import (
	"bytes"
	"context"
	"io"

	apiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"
	"gitlab.com/zharzhanov/mercury/internal/helpers"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type grpcServerV1 struct {
	uploadFile  grpctransport.Handler
	isFileExist grpctransport.Handler
}

type apiUploadFileRequest struct {
	server apiv1.FileStorageService_UploadFileServer
}

// NewGRPCServerV1 makes a set of endpoints available as a gRPC AddServer.
func NewGRPCServerV1(endpoints Endpoints, logger log.Logger) apiv1.FileStorageServiceServer {
	options := helpers.SetupServerOptions(logger)
	return &grpcServerV1{
		uploadFile: grpctransport.NewServer(
			endpoints.UploadFileEndpoint,
			decodeUploadFileRequestV1,
			encodeUploadFileResponseV1,
			options...,
		),
		isFileExist: grpctransport.NewServer(
			endpoints.IsFileExistEndpoint,
			decodeIsFileExistRequestV1,
			encodeIsFileExistResponseV1,
			options...,
		),
	}
}

// UploadFile Impl.
func (s *grpcServerV1) UploadFile(server apiv1.FileStorageService_UploadFileServer) error {
	rep, err := helpers.ServeGrpc(server.Context(), &apiUploadFileRequest{server: server}, s.uploadFile)
	if err != nil {
		return err
	}
	err = server.SendAndClose(rep.(*apiv1.UploadFileResponse))
	if err != nil {
		return errors.GRPCErrorEncoder(err)
	}
	return nil
}

func decodeUploadFileRequestV1(_ context.Context, grpcReq interface{}) (interface{}, error) {
	server := grpcReq.(*apiUploadFileRequest).server

	// Receiving upload header.
	req, err := server.Recv()
	if err == io.EOF {
		return nil, errors.NewErrInvalidArgument("unexpected end of stream")
	}
	if err != nil {
		return nil, err
	}

	// Get file info.
	fileInfo := req.GetFileInfo()
	if fileInfo == nil {
		return nil, errors.NewErrInvalidArgument("file information missing in request")
	}

	var buf bytes.Buffer

	// Start upload process.
	fileSize := 0
	//goland:noinspection GoLinterLocal
	for {
		// Receiving data.
		req, err := server.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		chunk := req.GetChunkData()

		// Check file size.
		fileSize += len(chunk)
		// if fileSize > maxFileSize {
		// 	return nil, errors.NewErrInvalidArgument("file is too large")
		// }

		// Write data.
		buf.Grow(len(chunk))
		_, err = buf.Write(chunk)
		if err != nil {
			return nil, err
		}
	}

	// Return decoded request.
	return uploadFileRequest{
		FileInfo: decodeFileInfo(fileInfo, int64(fileSize)),
		Data:     buf,
	}, nil
}

func encodeUploadFileResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(uploadFileResponse)
	if resp.Err != nil {
		return nil, resp.Err
	}
	return &apiv1.UploadFileResponse{
		FileUrl: string(resp.FileURL),
	}, nil
}

func decodeFileInfo(fileInfo *apiv1.FileInfo, fileSize int64) *domain.FileInfo {
	if fileInfo == nil {
		return nil
	}
	return &domain.FileInfo{
		Folder:   fileInfo.Folder,
		FileName: fileInfo.Filename,
		Override: fileInfo.Override,
		FileSize: fileSize,
	}
}

func (s *grpcServerV1) IsFileExist(ctx context.Context, request *apiv1.IsFileExistRequest) (*apiv1.IsFileExistResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.isFileExist)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.IsFileExistResponse), nil
}

func decodeIsFileExistRequestV1(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.IsFileExistRequest)
	return isFileExistRequest{
		Folder:   req.Folder,
		FileName: req.FileName,
	}, nil
}

func encodeIsFileExistResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(isFileExistResponse)
	if resp.Err != nil {
		return &apiv1.IsFileExistResponse{}, resp.Err
	}
	return &apiv1.IsFileExistResponse{
		IsExist: resp.IsExist,
	}, nil
}
