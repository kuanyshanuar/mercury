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
func NewGRPCServer(endpoints Endpoints, logger log.Logger) apiv1.ManagerServiceServer {
	options := helpers.SetupServerOptions(logger)

	return &grpcServer{
		create: grpctransport.NewServer(
			endpoints.CreateManagerEndpoint,
			decodeCreateManagerRequestV1,
			encodeCreateManagerResponseV1,
			options...,
		),
		list: grpctransport.NewServer(
			endpoints.ListManagersEndpoint,
			decodeListManagersRequestV1,
			encodeListManagersResponseV1,
			options...,
		),
		get: grpctransport.NewServer(
			endpoints.GetManagerEndpoint,
			decodeGetManagerRequestV1,
			encodeGetManagerResponseV1,
			options...,
		),
		update: grpctransport.NewServer(
			endpoints.UpdateManagerEndpoint,
			decodeUpdateManagerRequestV1,
			encodeUpdateManagerResponseV1,
			options...,
		),
		delete: grpctransport.NewServer(
			endpoints.DeleteManagerEndpoint,
			decodeDeleteManagerRequestV1,
			encodeDeleteManagerResponseV1,
			options...,
		),
	}
}

// CreateManager implements v1.ManagerServiceServer
func (s *grpcServer) CreateManager(ctx context.Context, request *apiv1.CreateManagerRequest) (*apiv1.CreateManagerResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.create)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.CreateManagerResponse), nil
}

func decodeCreateManagerRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.CreateManagerRequest)

	return createManagerRequest{
		Manager: decodeManagerV1(req.Manager),
	}, nil
}

func encodeCreateManagerResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(createManagerResponse)
	if resp.Err != nil {
		return &apiv1.CreateManagerResponse{}, resp.Err
	}
	return &apiv1.CreateManagerResponse{}, nil
}

// GetManager implements v1.ManagerServiceServer
func (s *grpcServer) GetManager(ctx context.Context, request *apiv1.GetManagerRequest) (*apiv1.GetManagerResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.get)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.GetManagerResponse), nil
}

func encodeGetManagerResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(getManagerResponse)
	if resp.Err != nil {
		return &apiv1.GetManagerResponse{}, resp.Err
	}

	return &apiv1.GetManagerResponse{
		Manager: encodeManagerV1(resp.Manager),
	}, nil
}

func decodeGetManagerRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.GetManagerRequest)

	return getManagerRequest{
		ManagerID: domain.ManagerID(req.Id),
	}, nil
}

// ListManagers implements v1.ManagerServiceServer
func (s *grpcServer) ListManagers(ctx context.Context, request *apiv1.ListManagersRequest) (*apiv1.ListManagersResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.list)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.ListManagersResponse), nil
}

func encodeListManagersResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(listManagerResponse)
	if resp.Err != nil {
		return &apiv1.ListManagersResponse{}, resp.Err
	}

	return &apiv1.ListManagersResponse{
		Managers: encodeManagersV1(resp.Managers),
		Total:    int64(resp.Total),
	}, nil
}

func decodeListManagersRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.ListManagersRequest)

	return listManagerRequest{
		Criteria: decodeManagerSearchCriteriaV1(req.Criteria),
	}, nil
}

// UpdateManager implements v1.ManagerServiceServer
func (s *grpcServer) UpdateManager(ctx context.Context, request *apiv1.UpdateManagerRequest) (*apiv1.UpdateManagerResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.update)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.UpdateManagerResponse), nil
}

func encodeUpdateManagerResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(updateManagerResponse)
	if resp.Err != nil {
		return &apiv1.UpdateManagerResponse{}, resp.Err
	}

	return &apiv1.UpdateManagerResponse{}, nil
}

func decodeUpdateManagerRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.UpdateManagerRequest)

	return updateManagerRequest{
		ManagerID: domain.ManagerID(req.Id),
		Manager:   decodeManagerV1(req.Manager),
	}, nil
}

// DeleteManager implements v1.ManagerServiceServer
func (s *grpcServer) DeleteManager(ctx context.Context, request *apiv1.DeleteManagerRequest) (*apiv1.DeleteManagerResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.delete)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.DeleteManagerResponse), nil
}

func decodeDeleteManagerRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.DeleteManagerRequest)

	return deleteManagerRequest{
		ManagerID: domain.ManagerID(req.Id),
	}, nil
}

func encodeDeleteManagerResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(deleteManagerResponse)
	if resp.Err != nil {
		return &apiv1.DeleteManagerResponse{}, resp.Err
	}

	return &apiv1.DeleteManagerResponse{}, nil
}
