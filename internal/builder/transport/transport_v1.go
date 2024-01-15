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
	create grpctransport.Handler
	list   grpctransport.Handler
	get    grpctransport.Handler
	update grpctransport.Handler
	delete grpctransport.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer.
func NewGRPCServer(endpoints Endpoints, logger log.Logger) apiv1.BuilderServiceServer {
	options := helpers.SetupServerOptions(logger)

	return &grpcServer{
		create: grpctransport.NewServer(
			endpoints.CreateBuilderEndpoint,
			decodeCreateBuilderRequestV1,
			encodeCreateBuilderResponseV1,
			options...,
		),
		list: grpctransport.NewServer(
			endpoints.ListBuildersEndpoint,
			decodeListBuildersRequestV1,
			encodeListBuildersResponseV1,
			options...,
		),
		get: grpctransport.NewServer(
			endpoints.GetBuilderEndpoint,
			decodeGetBuilderRequestV1,
			encodeGetBuilderResponseV1,
			options...,
		),
		update: grpctransport.NewServer(
			endpoints.UpdateBuilderEndpoint,
			decodeUpdateBuilderRequestV1,
			encodeUpdateBuilderResponseV1,
			options...,
		),
		delete: grpctransport.NewServer(
			endpoints.DeleteBuilderEndpoint,
			decodeDeleteBuilderRequestV1,
			encodeDeleteBuilderResponseV1,
			options...,
		),
	}
}

// CreateBuilder implements v1.BuilderServiceServer
func (s *grpcServer) CreateBuilder(ctx context.Context, request *apiv1.CreateBuilderRequest) (*apiv1.CreateBuilderResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.create)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.CreateBuilderResponse), nil
}

func decodeCreateBuilderRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.CreateBuilderRequest)

	return createBuilderRequest{
		Builder: decodeBuilderV1(req.Builder),
	}, nil
}

func encodeCreateBuilderResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(createBuilderResponse)
	if resp.Err != nil {
		return &apiv1.CreateBuilderResponse{}, resp.Err
	}
	return &apiv1.CreateBuilderResponse{
		Id: int64(resp.BuilderID),
	}, nil
}

// GetBuilder implements v1.BuilderServiceServer
func (s *grpcServer) GetBuilder(ctx context.Context, request *apiv1.GetBuilderRequest) (*apiv1.GetBuilderResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.get)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.GetBuilderResponse), nil
}

func encodeGetBuilderResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(getBuilderResponse)
	if resp.Err != nil {
		return &apiv1.GetBuilderResponse{}, resp.Err
	}

	return &apiv1.GetBuilderResponse{
		Builder: encodeBuilderV1(resp.Builder),
	}, nil
}

func decodeGetBuilderRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.GetBuilderRequest)

	return getBuilderRequest{
		BuilderID: domain.BuilderID(req.Id),
	}, nil
}

// ListBuilders implements v1.BuilderServiceServer
func (s *grpcServer) ListBuilders(ctx context.Context, request *apiv1.ListBuildersRequest) (*apiv1.ListBuildersResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.list)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.ListBuildersResponse), nil
}

func encodeListBuildersResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(listBuilderResponse)
	if resp.Err != nil {
		return &apiv1.ListBuildersResponse{}, resp.Err
	}

	return &apiv1.ListBuildersResponse{
		Builders: encodeBuildersV1(resp.Builders),
		Total:    int64(resp.Total),
	}, nil
}

func decodeListBuildersRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.ListBuildersRequest)

	return listBuilderRequest{
		Criteria: decodeBuilderSearchCriteriaV1(req.Criteria),
	}, nil
}

// UpdateBuilder implements v1.BuilderServiceServer
func (s *grpcServer) UpdateBuilder(ctx context.Context, request *apiv1.UpdateBuilderRequest) (*apiv1.UpdateBuilderResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.update)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.UpdateBuilderResponse), nil
}

func encodeUpdateBuilderResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(updateBuilderResponse)
	if resp.Err != nil {
		return &apiv1.UpdateBuilderResponse{}, resp.Err
	}

	return &apiv1.UpdateBuilderResponse{}, nil
}

func decodeUpdateBuilderRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.UpdateBuilderRequest)

	return updateBuilderRequest{
		BuilderID: domain.BuilderID(req.Id),
		Builder:   decodeBuilderV1(req.Builder),
	}, nil
}

// DeleteBuilder implements v1.BuilderServiceServer
func (s *grpcServer) DeleteBuilder(ctx context.Context, request *apiv1.DeleteBuilderRequest) (*apiv1.DeleteBuilderResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.delete)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.DeleteBuilderResponse), nil
}

func decodeDeleteBuilderRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.DeleteBuilderRequest)

	return deleteBuilderRequest{
		BuilderID: domain.BuilderID(req.Id),
	}, nil
}

func encodeDeleteBuilderResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(deleteBuilderResponse)
	if resp.Err != nil {
		return &apiv1.DeleteBuilderResponse{}, resp.Err
	}

	return &apiv1.DeleteBuilderResponse{}, nil
}
