package transport

import (
	"context"
	"fmt"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

// Endpoints collects all of the endpoints that compose an add service. It's meant to
// be used as a helper struct, to collect all of the endpoints into a single
// parameter.
type Endpoints struct {
	CreateCottageEndpoint       endpoint.Endpoint
	UpdateCottageEndpoint       endpoint.Endpoint
	DeleteCottageEndpoint       endpoint.Endpoint
	GetCottageEndpoint          endpoint.Endpoint
	ListCottageEndpoint         endpoint.Endpoint
	ListPopularCottagesEndpoint endpoint.Endpoint
	ListCottagesByIDEndpoint    endpoint.Endpoint
	UpdateHousePlanEndpoint     endpoint.Endpoint
	DeleteHousePlanEndpoint     endpoint.Endpoint
	CreateHousePlanEndpoint     endpoint.Endpoint
}

// NewEndpoint returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func NewEndpoint(
	service domain.CottageService,
	serviceSecretKey domain.ServiceSecretKey,
	logger log.Logger,
) Endpoints {
	factory := func(creator func(service domain.CottageService) endpoint.Endpoint, logKey string) endpoint.Endpoint {
		return helpers.SetupEndpoint(creator(service), serviceSecretKey, logger, "News", logKey)
	}

	return Endpoints{
		CreateCottageEndpoint:       factory(MakeCreateCottageEndpoint, "CreateCottage"),
		UpdateCottageEndpoint:       factory(MakeUpdateCottageEndpoint, "UpdateCottage"),
		DeleteCottageEndpoint:       factory(MakeDeleteCottageEndpoint, "DeleteCottage"),
		GetCottageEndpoint:          factory(MakeGetCottageEndpoint, "GetCottage"),
		ListCottageEndpoint:         factory(MakeListCottageEndpoint, "ListCottage"),
		ListPopularCottagesEndpoint: factory(MakeListPopularCottagesEndpoint, "ListPopularCottages"),
		ListCottagesByIDEndpoint:    factory(MakeListCottageByIDEndpoint, "ListCottagesById"),
		UpdateHousePlanEndpoint:     factory(MakeUpdateHousePlanEndpoint, "UpdateHousePlan"),
		DeleteHousePlanEndpoint:     factory(MakeDeleteHousePlanEndpoint, "DeleteHousePlan"),
		CreateHousePlanEndpoint:     factory(MakeCreateHousePlanEndpoint, "CreateHousePlan"),
	}
}

// MakeCreateCottageEndpoint Impl.
func MakeCreateCottageEndpoint(service domain.CottageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createCottageRequest)
		fmt.Println(req.Cottage)
		resp, err := service.CreateCottage(ctx, req.Cottage, helpers.CallerID(ctx))
		return createCottageResponse{
			Cottage: resp,
			Error:   err,
		}, err
	}
}

type createCottageRequest struct {
	Cottage *domain.Cottage
}

type createCottageResponse struct {
	Cottage *domain.Cottage
	Error   error
}

// MakeUpdateCottageEndpoint Impl.
func MakeUpdateCottageEndpoint(service domain.CottageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(updateCottageRequest)
		cottage, err := service.UpdateCottage(ctx, req.ID, req.Cottage, helpers.CallerID(ctx))
		return updateCottageResponse{
			Cottage: cottage,
			Error:   err,
		}, nil
	}
}

type updateCottageRequest struct {
	ID      int64
	Cottage *domain.Cottage
}

type updateCottageResponse struct {
	Cottage *domain.Cottage
	Error   error
}

// MakeDeleteCottageEndpoint Impl.
func MakeDeleteCottageEndpoint(service domain.CottageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteCottageRequest)
		err = service.DeleteCottage(ctx, req.ID, helpers.CallerID(ctx))
		return deleteCottageResponse{Error: err}, nil
	}
}

type deleteCottageRequest struct {
	ID int64
}

