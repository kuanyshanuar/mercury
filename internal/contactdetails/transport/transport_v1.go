package transport

import (
	"context"
	residenceapiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"
	"gitlab.com/zharzhanov/mercury/internal/helpers"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type grpcServer struct {
	createContactDetails          grpctransport.Handler
	createResidenceContactDetails grpctransport.Handler
	listContactDetails            grpctransport.Handler
	listResidenceContactDetails   grpctransport.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer.
func NewGRPCServer(endpoints Endpoints, logger log.Logger) residenceapiv1.ContactDetailsServiceServer {
	options := helpers.SetupServerOptions(logger)

	return &grpcServer{
		createContactDetails: grpctransport.NewServer(
			endpoints.CreateContactDetailsEndpoint,
			decodeCreateContactDetailsRequest,
			encodeCreateContactDetailsResponse,
			options...,
		),
		createResidenceContactDetails: grpctransport.NewServer(
			endpoints.CreateResidenceContactDetailsEndpoint,
			decodeCreateResidenceContactDetailsRequest,
			encodeCreateResidenceContactDetailsResponse,
			options...,
		),
		listContactDetails: grpctransport.NewServer(
			endpoints.ListContactDetailsEndpoint,
			decodeListContactDetailsRequest,
			encodeListContactDetailsResponse,
			options...,
		),
		listResidenceContactDetails: grpctransport.NewServer(
			endpoints.ListResidenceContactDetailsEndpoint,
			decodeListResidenceContactDetailsRequest,
			encodeListResidenceContactDetailsResponse,
			options...,
		),
	}
}

func (s *grpcServer) CreateContactDetails(ctx context.Context, request *residenceapiv1.CreateContactDetailsRequest) (*residenceapiv1.CreateContactDetailsResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.createContactDetails)
	if err != nil {
		return nil, err
	}
	return rep.(*residenceapiv1.CreateContactDetailsResponse), nil
}

func decodeCreateContactDetailsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*residenceapiv1.CreateContactDetailsRequest)

	return createContactDetailsRequest{
		ContactDetails: decodeContactDetails(req.ContactDetails),
	}, nil
}

func encodeCreateContactDetailsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(createContactDetailsResponse)
	if resp.Err != nil {
		return &residenceapiv1.CreateContactDetailsResponse{}, resp.Err
	}
	return &residenceapiv1.CreateContactDetailsResponse{}, nil
}

func (s *grpcServer) CreateResidenceContactDetails(ctx context.Context, request *residenceapiv1.CreateResidenceContactDetailsRequest) (*residenceapiv1.CreateResidenceContactDetailsResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.createResidenceContactDetails)
	if err != nil {
		return nil, err
	}
	return rep.(*residenceapiv1.CreateResidenceContactDetailsResponse), nil
}

func decodeCreateResidenceContactDetailsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*residenceapiv1.CreateResidenceContactDetailsRequest)

	return createResidenceContactDetailsRequest{
		ContactDetails: decodeResidenceContactDetails(req.ContactDetails),
	}, nil
}

func encodeCreateResidenceContactDetailsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(createResidenceContactDetailsResponse)
	if resp.Err != nil {
		return &residenceapiv1.CreateResidenceContactDetailsResponse{}, resp.Err
	}
	return &residenceapiv1.CreateResidenceContactDetailsResponse{}, nil
}

func (s *grpcServer) ListContactDetails(ctx context.Context, request *residenceapiv1.ListContactDetailsRequest) (*residenceapiv1.ListContactDetailsResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.listContactDetails)
	if err != nil {
		return nil, err
	}
	return rep.(*residenceapiv1.ListContactDetailsResponse), nil
}

func decodeListContactDetailsRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*residenceapiv1.ListContactDetailsRequest)
	_ = req

	return listContactDetailsRequest{}, nil
}

func encodeListContactDetailsResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(listContactDetailsResponse)
	if resp.Err != nil {
		return &residenceapiv1.ListContactDetailsResponse{}, resp.Err
	}
	return &residenceapiv1.ListContactDetailsResponse{
		ContactDetails: encodeListContactDetails(resp.ContactDetails),
	}, nil
}

func (s *grpcServer) ListResidenceContactDetails(ctx context.Context, request *residenceapiv1.ListResidenceContactDetailsRequest) (*residenceapiv1.ListResidenceContactDetailsResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.listResidenceContactDetails)
	if err != nil {
		return nil, err
	}
	return rep.(*residenceapiv1.ListResidenceContactDetailsResponse), nil
}

func decodeListResidenceContactDetailsRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*residenceapiv1.ListResidenceContactDetailsRequest)
	_ = req

	return listResidenceContactDetailsRequest{}, nil
}

func encodeListResidenceContactDetailsResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(listResidenceContactDetailsResponse)
	if resp.Err != nil {
		return &residenceapiv1.ListResidenceContactDetailsResponse{}, resp.Err
	}
	return &residenceapiv1.ListResidenceContactDetailsResponse{
		ContactDetails: encodeListResidenceContactDetails(resp.ContactDetails),
	}, nil
}
