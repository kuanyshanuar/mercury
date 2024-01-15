package transport

import (
	"context"

	residencesapiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type grpcServer struct {
	create         grpctransport.Handler
	list           grpctransport.Handler
	listPopular    grpctransport.Handler
	listByIDs      grpctransport.Handler
	get            grpctransport.Handler
	update         grpctransport.Handler
	delete         grpctransport.Handler
	createFlatPlan grpctransport.Handler
	updateFlatPlan grpctransport.Handler
	deleteFlatPlan grpctransport.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer.
func NewGRPCServer(endpoints Endpoints, logger log.Logger) residencesapiv1.ResidenceServiceServer {
	options := helpers.SetupServerOptions(logger)

	return &grpcServer{
		create: grpctransport.NewServer(
			endpoints.CreateResidenceEndpoint,
			decodeCreateResidenceRequestV1,
			encodeCreateResidenceResponseV1,
			options...,
		),
		list: grpctransport.NewServer(
			endpoints.ListResidencesEndpoint,
			decodeListResidenceRequestV1,
			encodeListResidenceResponseV1,
			options...,
		),
		listPopular: grpctransport.NewServer(
			endpoints.ListPopularResidencesEndpoint,
			decodeListPopularResidenceRequestV1,
			encodeListPopularResidenceResponseV1,
			options...,
		),
		listByIDs: grpctransport.NewServer(
			endpoints.ListResidencesByIDsEndpoint,
			decodeListResidenceByIDsRequestV1,
			encodeListResidenceByIDsResponseV1,
			options...,
		),
		get: grpctransport.NewServer(
			endpoints.GetResidenceEndpoint,
			decodeGetResidenceRequestV1,
			encodeGetResidenceResponseV1,
			options...,
		),
		update: grpctransport.NewServer(
			endpoints.UpdateResidenceEndpoint,
			decodeUpdateResidenceRequestV1,
			encodeUpdateResidenceResponseV1,
			options...,
		),
		delete: grpctransport.NewServer(
			endpoints.DeleteResidenceEndpoint,
			decodeDeleteResidenceRequestV1,
			encodeDeleteResidenceResponseV1,
			options...,
		),
		createFlatPlan: grpctransport.NewServer(
			endpoints.CreateFlatPlanEndpoint,
			decodeCreateFlatPlanRequestV1,
			encodeCreateFlatPlanResponseV1,
			options...,
		),
		updateFlatPlan: grpctransport.NewServer(
			endpoints.UpdateFlatPlanEndpoint,
			decodeUpdateFlatPlanRequestV1,
			encodeUpdateFlatPlanResponseV1,
			options...,
		),
		deleteFlatPlan: grpctransport.NewServer(
			endpoints.DeleteFlatPlanEndpoint,
			decodeDeleteFlatPlanRequestV1,
			encodeDeleteFlatPlanResponseV1,
			options...,
		),
	}
}

func decodeCreateResidenceRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*residencesapiv1.CreateResidenceRequest)

	return createResidenceRequest{
		Residence: decodeResidenceV1(req.Residence),
	}, nil
}

func encodeCreateResidenceResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(createResidenceResponse)
	if resp.Err != nil {
		return &residencesapiv1.CreateResidenceResponse{}, resp.Err
	}
	return &residencesapiv1.CreateResidenceResponse{
		Residence: encodeResidenceV1(resp.Residence),
	}, nil
}

func (s *grpcServer) CreateResidence(ctx context.Context, request *residencesapiv1.CreateResidenceRequest) (*residencesapiv1.CreateResidenceResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.create)
	if err != nil {
		return nil, err
	}
	return rep.(*residencesapiv1.CreateResidenceResponse), nil
}

func decodeListResidenceRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*residencesapiv1.ListResidencesRequest)

	return listResidencesRequest{
		Criteria: decodeResidenceSearchCriteria(req.Criteria),
	}, nil
}

func encodeListResidenceResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(listResidencesResponse)
	if resp.Err != nil {
		return &residencesapiv1.ListResidencesResponse{}, resp.Err
	}
	return &residencesapiv1.ListResidencesResponse{
		Residences: encodeListResidencesV1(resp.Residences),
		Total:      int64(resp.Total),
	}, nil
}

func (s *grpcServer) ListResidences(ctx context.Context, request *residencesapiv1.ListResidencesRequest) (*residencesapiv1.ListResidencesResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.list)
	if err != nil {
		return nil, err
	}
	return rep.(*residencesapiv1.ListResidencesResponse), nil
}

func decodeGetResidenceRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*residencesapiv1.GetResidenceRequest)

	return getResidenceRequest{
		ResidenceID: domain.ResidenceID(req.Id),
	}, nil
}

func encodeGetResidenceResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(getResidenceResponse)
	if resp.Err != nil {
		return &residencesapiv1.GetResidenceResponse{}, resp.Err
	}
	return &residencesapiv1.GetResidenceResponse{
		Residence: encodeResidenceV1(resp.Residence),
	}, nil
}

func (s *grpcServer) GetResidence(ctx context.Context, request *residencesapiv1.GetResidenceRequest) (*residencesapiv1.GetResidenceResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.get)
	if err != nil {
		return nil, err
	}
	return rep.(*residencesapiv1.GetResidenceResponse), nil
}

func decodeUpdateResidenceRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*residencesapiv1.UpdateResidenceRequest)
	return updateResidenceRequest{
		ResidenceID: domain.ResidenceID(req.Id),
		Residence:   decodeResidenceV1(req.Residence),
	}, nil
}

func encodeUpdateResidenceResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(updateResidenceResponse)
	if resp.Err != nil {
		return &residencesapiv1.UpdateResidenceResponse{}, resp.Err
	}
	return &residencesapiv1.UpdateResidenceResponse{
		Residence: encodeResidenceV1(resp.Residence),
	}, nil
}

func (s *grpcServer) UpdateResidence(ctx context.Context, request *residencesapiv1.UpdateResidenceRequest) (*residencesapiv1.UpdateResidenceResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.update)
	if err != nil {
		return nil, err
	}
	return rep.(*residencesapiv1.UpdateResidenceResponse), nil
}

func decodeDeleteResidenceRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*residencesapiv1.DeleteResidenceRequest)
	return deleteResidenceRequest{
		ResidenceID: domain.ResidenceID(req.Id),
	}, nil
}

func encodeDeleteResidenceResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(deleteResidenceResponse)
	if resp.Err != nil {
		return &residencesapiv1.DeleteResidenceResponse{}, resp.Err
	}
	return &residencesapiv1.DeleteResidenceResponse{}, nil
}

func (s *grpcServer) DeleteResidence(ctx context.Context, request *residencesapiv1.DeleteResidenceRequest) (*residencesapiv1.DeleteResidenceResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.delete)
	if err != nil {
		return nil, err
	}
	return rep.(*residencesapiv1.DeleteResidenceResponse), nil
}

func (s *grpcServer) ListResidencesByIDs(ctx context.Context, request *residencesapiv1.ListResidencesByIDsRequest) (*residencesapiv1.ListResidencesByIDsResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.listByIDs)
	if err != nil {
		return nil, err
	}
	return rep.(*residencesapiv1.ListResidencesByIDsResponse), nil
}

func decodeListResidenceByIDsRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*residencesapiv1.ListResidencesByIDsRequest)
	return listResidencesByIDsRequest{
		ResidenceIDs: decodeResidenceIDs(req.ResidenceIds),
	}, nil
}

func encodeListResidenceByIDsResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(listResidencesByIDsResponse)
	if resp.Err != nil {
		return &residencesapiv1.ListResidencesByIDsResponse{}, resp.Err
	}
	return &residencesapiv1.ListResidencesByIDsResponse{
		Residences: encodeListResidencesV1(resp.Residences),
	}, nil
}

