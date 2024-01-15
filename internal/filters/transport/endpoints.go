package transport

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

const (
	serviceName = "Filters"
)

// Endpoints collects all of the endpoints that compose an add service. It's meant to
// be used as a helper struct, to collect all of the endpoints into a single
// parameter.
type Endpoints struct {
	CreateCityEndpoint         endpoint.Endpoint
	ListCitiesEndpoint         endpoint.Endpoint
	UpdateCityEndpoint         endpoint.Endpoint
	DeleteCityEndpoint         endpoint.Endpoint
	ListDistrictsEndpoint      endpoint.Endpoint
	ListFiltersEndpoint        endpoint.Endpoint
	ListFilterBuildersEndpoint endpoint.Endpoint
	CreateFilterEndpoint       endpoint.Endpoint
	DeleteFilterEndpoint       endpoint.Endpoint
	CreateDistrictEndpoint     endpoint.Endpoint
	UpdateDistrictEndpoint     endpoint.Endpoint
	DeleteDistrictEndpoint     endpoint.Endpoint
}

// NewEndpoint returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func NewEndpoint(
	service domain.FiltersService,
	serviceSecretKey domain.ServiceSecretKey,
	logger log.Logger,
) Endpoints {
	factory := func(creator func(domain.FiltersService) endpoint.Endpoint, logKey string) endpoint.Endpoint {
		return helpers.SetupEndpoint(creator(service), serviceSecretKey, logger, serviceName, logKey)
	}

	return Endpoints{
		CreateCityEndpoint:         factory(MakeCreateCityEndpoint, "CreateCity"),
		ListCitiesEndpoint:         factory(MakeListCitiesEndpoint, "ListCities"),
		UpdateCityEndpoint:         factory(MakeUpdateCityEndpoint, "UpdateCity"),
		DeleteCityEndpoint:         factory(MakeDeleteCityEndpoint, "DeleteCity"),
		ListDistrictsEndpoint:      factory(MakeListDistrictsEndpoint, "ListDistricts"),
		ListFiltersEndpoint:        factory(MakeListFiltersEndpoint, "ListFilters"),
		ListFilterBuildersEndpoint: factory(MakeListFilterBuildersEndpoint, "ListFilterBuilders"),
		CreateFilterEndpoint:       factory(MakeCreateFilterEndpoint, "CreateFilter"),
		DeleteFilterEndpoint:       factory(MakeDeleteFilterEndpoint, "DeleteFilter"),
		CreateDistrictEndpoint:     factory(MakeCreateDistrictEndpoint, "CreateDistrict"),
		UpdateDistrictEndpoint:     factory(MakeUpdateDistrictEndpoint, "UpdateDistrict"),
		DeleteDistrictEndpoint:     factory(MakeDeleteDistrictEndpoint, "DeleteDistrict"),
	}
}

// MakeCreateCityEndpoint Impl.
func MakeCreateCityEndpoint(service domain.FiltersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createCityRequest)

		resp, err := service.CreateCity(ctx, req.City, helpers.CallerID(ctx))
		return createCityResponse{
			CityID: resp,
			Err:    err,
		}, nil
	}
}

type createCityRequest struct {
	City *domain.City
}

type createCityResponse struct {
	CityID domain.CityID
	Err    error
}

// MakeUpdateCityEndpoint Impl.
func MakeUpdateCityEndpoint(service domain.FiltersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(updateCityRequest)

		err = service.UpdateCity(ctx, req.CityID, req.City, helpers.CallerID(ctx))
		return updateCityResponse{
			Err: err,
		}, nil
	}
}

type updateCityRequest struct {
	CityID domain.CityID
	City   *domain.City
}

type updateCityResponse struct {
	Err error
}

// MakeDeleteCityEndpoint Impl.
func MakeDeleteCityEndpoint(service domain.FiltersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteCityRequest)

		err = service.DeleteCity(ctx, req.CityID, helpers.CallerID(ctx))
		return deleteCityResponse{
			Err: err,
		}, nil

	}
}

type deleteCityRequest struct {
	CityID domain.CityID
}

type deleteCityResponse struct {
	Err error
}

// MakeCreateFilterEndpoint Impl.
func MakeCreateFilterEndpoint(service domain.FiltersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createFilterRequest)

		resp, err := service.CreateFilter(ctx, req.Key, req.Filter, helpers.CallerID(ctx))
		return createFilterResponse{
			Filter: resp,
			Err:    err,
		}, nil
	}
}

