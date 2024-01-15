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
	createLeadBuilder grpctransport.Handler
	listLeadBuilders  grpctransport.Handler
	updateLeadBuilder grpctransport.Handler
	deleteLeadBuilder grpctransport.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer.
func NewGRPCServer(endpoints Endpoints, logger log.Logger) apiv1.LeadBuildersServiceServer {
	options := helpers.SetupServerOptions(logger)

	return &grpcServer{
		createLeadBuilder: grpctransport.NewServer(
			endpoints.CreateLeadBuilderEndpoint,
			decodeCreateLeadBuilderRequest,
			encodeCreateLeadBuilderResponse,
			options...,
		),
		listLeadBuilders: grpctransport.NewServer(
			endpoints.ListLeadBuildersEndpoint,
			decodeListLeadBuildersRequest,
			encodeListLeadBuildersResponse,
			options...,
		),
		updateLeadBuilder: grpctransport.NewServer(
			endpoints.UpdateLeadBuilderEndpoint,
			decodeUpdateLeadBuilderRequest,
			encodeUpdateLeadBuildersResponse,
			options...,
		),
		deleteLeadBuilder: grpctransport.NewServer(
			endpoints.DeleteLeadBuilderEndpoint,
			decodeDeleteLeadBuilderRequest,
			encodeDeleteLeadBuildersResponse,
			options...,
		),
	}
}

func (s *grpcServer) CreateLeadBuilder(ctx context.Context, request *apiv1.CreateLeadBuilderRequest) (*apiv1.CreateLeadBuilderResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.createLeadBuilder)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.CreateLeadBuilderResponse), nil
}

func decodeCreateLeadBuilderRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.CreateLeadBuilderRequest)

	return creteLeadBuildersRequest{
		Lead: decodeLeadBuilderV1(req.Lead),
	}, nil
}

func encodeCreateLeadBuilderResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(creteLeadBuildersResponse)
	if resp.Err != nil {
		return &apiv1.CreateLeadBuilderResponse{}, resp.Err
	}
	return &apiv1.CreateLeadBuilderResponse{}, nil
}

func (s *grpcServer) ListLeadBuilders(ctx context.Context, request *apiv1.ListLeadBuildersRequest) (*apiv1.ListLeadBuildersResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.listLeadBuilders)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.ListLeadBuildersResponse), nil
}

func decodeListLeadBuildersRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.ListLeadBuildersRequest)
	return listLeadBuildersRequest{
		Criteria: decodeLeadBuildersSearchCriteria(req.Criteria),
	}, nil
}

func encodeListLeadBuildersResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(listLeadBuildersResponse)
	if resp.Err != nil {
		return &apiv1.ListLeadBuildersResponse{}, resp.Err
	}
	return &apiv1.ListLeadBuildersResponse{
		Leads: encodeLeadBuildersV1(resp.Leads),
		Total: int64(resp.Total),
	}, nil
}

func (s *grpcServer) UpdateLeadBuilder(ctx context.Context, request *apiv1.UpdateLeadBuilderRequest) (*apiv1.UpdateLeadBuilderResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.updateLeadBuilder)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.UpdateLeadBuilderResponse), nil
}

func decodeUpdateLeadBuilderRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.UpdateLeadBuilderRequest)
	return updateLeadBuilderRequest{
		LeadID: domain.LeadID(req.LeadId),
		Lead:   decodeLeadBuilderV1(req.Lead),
	}, nil
}

func encodeUpdateLeadBuildersResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(updateLeadBuilderResponse)
	if resp.Err != nil {
		return &apiv1.UpdateLeadBuilderResponse{}, resp.Err
	}
	return &apiv1.UpdateLeadBuilderResponse{}, nil
}

func (s *grpcServer) DeleteLeadBuilder(ctx context.Context, request *apiv1.DeleteLeadBuilderRequest) (*apiv1.DeleteLeadBuilderResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.deleteLeadBuilder)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.DeleteLeadBuilderResponse), nil
}

func decodeDeleteLeadBuilderRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.DeleteLeadBuilderRequest)
	return deleteLeadBuilderRequest{
		LeadID: domain.LeadID(req.LeadId),
	}, nil
}

func encodeDeleteLeadBuildersResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(deleteLeadBuilderResponse)
	if resp.Err != nil {
		return &apiv1.DeleteLeadBuilderResponse{}, resp.Err
	}
	return &apiv1.DeleteLeadBuilderResponse{}, nil
}
