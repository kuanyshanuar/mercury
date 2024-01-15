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
	createCity         grpctransport.Handler
	listCities         grpctransport.Handler
	updateCity         grpctransport.Handler
	deleteCity         grpctransport.Handler
	listDistricts      grpctransport.Handler
	listFilters        grpctransport.Handler
	listFiltersV2      grpctransport.Handler
	listBuilderFilters grpctransport.Handler
	createFilter       grpctransport.Handler
	deleteFilter       grpctransport.Handler
	createDistrict     grpctransport.Handler
	updateDistrict     grpctransport.Handler
	deleteDistrict     grpctransport.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer.
func NewGRPCServer(endpoints Endpoints, logger log.Logger) apiv1.ResidenceFilterServiceServer {
	options := helpers.SetupServerOptions(logger)

	return &grpcServer{
		createCity: grpctransport.NewServer(
			endpoints.CreateCityEndpoint,
			decodeCreateCityRequest,
			encodeCreateCityResponse,
			options...,
		),
		listCities: grpctransport.NewServer(
			endpoints.ListCitiesEndpoint,
			decodeListCitiesRequest,
			encodeListCitiesResponse,
			options...,
		),
		updateCity: grpctransport.NewServer(
			endpoints.UpdateCityEndpoint,
			decodeUpdateCityRequest,
			encodeUpdateCityResponse,
			options...,
		),
		deleteCity: grpctransport.NewServer(
			endpoints.DeleteCityEndpoint,
			decodeDeleteCityRequest,
			encodeDeleteCityResponse,
			options...,
		),
		listFilters: grpctransport.NewServer(
			endpoints.ListFiltersEndpoint,
			decodeListFiltersRequest,
			encodeListFiltersResponse,
			options...,
		),
		listFiltersV2: grpctransport.NewServer(
			endpoints.ListFiltersEndpoint,
			decodeListFiltersRequestV2,
			encodeListFiltersResponseV2,
			options...,
		),
		listDistricts: grpctransport.NewServer(
			endpoints.ListDistrictsEndpoint,
			decodeListDistrictsRequest,
			encodeListDistrictsResponse,
			options...,
		),
		listBuilderFilters: grpctransport.NewServer(
			endpoints.ListFilterBuildersEndpoint,
			decodeListBuildersRequest,
			encodeListBuildersResponse,
			options...,
		),
		createFilter: grpctransport.NewServer(
			endpoints.CreateFilterEndpoint,
			decodeCreateFilterRequest,
			encodeCreateFilterResponse,
			options...,
		),
		deleteFilter: grpctransport.NewServer(
			endpoints.DeleteFilterEndpoint,
			decodeDeleteFilterRequest,
			encodeDeleteFilterResponse,
			options...,
		),
		createDistrict: grpctransport.NewServer(
			endpoints.CreateDistrictEndpoint,
			decodeCreateDistrictRequest,
			encodeCreateDistrictResponse,
			options...,
		),
		updateDistrict: grpctransport.NewServer(
			endpoints.UpdateDistrictEndpoint,
			decodeUpdateDistrictRequest,
			encodeUpdateDistrictResponse,
			options...,
		),
		deleteDistrict: grpctransport.NewServer(
			endpoints.DeleteDistrictEndpoint,
			decodeDeleteDistrictRequest,
			encodeDeleteDistrictResponse,
			options...,
		),
	}
}

func (s *grpcServer) ListCities(ctx context.Context, request *apiv1.ListCitiesRequest) (*apiv1.ListCitiesResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.listCities)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.ListCitiesResponse), nil
}

func decodeListCitiesRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.ListCitiesRequest)
	_ = req

	return listCitiesRequest{
		Criteria: decodeCitySearchCriteria(req.Criteria),
	}, nil
}

func encodeListCitiesResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(listCitiesResponse)
	if resp.Err != nil {
		return &apiv1.ListCitiesResponse{}, resp.Err
	}
	return &apiv1.ListCitiesResponse{
		Cities: encodeCities(resp.Cities),
		Total:  int64(resp.Total),
	}, nil
}

