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
	createLeadResidence grpctransport.Handler
	listLeadResidences  grpctransport.Handler
	updateLeadResidence grpctransport.Handler
	deleteLeadResidence grpctransport.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer.
func NewGRPCServer(endpoints Endpoints, logger log.Logger) apiv1.LeadResidencesServiceServer {
	options := helpers.SetupServerOptions(logger)

	return &grpcServer{
		createLeadResidence: grpctransport.NewServer(
			endpoints.CreateLeadResidenceEndpoint,
			decodeCreateLeadResidenceRequest,
			encodeCreateLeadResidenceResponse,
			options...,
		),
		listLeadResidences: grpctransport.NewServer(
			endpoints.ListLeadResidencesEndpoint,
			decodeListLeadResidencesRequest,
			encodeListLeadResidencesResponse,
			options...,
		),
		updateLeadResidence: grpctransport.NewServer(
			endpoints.UpdateLeadResidenceEndpoint,
			decodeUpdateLeadResidenceRequest,
			encodeUpdateLeadResidencesResponse,
			options...,
		),
		deleteLeadResidence: grpctransport.NewServer(
			endpoints.DeleteLeadResidenceEndpoint,
			decodeDeleteLeadResidenceRequest,
			encodeDeleteLeadResidencesResponse,
			options...,
		),
	}
}

func (s *grpcServer) CreateLeadResidence(ctx context.Context, request *apiv1.CreateLeadResidenceRequest) (*apiv1.CreateLeadResidenceResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.createLeadResidence)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.CreateLeadResidenceResponse), nil
}

func decodeCreateLeadResidenceRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.CreateLeadResidenceRequest)

	return creteLeadResidencesRequest{
		Lead: decodeLeadV1(req.Lead),
	}, nil
}

func encodeCreateLeadResidenceResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(creteLeadResidencesResponse)
	if resp.Err != nil {
		return &apiv1.CreateLeadResidenceResponse{}, resp.Err
	}
	return &apiv1.CreateLeadResidenceResponse{}, nil
}

func (s *grpcServer) ListLeadResidences(ctx context.Context, request *apiv1.ListLeadResidencesRequest) (*apiv1.ListLeadResidencesResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.listLeadResidences)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.ListLeadResidencesResponse), nil
}

func decodeListLeadResidencesRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.ListLeadResidencesRequest)
	return listLeadResidencesRequest{
		Criteria: decodeLeadsSearchCriteria(req.Criteria),
	}, nil
}

func encodeListLeadResidencesResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(listLeadResidencesResponse)
	if resp.Err != nil {
		return &apiv1.ListLeadResidencesResponse{}, resp.Err
	}
	return &apiv1.ListLeadResidencesResponse{
		Leads: encodeLeadsV1(resp.Leads),
		Total: int64(resp.Total),
	}, nil
}

func (s *grpcServer) UpdateLeadResidence(ctx context.Context, request *apiv1.UpdateLeadResidenceRequest) (*apiv1.UpdateLeadResidenceResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.updateLeadResidence)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.UpdateLeadResidenceResponse), nil
}

func decodeUpdateLeadResidenceRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.UpdateLeadResidenceRequest)
	return updateLeadResidenceRequest{
		LeadID: domain.LeadID(req.LeadId),
		Lead:   decodeLeadV1(req.Lead),
	}, nil
}

func encodeUpdateLeadResidencesResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(updateLeadResidenceResponse)
	if resp.Err != nil {
		return &apiv1.UpdateLeadResidenceResponse{}, resp.Err
	}
	return &apiv1.UpdateLeadResidenceResponse{}, nil
}

func (s *grpcServer) DeleteLeadResidence(ctx context.Context, request *apiv1.DeleteLeadResidenceRequest) (*apiv1.DeleteLeadResidenceResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.deleteLeadResidence)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.DeleteLeadResidenceResponse), nil
}

func decodeDeleteLeadResidenceRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.DeleteLeadResidenceRequest)
	return deleteLeadResidenceRequest{
		LeadID: domain.LeadID(req.LeadId),
	}, nil
}

func encodeDeleteLeadResidencesResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(deleteLeadResidenceResponse)
	if resp.Err != nil {
		return &apiv1.DeleteLeadResidenceResponse{}, resp.Err
	}
	return &apiv1.DeleteLeadResidenceResponse{}, nil
}