type createFilterRequest struct {
	Key    string
	Filter *domain.Filter
}

type createFilterResponse struct {
	Filter *domain.Filter
	Err    error
}

// MakeDeleteFilterEndpoint Impl.
func MakeDeleteFilterEndpoint(service domain.FiltersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteFilterRequest)

		err = service.DeleteFilter(ctx, req.ID, req.Key, helpers.CallerID(ctx))
		return deleteFilterResponse{
			Err: err,
		}, nil
	}
}

type deleteFilterRequest struct {
	ID  int64
	Key string
}

type deleteFilterResponse struct {
	Err error
}

// MakeListDistrictsEndpoint Impl.
func MakeListDistrictsEndpoint(service domain.FiltersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(listDistrictRequest)

		resp, total, err := service.ListDistricts(ctx, req.Criteria, helpers.CallerID(ctx))
		return listDistrictResponse{
			Districts: resp,
			Total:     total,
			Err:       err,
		}, err
	}
}

type listDistrictRequest struct {
	Criteria domain.DistrictSearchCriteria
}

type listDistrictResponse struct {
	Districts []*domain.District
	Total     domain.Total
	Err       error
}

// MakeListCitiesEndpoint Impl.
func MakeListCitiesEndpoint(service domain.FiltersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(listCitiesRequest)
		_ = req

		resp, total, err := service.ListCities(ctx, req.Criteria, helpers.CallerID(ctx))
		return listCitiesResponse{
			Cities: resp,
			Total:  total,
			Err:    err,
		}, err
	}
}

type listCitiesRequest struct {
	Criteria domain.CitySearchCriteria
}

type listCitiesResponse struct {
	Cities []*domain.City
	Total  domain.Total
	Err    error
}

// MakeListFiltersEndpoint Impl.
func MakeListFiltersEndpoint(service domain.FiltersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(listFiltersRequest)
		_ = req

		resp, err := service.ListFilters(ctx, helpers.CallerID(ctx))
		return listFiltersResponse{
			Filter: resp,
			Err:    err,
		}, nil
	}
}

type listFiltersRequest struct {
}

type listFiltersResponse struct {
	Filter map[string][]*domain.Filter
	Err    error
}

// MakeListFilterBuildersEndpoint Impl.
func MakeListFilterBuildersEndpoint(service domain.FiltersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(listFilterBuildersRequest)
		_ = req

		resp, err := service.ListBuilders(ctx, helpers.CallerID(ctx))

		return listFilterBuildersResponse{
			Builders: resp,
			Err:      err,
		}, nil
	}
}

type listFilterBuildersRequest struct {
}

type listFilterBuildersResponse struct {
	Builders []*domain.FilterBuilder
	Err      error
}

// MakeCreateDistrictEndpoint Impl.
func MakeCreateDistrictEndpoint(service domain.FiltersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createDistrictRequest)

		resp, err := service.CreateDistrict(ctx, req.District, helpers.CallerID(ctx))
		return createDistrictResponse{
			DistrictID: resp,
			Err:        err,
		}, nil
	}
}

type createDistrictRequest struct {
	District *domain.District
}

type createDistrictResponse struct {
	DistrictID domain.DistrictID
	Err        error
}

// MakeUpdateDistrictEndpoint Impl.
func MakeUpdateDistrictEndpoint(service domain.FiltersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(updateDistrictRequest)

		err = service.UpdateDistrict(ctx, req.DistrictID, req.District, helpers.CallerID(ctx))
		return updateDistrictResponse{
			Err: err,
		}, nil
	}
}

type updateDistrictRequest struct {
	DistrictID domain.DistrictID
	District   *domain.District
}

type updateDistrictResponse struct {
	Err error
}

// MakeDeleteDistrictEndpoint Impl.
func MakeDeleteDistrictEndpoint(service domain.FiltersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteDistrictRequest)

		err = service.DeleteDistrict(ctx, req.DistrictID, helpers.CallerID(ctx))
		return deleteDistrictResponse{
			Err: err,
		}, nil
	}
}

type deleteDistrictRequest struct {
	DistrictID domain.DistrictID
}

type deleteDistrictResponse struct {
	Err error
}
