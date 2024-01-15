package transport

import (
	"context"

	apiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"

	"gitlab.com/zharzhanov/mercury/internal/helpers"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type grpcServer struct {
	createLeadCottage grpctransport.Handler
	listLeadCottages  grpctransport.Handler
	updateLeadCottage grpctransport.Handler
	deleteLeadCottage grpctransport.Handler
	getLeadCottage    grpctransport.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer.
func NewGRPCServer(endpoints Endpoints, logger log.Logger) apiv1.LeadCottageServiceServer {
	options := helpers.SetupServerOptions(logger)

	return &grpcServer{
		createLeadCottage: grpctransport.NewServer(
			endpoints.CreateLeadCottageEndpoint,
			decodeCreateLeadCottageRequest,
			encodeCreateLeadCottageResponse,
			options...,
		),
		listLeadCottages: grpctransport.NewServer(
			endpoints.ListLeadCottagesEndpoint,
			decodeListLeadCottagesRequest,
			encodeListLeadCottagesResponse,
			options...,
		),
		updateLeadCottage: grpctransport.NewServer(
			endpoints.UpdateLeadCottageEndpoint,
			decodeUpdateLeadCottageRequest,
			encodeUpdateLeadCottagesResponse,
			options...,
		),
		deleteLeadCottage: grpctransport.NewServer(
			endpoints.DeleteLeadCottageEndpoint,
			decodeDeleteLeadCottageRequest,
			encodeDeleteLeadCottagesResponse,
			options...,
		),
		getLeadCottage: grpctransport.NewServer(
			endpoints.GetLeadCottageEndpoint,
			decodeGetLeadCottageRequest,
			encodeGetLeadCottagesResponse,
			options...,
		),
	}
}

func (s *grpcServer) CreateLeadCottage(ctx context.Context, request *apiv1.CreateLeadCottageRequest) (*apiv1.CreateLeadCottageResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.createLeadCottage)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.CreateLeadCottageResponse), nil
}

func decodeCreateLeadCottageRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.CreateLeadCottageRequest)

	return createLeadCottagesRequest{

		Lead: decodeLeadV1(req.LeadCottage),
	}, nil
}

func encodeCreateLeadCottageResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(creteLeadCottagesResponse)
	if resp.Err != nil {
		return &apiv1.CreateLeadCottageResponse{
			LeadId: 0,
		}, resp.Err
	}
	return &apiv1.CreateLeadCottageResponse{
		LeadId: resp.ID,
	}, nil
}

func (s *grpcServer) ListLeadCottages(ctx context.Context, request *apiv1.ListLeadCottagesRequest) (*apiv1.ListLeadCottagesResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.listLeadCottages)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.ListLeadCottagesResponse), nil
}

func decodeListLeadCottagesRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.ListLeadCottagesRequest)
	return listLeadCottagesRequest{
		Criteria: decodeLeadsSearchCriteria(req.Criteria),
	}, nil
}

func encodeListLeadCottagesResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(listLeadCottagesResponse)
	if resp.Err != nil {
		return &apiv1.ListLeadCottagesResponse{}, resp.Err
	}
	return &apiv1.ListLeadCottagesResponse{
		LeadCottages: encodeLeadsV1(resp.Leads),
		Total:        int64(resp.Total),
	}, nil
}

func (s *grpcServer) UpdateLeadCottage(ctx context.Context, request *apiv1.UpdateLeadCottageRequest) (*apiv1.UpdateLeadCottageResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.updateLeadCottage)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.UpdateLeadCottageResponse), nil
}

func decodeUpdateLeadCottageRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.UpdateLeadCottageRequest)
	return updateLeadCottageRequest{
		LeadID: req.LeadId,
		Lead:   decodeLeadV1(req.LeadCottage),
	}, nil
}

func encodeUpdateLeadCottagesResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(updateLeadCottageResponse)
	if resp.Err != nil {
		return &apiv1.UpdateLeadCottageResponse{}, resp.Err
	}
	return &apiv1.UpdateLeadCottageResponse{}, nil
}

func (s *grpcServer) DeleteLeadCottage(ctx context.Context, request *apiv1.DeleteLeadCottageRequest) (*apiv1.DeleteLeadCottageResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.deleteLeadCottage)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.DeleteLeadCottageResponse), nil
}

func decodeDeleteLeadCottageRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.DeleteLeadCottageRequest)
	return deleteLeadCottageRequest{
		LeadID: req.LeadId,
	}, nil
}

func encodeDeleteLeadCottagesResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(deleteLeadCottageResponse)
	if resp.Err != nil {
		return &apiv1.DeleteLeadCottageResponse{}, resp.Err
	}
	return &apiv1.DeleteLeadCottageResponse{}, nil
}

func (s *grpcServer) GetLeadCottage(ctx context.Context, request *apiv1.GetLeadCottageRequest) (*apiv1.GetLeadCottageResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.getLeadCottage)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.GetLeadCottageResponse), nil
}

func decodeGetLeadCottageRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.GetLeadCottageRequest)
	return getLeadCottageRequest{
		LeadID: req.LeadId,
	}, nil
}

func encodeGetLeadCottagesResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(getLeadCottageResponse)
	if resp.Err != nil {
		return &apiv1.GetLeadCottageResponse{}, resp.Err
	}
	return &apiv1.GetLeadCottageResponse{
		LeadCottage: encodeLeadV1(resp.Lead),
	}, nil
}
