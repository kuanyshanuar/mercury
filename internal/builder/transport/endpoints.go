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
	CreateBuilderEndpoint endpoint.Endpoint
	GetBuilderEndpoint    endpoint.Endpoint
	ListBuildersEndpoint  endpoint.Endpoint
	UpdateBuilderEndpoint endpoint.Endpoint
	DeleteBuilderEndpoint endpoint.Endpoint
}

// NewEndpoint returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func NewEndpoint(
	service domain.BuilderService,
	serviceSecretKey domain.ServiceSecretKey,
	logger log.Logger,
) Endpoints {
	factory := func(creator func(domain.BuilderService) endpoint.Endpoint, logKey string) endpoint.Endpoint {
		return helpers.SetupEndpoint(creator(service), serviceSecretKey, logger, "Builder", logKey)
	}

	return Endpoints{
		CreateBuilderEndpoint: factory(MakeCreateBuilderEndpoint, "CreateBuilder"),
		GetBuilderEndpoint:    factory(MakeGetBuilderEndpoint, "GetBuilder"),
		ListBuildersEndpoint:  factory(MakeListBuildersEndpoint, "ListBuilders"),
		UpdateBuilderEndpoint: factory(MakeUpdateBuilderEndpoint, "UpdateBuilder"),
		DeleteBuilderEndpoint: factory(MakeDeleteBuilderEndpoint, "DeleteBuilder"),
	}
}

// MakeCreateBuilderEndpoint Impl.
func MakeCreateBuilderEndpoint(service domain.BuilderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createBuilderRequest)

		resp, err := service.CreateBuilder(
			ctx,
			req.Builder,
			helpers.CallerID(ctx),
		)
		return createBuilderResponse{
			BuilderID: resp,
			Err:       err,
		}, nil
	}
}

type createBuilderRequest struct {
	Builder *domain.Builder
}

type createBuilderResponse struct {
	BuilderID domain.BuilderID
	Err       error
}

// MakeGetBuilderEndpoint Impl.
func MakeGetBuilderEndpoint(service domain.BuilderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getBuilderRequest)

		resp, err := service.GetBuilder(
			ctx,
			req.BuilderID,
			helpers.CallerID(ctx),
		)
		return getBuilderResponse{
			Builder: resp,
			Err:     err,
		}, nil
	}
}

type getBuilderRequest struct {
	BuilderID domain.BuilderID
}

type getBuilderResponse struct {
	Builder *domain.Builder
	Err     error
}

// MakeListBuildersEndpoint Impl.
func MakeListBuildersEndpoint(service domain.BuilderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(listBuilderRequest)

		resp, total, err := service.ListBuilders(
			ctx,
			req.Criteria,
			helpers.CallerID(ctx),
		)
		return listBuilderResponse{
			Builders: resp,
			Total:    total,
			Err:      err,
		}, nil
	}
}

type listBuilderRequest struct {
	Criteria domain.BuilderSearchCriteria
}

type listBuilderResponse struct {
	Builders []*domain.Builder
	Total    domain.Total
	Err      error
}

// MakeUpdateBuilderEndpoint Impl.
func MakeUpdateBuilderEndpoint(service domain.BuilderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(updateBuilderRequest)

		err = service.UpdateBuilder(
			ctx,
			req.BuilderID,
			req.Builder,
			helpers.CallerID(ctx),
		)
		return updateBuilderResponse{
			Err: err,
		}, nil
	}
}

type updateBuilderRequest struct {
	BuilderID domain.BuilderID
	Builder   *domain.Builder
}

type updateBuilderResponse struct {
	Err error
}

// MakeDeleteBuilderEndpoint Impl.
func MakeDeleteBuilderEndpoint(service domain.BuilderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteBuilderRequest)

		err = service.DeleteBuilder(
			ctx,
			req.BuilderID,
			helpers.CallerID(ctx),
		)
		return deleteBuilderResponse{
			Err: err,
		}, nil
	}
}

type deleteBuilderRequest struct {
	BuilderID domain.BuilderID
}

type deleteBuilderResponse struct {
	Err error
}
