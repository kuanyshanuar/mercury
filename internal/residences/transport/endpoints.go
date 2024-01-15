package transport

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

// Endpoints collects all of the endpoints that compose an add service. It's meant to
// be used as a helper struct, to collect all of the endpoints into a single
// parameter.
type Endpoints struct {
	CreateResidenceEndpoint       endpoint.Endpoint
	ListResidencesEndpoint        endpoint.Endpoint
	ListPopularResidencesEndpoint endpoint.Endpoint
	ListResidencesByIDsEndpoint   endpoint.Endpoint
	GetResidenceEndpoint          endpoint.Endpoint
	UpdateResidenceEndpoint       endpoint.Endpoint
	DeleteResidenceEndpoint       endpoint.Endpoint
	CreateFlatPlanEndpoint        endpoint.Endpoint
	UpdateFlatPlanEndpoint        endpoint.Endpoint
	DeleteFlatPlanEndpoint        endpoint.Endpoint
}

// NewEndpoint returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func NewEndpoint(
	service domain.ResidencesService,
	serviceSecretKey domain.ServiceSecretKey,
	logger log.Logger,
) Endpoints {
	factory := func(creator func(domain.ResidencesService) endpoint.Endpoint, logKey string) endpoint.Endpoint {
		return helpers.SetupEndpoint(creator(service), serviceSecretKey, logger, "Residences", logKey)
	}

	return Endpoints{
		CreateResidenceEndpoint:       factory(MakeCreateResidenceEndpoint, "CreateResidence"),
		ListResidencesEndpoint:        factory(MakeListResidencesEndpoint, "ListResidences"),
		ListPopularResidencesEndpoint: factory(MakeListPopularResidencesEndpoint, "ListPopularResidences"),
		ListResidencesByIDsEndpoint:   factory(MakeListResidencesByIDsEndpoint, "ListResidencesByIDs"),
		GetResidenceEndpoint:          factory(MakeGetResidenceEndpoint, "GetResidence"),
		UpdateResidenceEndpoint:       factory(MakeUpdateResidenceEndpoint, "UpdateResidence"),
		DeleteResidenceEndpoint:       factory(MakeDeleteResidenceEndpoint, "DeleteResidence"),
		CreateFlatPlanEndpoint:        factory(MakeCreateFlatPlanEndpoint, "CreateFlatPlan"),
		UpdateFlatPlanEndpoint:        factory(MakeUpdateFlatPlanEndpoint, "UpdateFlatPlan"),
		DeleteFlatPlanEndpoint:        factory(MakeDeleteFlatPlanEndpoint, "DeleteFlatPlan"),
	}
}

// MakeCreateResidenceEndpoint - Impl.
func MakeCreateResidenceEndpoint(service domain.ResidencesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createResidenceRequest)

		resp, err := service.CreateResidence(ctx, req.Residence, helpers.CallerID(ctx))
		return createResidenceResponse{
			Residence: resp,
			Err:       err,
		}, nil
	}
}

type createResidenceRequest struct {
	Residence *domain.Residence
}

type createResidenceResponse struct {
	Residence *domain.Residence
	Err       error
}

// MakeListResidencesEndpoint - Impl.
func MakeListResidencesEndpoint(service domain.ResidencesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(listResidencesRequest)

		resp, total, err := service.ListResidences(ctx, req.Criteria, helpers.CallerID(ctx))
		return listResidencesResponse{
			Residences: resp,
			Total:      total,
			Err:        err,
		}, nil
	}
}

type listResidencesRequest struct {
	Criteria domain.ResidenceSearchCriteria
}

type listResidencesResponse struct {
	Residences []*domain.Residence
	Total      domain.Total
	Err        error
}

// MakeListPopularResidencesEndpoint Impl.
func MakeListPopularResidencesEndpoint(service domain.ResidencesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(listPopularResidencesRequest)

		resp, total, err := service.ListPopularResidences(ctx, req.Criteria, helpers.CallerID(ctx))
		return listPopularResidencesResponse{
			Residences: resp,
			Total:      total,
			Err:        err,
		}, nil
	}
}

