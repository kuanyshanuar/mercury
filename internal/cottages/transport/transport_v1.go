package transport

import (
	"context"
	cottageserviceapi "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"gitlab.com/zharzhanov/mercury/internal/helpers"
)

type grpcServer struct {
	createCottage       grpctransport.Handler
	updateCottage       grpctransport.Handler
	deleteCottage       grpctransport.Handler
	getCottage          grpctransport.Handler
	listCottage         grpctransport.Handler
	listPopularCottages grpctransport.Handler
	listCottagesByID    grpctransport.Handler
	updateHousePlan     grpctransport.Handler
	deleteHousePlan     grpctransport.Handler
	createHousePlan     grpctransport.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer.
func NewGRPCServer(endpoints Endpoints, logger log.Logger) cottageserviceapi.CottageServiceServer {
	options := helpers.SetupServerOptions(logger)
	return &grpcServer{
		createCottage: grpctransport.NewServer(
			endpoints.CreateCottageEndpoint,
			decodeCreateCottageRequest,
			encodeCreateCottageResponse,
			options...,
		),
		updateCottage: grpctransport.NewServer(
			endpoints.UpdateCottageEndpoint,
			decodeUpdateCottageRequest,
			encodeUpdateCottageResponse,
			options...,
		),
		deleteCottage: grpctransport.NewServer(
			endpoints.DeleteCottageEndpoint,
			decodeDeleteCottageRequest,
			encodeDeleteCottageResponse,
			options...,
		),
		getCottage: grpctransport.NewServer(
			endpoints.GetCottageEndpoint,
			decodeGetCottageRequest,
			encodeGetCottageResponse,
			options...,
		),
		listCottage: grpctransport.NewServer(
			endpoints.ListCottageEndpoint,
			decodeListCottageRequest,
			encodeListCottageResponse,
			options...,
		),
		listPopularCottages: grpctransport.NewServer(
			endpoints.ListPopularCottagesEndpoint,
			decodeListPopularCottagesRequest,
			encodeListPopularCottagesResponse,
			options...,
		),
		listCottagesByID: grpctransport.NewServer(
			endpoints.ListCottagesByIDEndpoint,
			decodeListCottageByIDRequest,
			encodeListCottageByIDResponse,
			options...,
		),
		updateHousePlan: grpctransport.NewServer(
			endpoints.UpdateHousePlanEndpoint,
			decodeUpdateCottagePlanRequest,
			encodeUpdateCottagePlanResponse,
			options...,
		),
		deleteHousePlan: grpctransport.NewServer(
			endpoints.DeleteHousePlanEndpoint,
			decodeDeleteCottagePlanRequest,
			encodeDeleteCottagePlanResponse,
			options...,
		),
		createHousePlan: grpctransport.NewServer(
			endpoints.CreateHousePlanEndpoint,
			decodeCreateCottagePlanRequest,
			encodeCreateCottagePlanResponse,
			options...,
		),
	}

}

func (s *grpcServer) CreateCottage(ctx context.Context, request *cottageserviceapi.CreateCottageRequest) (*cottageserviceapi.CreateCottageResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.createCottage)
	if err != nil {
		return nil, err
	}
	return rep.(*cottageserviceapi.CreateCottageResponse), nil
}

func decodeCreateCottageRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*cottageserviceapi.CreateCottageRequest)
	return createCottageRequest{
		Cottage: decodeCottageWriteV1(req.Cottage),
	}, nil
}

func encodeCreateCottageResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(createCottageResponse)
	if resp.Error != nil {
		return &cottageserviceapi.CreateCottageResponse{}, resp.Error
	}

	return &cottageserviceapi.CreateCottageResponse{
		Cottage: encodeCottageReadV1(resp.Cottage),
	}, nil
}

func (s *grpcServer) UpdateCottage(
	ctx context.Context,
	request *cottageserviceapi.UpdateCottageRequest,
) (*cottageserviceapi.UpdateCottageResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.updateCottage)
	if err != nil {
		return nil, err
	}
	return rep.(*cottageserviceapi.UpdateCottageResponse), nil
}

func decodeUpdateCottageRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*cottageserviceapi.UpdateCottageRequest)
	return updateCottageRequest{
		ID:      req.Id,
		Cottage: decodeCottageWriteV1(req.Cottage),
	}, nil
}

func encodeUpdateCottageResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(updateCottageResponse)
	if resp.Error != nil {
		return &cottageserviceapi.UpdateCottageResponse{}, resp.Error
	}

	return &cottageserviceapi.UpdateCottageResponse{
		Cottage: encodeCottageReadV1(resp.Cottage),
	}, nil
}

func (s *grpcServer) DeleteCottage(ctx context.Context, request *cottageserviceapi.DeleteCottageRequest) (*cottageserviceapi.DeleteCottageResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.deleteCottage)
	if err != nil {
		return nil, err
	}
	return rep.(*cottageserviceapi.DeleteCottageResponse), nil
}

func decodeDeleteCottageRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*cottageserviceapi.DeleteCottageRequest)
	return deleteCottageRequest{
		ID: req.Id,
	}, nil
}

func encodeDeleteCottageResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(deleteCottageResponse)
	if resp.Error != nil {
		return &cottageserviceapi.DeleteCottageResponse{}, resp.Error
	}

	return &cottageserviceapi.DeleteCottageResponse{}, nil
}

func (s *grpcServer) ListCottages(ctx context.Context, request *cottageserviceapi.ListCottagesRequest) (*cottageserviceapi.ListCottagesResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.listCottage)
	if err != nil {
		return nil, err
	}

	return rep.(*cottageserviceapi.ListCottagesResponse), nil
}

func decodeListCottageRequest(
	_ context.Context,
	request interface{},
) (interface{}, error) {
	req := request.(*cottageserviceapi.ListCottagesRequest)

	return listCottageRequest{
		Criteria: decodeCriteriaV1(req.Criteria),
	}, nil
}

func encodeListCottageResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(listCottageResponse)
	if resp.Error != nil {
		return &cottageserviceapi.ListCottagesResponse{}, resp.Error
	}

	return &cottageserviceapi.ListCottagesResponse{
		Cottage:    encodeCottageListV1(resp.Cottage),
		TotalCount: int64(resp.Total),
	}, nil

}

func (s *grpcServer) GetCottage(ctx context.Context, request *cottageserviceapi.GetCottageRequest) (*cottageserviceapi.GetCottageResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.getCottage)
	if err != nil {
		return nil, err
	}

	return rep.(*cottageserviceapi.GetCottageResponse), nil
}

func decodeGetCottageRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*cottageserviceapi.GetCottageRequest)

	return getCottageRequest{
		ID: req.ID,
	}, nil
}

func encodeGetCottageResponse(_ context.Context, request interface{}) (interface{}, error) {
	resp := request.(getCottageResponse)
	if resp.Error != nil {
		return &cottageserviceapi.GetCottageResponse{}, resp.Error
	}

	return &cottageserviceapi.GetCottageResponse{
		Cottage: encodeCottageReadV1(resp.Cottage),
	}, nil
}

func (s *grpcServer) ListCottageByID(ctx context.Context, request *cottageserviceapi.ListCottageByIDRequest) (*cottageserviceapi.ListCottageByIDResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.listCottagesByID)
	if err != nil {
		return nil, err
	}
	return rep.(*cottageserviceapi.ListCottageByIDResponse), nil
}

func decodeListCottageByIDRequest(
	_ context.Context,
	request interface{},
) (interface{}, error) {
	req := request.(*cottageserviceapi.ListCottageByIDRequest)
	return listCottagesByIDRequest{
		IDs: req.Ids,
	}, nil
}

func encodeListCottageByIDResponse(_ context.Context, request interface{}) (interface{}, error) {
	resp := request.(listCottagesByIDResponse)
	if resp.Error != nil {
		return &cottageserviceapi.ListCottageByIDResponse{}, resp.Error
	}

	return &cottageserviceapi.ListCottageByIDResponse{
		Cottage: encodeCottageListV1(resp.Cottage),
	}, nil
}