func (s *grpcServer) ListDistricts(ctx context.Context, request *apiv1.ListDistrictsRequest) (*apiv1.ListDistrictsResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.listDistricts)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.ListDistrictsResponse), nil
}

func decodeListDistrictsRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.ListDistrictsRequest)

	return listDistrictRequest{
		Criteria: decodeDistrictSearchCriteria(req.Criteria),
	}, nil
}

func encodeListDistrictsResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(listDistrictResponse)
	if resp.Err != nil {
		return &apiv1.ListDistrictsResponse{}, resp.Err
	}
	return &apiv1.ListDistrictsResponse{
		Districts: encodeDistricts(resp.Districts),
		Total:     int64(resp.Total),
	}, nil
}

func (s *grpcServer) ListFilters(ctx context.Context, request *apiv1.ListFiltersRequest) (*apiv1.ListFiltersResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.listFilters)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.ListFiltersResponse), nil
}

func decodeListFiltersRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.ListFiltersRequest)
	_ = req

	return listFiltersRequest{}, nil
}

func encodeListFiltersResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(listFiltersResponse)
	if resp.Err != nil {
		return &apiv1.ListFiltersResponse{}, resp.Err
	}
	return &apiv1.ListFiltersResponse{
		Filters: encodeFilters(resp.Filter),
	}, nil
}

func (s *grpcServer) ListFiltersV2(ctx context.Context, request *apiv1.ListFiltersRequestV2) (*apiv1.ListFiltersResponseV2, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.listFiltersV2)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.ListFiltersResponseV2), nil
}

func decodeListFiltersRequestV2(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.ListFiltersRequestV2)
	_ = req

	return listFiltersRequest{}, nil
}

func encodeListFiltersResponseV2(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(listFiltersResponse)
	if resp.Err != nil {
		return &apiv1.ListFiltersResponseV2{}, resp.Err
	}
	return &apiv1.ListFiltersResponseV2{
		Filters: encodeFiltersV2(resp.Filter),
	}, nil
}

func (s *grpcServer) ListFilterBuilders(ctx context.Context, request *apiv1.ListFilterBuildersRequest) (*apiv1.ListFilterBuildersResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.listBuilderFilters)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.ListFilterBuildersResponse), nil
}

func decodeListBuildersRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.ListFilterBuildersRequest)
	_ = req

	return listFilterBuildersRequest{}, nil
}

func encodeListBuildersResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(listFilterBuildersResponse)
	if resp.Err != nil {
		return &apiv1.ListFilterBuildersResponse{}, resp.Err
	}
	return &apiv1.ListFilterBuildersResponse{
		Builders: encodeFilterBuilders(resp.Builders),
	}, nil
}

func (s *grpcServer) CreateFilter(ctx context.Context, request *apiv1.CreateFilterRequest) (*apiv1.CreateFilterResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.createFilter)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.CreateFilterResponse), nil
}

func decodeCreateFilterRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.CreateFilterRequest)
	_ = req

	return createFilterRequest{
		Key:    req.Key,
		Filter: decodeFilterV1(req.Filter),
	}, nil
}

func encodeCreateFilterResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(createFilterResponse)
	if resp.Err != nil {
		return &apiv1.CreateFilterResponse{}, resp.Err
	}
	return &apiv1.CreateFilterResponse{
		Filter: encodeFilterV1(resp.Filter),
	}, nil
}

func (s *grpcServer) DeleteFilter(ctx context.Context, request *apiv1.DeleteFilterRequest) (*apiv1.DeleteFilterResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.deleteFilter)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.DeleteFilterResponse), nil
}

func decodeDeleteFilterRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.DeleteFilterRequest)
	_ = req

	return deleteFilterRequest{
		ID:  req.Id,
		Key: req.Key,
	}, nil
}

func encodeDeleteFilterResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(deleteFilterResponse)
	if resp.Err != nil {
		return &apiv1.DeleteFilterResponse{}, resp.Err
	}
	return &apiv1.DeleteFilterResponse{}, nil
}