type listPopularResidencesRequest struct {
	Criteria domain.ResidenceSearchCriteria
}

type listPopularResidencesResponse struct {
	Residences []*domain.Residence
	Total      domain.Total
	Err        error
}

// MakeListResidencesByIDsEndpoint Impl.
func MakeListResidencesByIDsEndpoint(service domain.ResidencesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(listResidencesByIDsRequest)

		resp, err := service.ListResidencesByIDs(ctx, req.ResidenceIDs, helpers.CallerID(ctx))
		return listResidencesByIDsResponse{
			Residences: resp,
			Err:        err,
		}, nil
	}
}

type listResidencesByIDsRequest struct {
	ResidenceIDs []domain.ResidenceID
}

type listResidencesByIDsResponse struct {
	Residences []*domain.Residence
	Err        error
}

// MakeGetResidenceEndpoint - Impl.
func MakeGetResidenceEndpoint(service domain.ResidencesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getResidenceRequest)

		resp, err := service.GetResidence(ctx, req.ResidenceID, helpers.CallerID(ctx))

		return getResidenceResponse{
			Residence: resp,
			Err:       err,
		}, nil
	}
}

type getResidenceRequest struct {
	ResidenceID domain.ResidenceID
}

type getResidenceResponse struct {
	Residence *domain.Residence
	Err       error
}

// MakeUpdateResidenceEndpoint - Impl.
func MakeUpdateResidenceEndpoint(service domain.ResidencesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(updateResidenceRequest)

		resp, err := service.UpdateResidence(ctx, req.ResidenceID, req.Residence, helpers.CallerID(ctx))
		return updateResidenceResponse{
			Residence: resp,
			Err:       err,
		}, nil
	}
}

type updateResidenceRequest struct {
	ResidenceID domain.ResidenceID
	Residence   *domain.Residence
}

type updateResidenceResponse struct {
	Residence *domain.Residence
	Err       error
}

// MakeDeleteResidenceEndpoint - Impl.
func MakeDeleteResidenceEndpoint(service domain.ResidencesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteResidenceRequest)

		err = service.DeleteResidence(ctx, req.ResidenceID, helpers.CallerID(ctx))
		return deleteResidenceResponse{
			Err: err,
		}, nil
	}
}

type deleteResidenceRequest struct {
	ResidenceID domain.ResidenceID
}

type deleteResidenceResponse struct {
	Err error
}

// MakeCreateFlatPlanEndpoint Impl.
func MakeCreateFlatPlanEndpoint(service domain.ResidencesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createFlatPlanRequest)

		resp, err := service.CreateFlatPlan(ctx, req.FlatPlan, helpers.CallerID(ctx))
		return createFlatPlanResponse{
			FlatPlan: resp,
			Err:      err,
		}, nil
	}
}

type createFlatPlanRequest struct {
	FlatPlan *domain.FlatPlan
}

type createFlatPlanResponse struct {
	FlatPlan *domain.FlatPlan
	Err      error
}

// MakeUpdateFlatPlanEndpoint Impl.
func MakeUpdateFlatPlanEndpoint(service domain.ResidencesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(updateFlatPlanRequest)

		resp, err := service.UpdateFlatPlan(ctx, req.ID, req.FlatPlan, helpers.CallerID(ctx))
		return updateFlatPlanResponse{
			FlatPlan: resp,
			Err:      err,
		}, nil
	}
}

type updateFlatPlanRequest struct {
	ID       domain.FlatPlanID
	FlatPlan *domain.FlatPlan
}

type updateFlatPlanResponse struct {
	FlatPlan *domain.FlatPlan
	Err      error
}

// MakeDeleteFlatPlanEndpoint Impl.
func MakeDeleteFlatPlanEndpoint(service domain.ResidencesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteFlatPlanRequest)

		err = service.DeleteFlatPlan(ctx, req.ID, helpers.CallerID(ctx))
		return deleteFlatPlanResponse{
			Err: err,
		}, nil
	}
}

type deleteFlatPlanRequest struct {
	ID domain.FlatPlanID
}

type deleteFlatPlanResponse struct {
	Err error
}
