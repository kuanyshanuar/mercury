package transport

import (
	"context"

	apiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type grpcServer struct {
	list   grpctransport.Handler
	update grpctransport.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer.
func NewGRPCServer(endpoints Endpoints, logger log.Logger) apiv1.SystemServiceServer {
	options := helpers.SetupServerOptions(logger)

	return &grpcServer{
		list: grpctransport.NewServer(
			endpoints.ListUsersEndpoint,
			decodeListUsersRequestV1,
			encodeListUsersResponseV1,
			options...,
		),
		update: grpctransport.NewServer(
			endpoints.UpdateUserEndpoint,
			decodeUpdateUserRequestV1,
			encodeUpdateUserResponseV1,
			options...,
		),
	}
}

func (s *grpcServer) ListUsers(ctx context.Context, request *apiv1.ListUsersRequest) (*apiv1.ListUsersResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.list)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.ListUsersResponse), nil
}

func decodeListUsersRequestV1(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.ListUsersRequest)

	return listUsersRequest{
		Criteria: decodeUserSearchCriteriaV1(req.Criteria),
	}, nil
}

func encodeListUsersResponseV1(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(listUsersResponse)
	if resp.Err != nil {
		return &apiv1.ListUsersResponse{}, resp.Err
	}

	return &apiv1.ListUsersResponse{
		Users: encodeUsersV1(resp.Users),
		Total: int64(resp.Total),
	}, nil
}

func (s *grpcServer) UpdateUser(ctx context.Context, request *apiv1.UpdateUserRequest) (*apiv1.UpdateUserResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.update)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.UpdateUserResponse), nil
}

func decodeUpdateUserRequestV1(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.UpdateUserRequest)

	return updateUserRequest{
		UserID: domain.UserID(req.UserId),
		User:   decodeUserV1(req.User),
	}, nil
}

func encodeUpdateUserResponseV1(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(updateUserResponse)
	if resp.Err != nil {
		return &apiv1.UpdateUserResponse{}, resp.Err
	}

	return &apiv1.UpdateUserResponse{}, nil
}