func (s *grpcServer) ListPopularCottages(ctx context.Context, request *cottageserviceapi.ListPopularCottagesRequest) (*cottageserviceapi.ListPopularCottagesResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.listPopularCottages)
	if err != nil {
		return nil, err
	}
	return rep.(*cottageserviceapi.ListPopularCottagesResponse), nil
}

func decodeListPopularCottagesRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*cottageserviceapi.ListPopularCottagesRequest)
	return listPopularCottageRequest{
		Criteria: decodeCriteriaV1(req.Criteria),
	}, nil
}

func encodeListPopularCottagesResponse(_ context.Context, request interface{}) (interface{}, error) {
	resp := request.(listPopularCottageResponse)
	if resp.Error != nil {
		return &cottageserviceapi.ListPopularCottagesResponse{}, resp.Error
	}

	return &cottageserviceapi.ListPopularCottagesResponse{
		Cottage: encodeCottageListV1(resp.Cottage),
	}, nil
}

func (s *grpcServer) CreateCottagePlan(ctx context.Context, request *cottageserviceapi.CreateCottagePlanRequest) (*cottageserviceapi.CreateCottagePlanResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.createHousePlan)
	if err != nil {
		return nil, err
	}
	return rep.(*cottageserviceapi.CreateCottagePlanResponse), nil
}

func decodeCreateCottagePlanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*cottageserviceapi.CreateCottagePlanRequest)
	return createHousePlanRequest{
		HousePlan: decodeHousePlanV1(req.Plan),
	}, nil

}

func encodeCreateCottagePlanResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(createHousePlanResponse)
	if resp.Error != nil {
		return &cottageserviceapi.CreateCottagePlanResponse{}, resp.Error
	}

	return &cottageserviceapi.CreateCottagePlanResponse{
		Plan: encodeHousePlanV1(resp.CreatedPlan),
	}, nil
}

func (s *grpcServer) UpdateCottagePlan(ctx context.Context, request *cottageserviceapi.UpdateCottagePlanRequest) (*cottageserviceapi.UpdateCottagePlanResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.updateHousePlan)
	if err != nil {
		return nil, err
	}
	return rep.(*cottageserviceapi.UpdateCottagePlanResponse), nil
}

func decodeUpdateCottagePlanRequest(
	_ context.Context,
	request interface{},
) (interface{}, error) {
	req := request.(*cottageserviceapi.UpdateCottagePlanRequest)
	return updateHousePlanRequest{
		HousePlanID: req.PlanId,
		HousePlan:   decodeHousePlanV1(req.Plan),
	}, nil

}
func encodeUpdateCottagePlanResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(updateHousePlanResponse)
	if resp.Error != nil {
		return &cottageserviceapi.UpdateCottagePlanResponse{}, resp.Error
	}

	return &cottageserviceapi.UpdateCottagePlanResponse{
		Plan: encodeHousePlanV1(resp.HousePlan),
	}, nil

}

func decodeDeleteCottagePlanRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*cottageserviceapi.DeleteCottagePlanRequest)
	return deleteHousePlanRequest{
		HousePlanID: req.Plan,
	}, nil

}

func encodeDeleteCottagePlanResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(deleteHousePlanResponse)
	if resp.Error != nil {
		return &cottageserviceapi.DeleteCottagePlanResponse{}, resp.Error
	}

	return &cottageserviceapi.DeleteCottagePlanResponse{}, nil
}

func (s *grpcServer) DeleteCottagePlan(ctx context.Context, request *cottageserviceapi.DeleteCottagePlanRequest) (*cottageserviceapi.DeleteCottagePlanResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.deleteHousePlan)
	if err != nil {
		return nil, err
	}

	return rep.(*cottageserviceapi.DeleteCottagePlanResponse), nil
}