func (s *grpcServer) CreateCity(ctx context.Context, request *apiv1.CreateCityRequest) (*apiv1.CreateCityResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.createCity)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.CreateCityResponse), nil
}

func decodeCreateCityRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.CreateCityRequest)
	_ = req

	return createCityRequest{
		City: decodeCityV1(req.City),
	}, nil
}

func encodeCreateCityResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(createCityResponse)
	if resp.Err != nil {
		return &apiv1.CreateCityResponse{}, resp.Err
	}
	return &apiv1.CreateCityResponse{
		Id: int64(resp.CityID),
	}, nil
}

func (s *grpcServer) UpdateCity(ctx context.Context, request *apiv1.UpdateCityRequest) (*apiv1.UpdateCityResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.updateCity)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.UpdateCityResponse), nil
}

func decodeUpdateCityRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.UpdateCityRequest)
	_ = req

	return updateCityRequest{
		CityID: domain.CityID(req.Id),
		City:   decodeCityV1(req.City),
	}, nil
}

func encodeUpdateCityResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(updateCityResponse)
	if resp.Err != nil {
		return &apiv1.UpdateCityResponse{}, resp.Err
	}
	return &apiv1.UpdateCityResponse{}, nil
}

func (s *grpcServer) DeleteCity(ctx context.Context, request *apiv1.DeleteCityRequest) (*apiv1.DeleteCityResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.deleteCity)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.DeleteCityResponse), nil
}

func decodeDeleteCityRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.DeleteCityRequest)
	_ = req

	return deleteCityRequest{
		CityID: domain.CityID(req.Id),
	}, nil
}

func encodeDeleteCityResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(deleteCityResponse)
	if resp.Err != nil {
		return &apiv1.DeleteCityResponse{}, resp.Err
	}
	return &apiv1.DeleteCityResponse{}, nil
}

func (s *grpcServer) CreateDistrict(ctx context.Context, request *apiv1.CreateDistrictRequest) (*apiv1.CreateDistrictResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.createDistrict)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.CreateDistrictResponse), nil
}

func decodeCreateDistrictRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.CreateDistrictRequest)
	_ = req

	return createDistrictRequest{
		District: decodeDistrictV1(req.District),
	}, nil
}

func encodeCreateDistrictResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(createDistrictResponse)
	if resp.Err != nil {
		return &apiv1.CreateDistrictResponse{}, resp.Err
	}
	return &apiv1.CreateDistrictResponse{
		Id: int64(resp.DistrictID),
	}, nil
}

func (s *grpcServer) UpdateDistrict(ctx context.Context, request *apiv1.UpdateDistrictRequest) (*apiv1.UpdateDistrictResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.updateDistrict)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.UpdateDistrictResponse), nil
}

func decodeUpdateDistrictRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.UpdateDistrictRequest)
	_ = req

	return updateDistrictRequest{
		DistrictID: domain.DistrictID(req.DistrictId),
		District:   decodeDistrictV1(req.District),
	}, nil
}

func encodeUpdateDistrictResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(updateDistrictResponse)
	if resp.Err != nil {
		return &apiv1.UpdateDistrictResponse{}, resp.Err
	}
	return &apiv1.UpdateDistrictResponse{}, nil
}

func (s *grpcServer) DeleteDistrict(ctx context.Context, request *apiv1.DeleteDistrictRequest) (*apiv1.DeleteDistrictResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.deleteDistrict)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.DeleteDistrictResponse), nil
}

func decodeDeleteDistrictRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.DeleteDistrictRequest)
	_ = req

	return deleteDistrictRequest{
		DistrictID: domain.DistrictID(req.DistrictId),
	}, nil
}

func encodeDeleteDistrictResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(deleteDistrictResponse)
	if resp.Err != nil {
		return &apiv1.DeleteDistrictResponse{}, resp.Err
	}
	return &apiv1.DeleteDistrictResponse{}, nil
}