func (s *grpcServer) ListPopularResidences(ctx context.Context, request *residencesapiv1.ListPopularResidencesRequest) (*residencesapiv1.ListPopularResidencesResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.listPopular)
	if err != nil {
		return nil, err
	}
	return rep.(*residencesapiv1.ListPopularResidencesResponse), nil
}

func decodeListPopularResidenceRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*residencesapiv1.ListPopularResidencesRequest)
	return listPopularResidencesRequest{
		Criteria: decodeResidenceSearchCriteria(req.Criteria),
	}, nil
}

func encodeListPopularResidenceResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(listPopularResidencesResponse)
	if resp.Err != nil {
		return &residencesapiv1.ListPopularResidencesResponse{}, resp.Err
	}
	return &residencesapiv1.ListPopularResidencesResponse{
		Residences: encodeListResidencesV1(resp.Residences),
		Total:      int64(resp.Total),
	}, nil
}

func (s *grpcServer) CreateFlatPlan(ctx context.Context, request *residencesapiv1.CreateFlatPlanRequest) (*residencesapiv1.CreateFlatPlanResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.createFlatPlan)
	if err != nil {
		return nil, err
	}
	return rep.(*residencesapiv1.CreateFlatPlanResponse), nil
}

func decodeCreateFlatPlanRequestV1(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*residencesapiv1.CreateFlatPlanRequest)

	return createFlatPlanRequest{
		FlatPlan: decodeFlatPlanV1(req.GetFlatPlan()),
	}, nil
}

func encodeCreateFlatPlanResponseV1(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(createFlatPlanResponse)
	if resp.Err != nil {
		return &residencesapiv1.CreateFlatPlanResponse{}, resp.Err
	}
	return &residencesapiv1.CreateFlatPlanResponse{
		FlatPlan: encodeFlatPlanV1(resp.FlatPlan),
	}, nil
}

func (s *grpcServer) UpdateFlatPlan(ctx context.Context, request *residencesapiv1.UpdateFlatPlanRequest) (*residencesapiv1.UpdateFlatPlanResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.updateFlatPlan)
	if err != nil {
		return nil, err
	}
	return rep.(*residencesapiv1.UpdateFlatPlanResponse), nil
}

func decodeUpdateFlatPlanRequestV1(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*residencesapiv1.UpdateFlatPlanRequest)

	return updateFlatPlanRequest{
		ID:       domain.FlatPlanID(req.Id),
		FlatPlan: decodeFlatPlanV1(req.FlatPlan),
	}, nil
}

func encodeUpdateFlatPlanResponseV1(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(updateFlatPlanResponse)

	if resp.Err != nil {
		return &residencesapiv1.UpdateFlatPlanResponse{}, resp.Err
	}
	return &residencesapiv1.UpdateFlatPlanResponse{
		FlatPlan: encodeFlatPlanV1(resp.FlatPlan),
	}, nil
}

func (s *grpcServer) DeleteFlatPlan(ctx context.Context, request *residencesapiv1.DeleteFlatPlanRequest) (*residencesapiv1.DeleteFlatPlanResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.deleteFlatPlan)
	if err != nil {
		return nil, err
	}
	return rep.(*residencesapiv1.DeleteFlatPlanResponse), nil
}

func decodeDeleteFlatPlanRequestV1(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*residencesapiv1.DeleteFlatPlanRequest)

	return deleteFlatPlanRequest{
		ID: domain.FlatPlanID(req.Id),
	}, nil
}

func encodeDeleteFlatPlanResponseV1(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(deleteFlatPlanResponse)
	if resp.Err != nil {
		return &residencesapiv1.DeleteFlatPlanResponse{}, resp.Err
	}
	return &residencesapiv1.DeleteFlatPlanResponse{}, nil
}
