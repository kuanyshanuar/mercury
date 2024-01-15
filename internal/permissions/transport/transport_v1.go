package transport

import (
	"context"
	imsapiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/identityserviceapi/v1"
	"gitlab.com/zharzhanov/mercury/internal/domain"

	"gitlab.com/zharzhanov/mercury/internal/helpers"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type grpcServer struct {
	allow grpctransport.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer.
func NewGRPCServer(endpoints Endpoints, logger log.Logger) imsapiv1.PermissionServiceServer {
	options := helpers.SetupServerOptions(logger)

	return &grpcServer{
		allow: grpctransport.NewServer(
			endpoints.AllowPermissionEndpoint,
			decodeAllowPermissionRequest,
			encodeAllowPermissionRequest,
			options...,
		),
	}
}

func (s *grpcServer) Allow(ctx context.Context, request *imsapiv1.AllowPermissionRequest) (*imsapiv1.AllowPermissionResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.allow)
	if err != nil {
		return nil, err
	}
	return rep.(*imsapiv1.AllowPermissionResponse), nil
}

func decodeAllowPermissionRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*imsapiv1.AllowPermissionRequest)

	return allowPermissionRequest{
		PermissionKey: req.PermissionKey,
		UserID:        domain.UserID(req.UserId),
		RoleID:        domain.RoleID(req.RoleId),
	}, nil
}

func encodeAllowPermissionRequest(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(allowPermissionResponse)
	if resp.Err != nil {
		return nil, resp.Err
	}

	return &imsapiv1.AllowPermissionResponse{
		IsAllowed: resp.isAllowed,
	}, nil
}

func (s *grpcServer) List(ctx context.Context, request *imsapiv1.ListPermissionsRequest) (*imsapiv1.ListPermissionsResponse, error) {
	//TODO implement me
	return nil, nil
}
