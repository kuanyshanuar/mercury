package transport

import (
	"context"

	apiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"

	"gitlab.com/zharzhanov/mercury/internal/helpers"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type grpcServer struct {
	addUserCottage    grpctransport.Handler
	deleteUserCottage grpctransport.Handler
	listUserCottage   grpctransport.Handler
}

// NewGrpcServer - returns grpc server for user cottages
func NewGrpcServer(endpoint Endpoints, logger log.Logger) apiv1.UserCottagesServiceServer {
	options := helpers.SetupServerOptions(logger)
	return &grpcServer{
		addUserCottage: grpctransport.NewServer(
			endpoint.addUserCottage,
			decodeAddUserCottage,
			encodeAddUserCottage,
			options...,
		),
		deleteUserCottage: grpctransport.NewServer(
			endpoint.deleteUserCottage,
			decodeDeleteUserCottage,
			encodeDeleteUserCottage,
			options...,
		),
		listUserCottage: grpctransport.NewServer(
			endpoint.listUserCottage,
			decodeListUserCottage,
			encodeListUserCottage,
			options...,
		),
	}
}
func decodeAddUserCottage(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.AddUserCottageRequest)
	return addUserCottageRequest{
		UserID:    req.UserId,
		CottageID: req.CottageId,
	}, nil
}

func encodeAddUserCottage(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(addUserCottageResponse)
	if resp.Error != nil {
		return &apiv1.AddUserCottageResponse{}, resp.Error
	}
	return &apiv1.AddUserCottageResponse{}, nil
}
func (g *grpcServer) AddUserCottage(ctx context.Context, request *apiv1.AddUserCottageRequest) (*apiv1.AddUserCottageResponse, error) {

	resp, err := helpers.ServeGrpc(ctx, request, g.addUserCottage)
	if err != nil {
		return nil, err
	}
	return resp.(*apiv1.AddUserCottageResponse), nil
}

func decodeDeleteUserCottage(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.DeleteUserCottageRequest)

	return deleteUserCottageRequest{
		CottageID: req.CottageId,
		UserID:    req.UserId,
	}, nil
}

func encodeDeleteUserCottage(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(deleteUserCottageResponse)
	if resp.Error != nil {
		return &apiv1.DeleteUserCottageResponse{}, resp.Error
	}
	return &apiv1.DeleteUserCottageResponse{}, nil
}

func (g *grpcServer) DeleteUserCottage(ctx context.Context, request *apiv1.DeleteUserCottageRequest) (*apiv1.DeleteUserCottageResponse, error) {
	_, err := helpers.ServeGrpc(ctx, request, g.deleteUserCottage)
	if err != nil {
		return nil, err
	}
	return &apiv1.DeleteUserCottageResponse{}, nil
}

func decodeListUserCottage(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.ListUserCottageRequest)

	return listUserCottageRequest{
		UserID:   req.UserId,
		Criteria: decodeFavouriteCottageSearchCriteria(req.Criteria),
	}, nil
}

func encodeListUserCottage(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(listUserCottageResponse)
	if resp.Error != nil {
		return &apiv1.ListUserCottageResponse{
			Total:        int64(resp.Total),
			UserCottages: resp.Result,
		}, resp.Error
	}
	return &apiv1.ListUserCottageResponse{
		Total:        int64(resp.Total),
		UserCottages: resp.Result,
	}, nil
}

func (g *grpcServer) ListUserCottage(ctx context.Context, request *apiv1.ListUserCottageRequest) (*apiv1.ListUserCottageResponse, error) {
	response, err := helpers.ServeGrpc(ctx, request, g.listUserCottage)
	if err != nil {
		return nil, err
	}
	return response.(*apiv1.ListUserCottageResponse), nil
}