type deleteCottageResponse struct {
	Error error
}

// MakeGetCottageEndpoint Impl.
func MakeGetCottageEndpoint(service domain.CottageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getCottageRequest)
		cottage, err := service.GetCottage(ctx, req.ID, helpers.CallerID(ctx))
		return getCottageResponse{
			Cottage: cottage,
			Error:   err,
		}, nil
	}
}

type getCottageRequest struct {
	ID int64
}

type getCottageResponse struct {
	Cottage *domain.Cottage
	Error   error
}

// MakeListCottageEndpoint Impl.
func MakeListCottageEndpoint(service domain.CottageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(listCottageRequest)
		cottage, total, err := service.ListCottage(ctx, req.Criteria, helpers.CallerID(ctx))
		return listCottageResponse{
			Cottage: cottage,
			Total:   total,
			Error:   err,
		}, nil
	}
}

type listCottageRequest struct {
	Criteria domain.CottageSearchCriteria
}

type listCottageResponse struct {
	Cottage []*domain.Cottage
	Total   domain.Total
	Error   error
}

// MakeListPopularCottagesEndpoint Impl.
func MakeListPopularCottagesEndpoint(service domain.CottageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(listPopularCottageRequest)
		cottage, total, err := service.ListPopularCottages(ctx, req.Criteria, helpers.CallerID(ctx))
		return listPopularCottageResponse{
			Cottage: cottage,
			Total:   total,
			Error:   err,
		}, nil
	}
}

type listPopularCottageRequest struct {
	Criteria domain.CottageSearchCriteria
}

type listPopularCottageResponse struct {
	Cottage []*domain.Cottage
	Total   domain.Total
	Error   error
}

// MakeListCottageByIDEndpoint Impl.
func MakeListCottageByIDEndpoint(service domain.CottageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(listCottagesByIDRequest)

		cottage, err := service.ListCottagesByIDs(ctx, req.IDs, helpers.CallerID(ctx))
		return listCottagesByIDResponse{
			Cottage: cottage,
			Error:   err,
		}, nil
	}
}

type listCottagesByIDRequest struct {
	IDs []int64
}

type listCottagesByIDResponse struct {
	Cottage []*domain.Cottage
	Error   error
}

// MakeUpdateHousePlanEndpoint Impl.
func MakeUpdateHousePlanEndpoint(service domain.CottageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(updateHousePlanRequest)
		housePlan, err := service.UpdateHousePlan(ctx, req.HousePlanID, req.HousePlan, helpers.CallerID(ctx))
		return updateHousePlanResponse{
			HousePlan: housePlan,
			Error:     err,
		}, nil
	}
}

type updateHousePlanRequest struct {
	HousePlanID int64
	HousePlan   *domain.HousePlan
}

type updateHousePlanResponse struct {
	HousePlan *domain.HousePlan
	Error     error
}

// MakeCreateHousePlanEndpoint Impl.
func MakeCreateHousePlanEndpoint(service domain.CottageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createHousePlanRequest)
		plan, err := service.CreateHousePlan(ctx, req.HousePlan, helpers.CallerID(ctx))
		return createHousePlanResponse{
			CreatedPlan: plan,
			Error:       err,
		}, nil
	}
}

type createHousePlanRequest struct {
	HousePlan *domain.HousePlan
}

type createHousePlanResponse struct {
	CreatedPlan *domain.HousePlan
	Error       error
}

// MakeDeleteHousePlanEndpoint Impl.
func MakeDeleteHousePlanEndpoint(service domain.CottageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteHousePlanRequest)
		err = service.DeleteHousePlan(ctx, req.HousePlanID, helpers.CallerID(ctx))
		return deleteHousePlanResponse{
			Error: err,
		}, nil
	}
}

type deleteHousePlanRequest struct {
	HousePlanID int64
}

type deleteHousePlanResponse struct {
	Error error
}
