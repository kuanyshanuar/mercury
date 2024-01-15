package transport

import (
	"context"
	apiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"
	"gitlab.com/zharzhanov/mercury/internal/helpers"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type grpcServer struct {
	subscribeEndpoint    grpctransport.Handler
	unsubscribeEndpoint  grpctransport.Handler
	listBuildersEndpoint grpctransport.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer.
func NewGRPCServer(endpoints Endpoints, logger log.Logger) apiv1.SubscribersServiceServer {
	options := helpers.SetupServerOptions(logger)

	return &grpcServer{
		subscribeEndpoint: grpctransport.NewServer(
			endpoints.SubscribeEndpoint,
			decodeSubscribeRequest,
			encodeSubscribeResponse,
			options...,
		),
		unsubscribeEndpoint: grpctransport.NewServer(
			endpoints.UnsubscribeEndpoint,
			decodeUnsubscribeRequest,
			encodeUnsubscribeResponse,
			options...,
		),
		listBuildersEndpoint: grpctransport.NewServer(
			endpoints.ListBuildersEndpoint,
			decodeListBuildersRequest,
			encodeListBuildersResponse,
			options...,
		),
	}
}

func (s *grpcServer) Subscribe(ctx context.Context, request *apiv1.SubscribeRequest) (*apiv1.SubscribeResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.subscribeEndpoint)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.SubscribeResponse), nil
}

func decodeSubscribeRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.SubscribeRequest)

	return subscribeRequest{
		SubscriberID: req.UserId,
		BuilderID:    req.BuilderId,
	}, nil
}

func encodeSubscribeResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(subscribeResponse)

	if resp.Err != nil {
		return &apiv1.SubscribeResponse{}, resp.Err
	}
	return &apiv1.SubscribeResponse{}, nil
}

func (s *grpcServer) Unsubscribe(ctx context.Context, request *apiv1.UnsubscribeRequest) (*apiv1.UnsubscribeResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.unsubscribeEndpoint)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.UnsubscribeResponse), nil
}

func decodeUnsubscribeRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.UnsubscribeRequest)

	return unsubscribeRequest{
		SubscriberID: req.UserId,
		BuilderID:    req.BuilderId,
	}, nil
}

func encodeUnsubscribeResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(unsubscribeResponse)

	if resp.Err != nil {
		return &apiv1.UnsubscribeResponse{}, resp.Err
	}
	return &apiv1.UnsubscribeResponse{}, nil
}

func (s *grpcServer) ListSubscribedBuilders(ctx context.Context, request *apiv1.ListSubscribedBuildersRequest) (*apiv1.ListSubscribedBuildersResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.listBuildersEndpoint)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.ListSubscribedBuildersResponse), nil
}

func decodeListBuildersRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.ListSubscribedBuildersRequest)

	return listBuildersRequest{
		SubscriberID: req.UserId,
		Criteria:     encodeSearchCriteria(req.Criteria),
	}, nil
}

func encodeListBuildersResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(listBuildersResponse)

	if resp.Err != nil {
		return &apiv1.ListSubscribedBuildersResponse{}, resp.Err
	}
	return &apiv1.ListSubscribedBuildersResponse{
		Builders: encodeBuildersV1(resp.Builders),
		Total:    int64(resp.Total),
	}, nil
}
