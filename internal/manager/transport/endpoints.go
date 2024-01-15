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
	CreateManagerEndpoint endpoint.Endpoint
	GetManagerEndpoint    endpoint.Endpoint
	ListManagersEndpoint  endpoint.Endpoint
	UpdateManagerEndpoint endpoint.Endpoint
	DeleteManagerEndpoint endpoint.Endpoint
}

// NewEndpoint returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func NewEndpoint(
	service domain.ManagerService,
	serviceSecretKey domain.ServiceSecretKey,
	logger log.Logger,
) Endpoints {
	factory := func(creator func(domain.ManagerService) endpoint.Endpoint, logKey string) endpoint.Endpoint {
		return helpers.SetupEndpoint(creator(service), serviceSecretKey, logger, "Manager", logKey)
	}

	return Endpoints{
		CreateManagerEndpoint: factory(MakeCreateManagerEndpoint, "CreateManager"),
		GetManagerEndpoint:    factory(MakeGetManagerEndpoint, "GetManager"),
		ListManagersEndpoint:  factory(MakeListManagersEndpoint, "ListManagers"),
		UpdateManagerEndpoint: factory(MakeUpdateManagerEndpoint, "UpdateManager"),
		DeleteManagerEndpoint: factory(MakeDeleteManagerEndpoint, "DeleteManager"),
	}
}

// MakeCreateManagerEndpoint Impl.
func MakeCreateManagerEndpoint(service domain.ManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createManagerRequest)

		res, err := service.CreateManager(
			ctx,
			req.Manager,
			helpers.CallerID(ctx),
		)
		return createManagerResponse{
			ID:  res,
			Err: err,
		}, nil
	}
}

type createManagerRequest struct {
	Manager *domain.Manager
}

type createManagerResponse struct {
	ID  domain.ManagerID
	Err error
}

// MakeGetManagerEndpoint Impl.
func MakeGetManagerEndpoint(service domain.ManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getManagerRequest)

		resp, err := service.GetManager(
			ctx,
			req.ManagerID,
			helpers.CallerID(ctx),
		)
		return getManagerResponse{
			Manager: resp,
			Err:     err,
		}, nil
	}
}

type getManagerRequest struct {
	ManagerID domain.ManagerID
}

type getManagerResponse struct {
	Manager *domain.Manager
	Err     error
}

// MakeListManagersEndpoint Impl.
func MakeListManagersEndpoint(service domain.ManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(listManagerRequest)

		resp, total, err := service.ListManagers(
			ctx,
			req.Criteria,
			helpers.CallerID(ctx),
		)
		return listManagerResponse{
			Managers: resp,
			Total:    total,
			Err:      err,
		}, nil
	}
}

type listManagerRequest struct {
	Criteria domain.ManagerSearchCriteria
}

type listManagerResponse struct {
	Managers []*domain.Manager
	Total    domain.Total
	Err      error
}

// MakeUpdateManagerEndpoint Impl.
func MakeUpdateManagerEndpoint(service domain.ManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(updateManagerRequest)

		err = service.UpdateManager(
			ctx,
			req.ManagerID,
			req.Manager,
			helpers.CallerID(ctx),
		)
		return updateManagerResponse{
			Err: err,
		}, nil
	}
}

type updateManagerRequest struct {
	ManagerID domain.ManagerID
	Manager   *domain.Manager
}

type updateManagerResponse struct {
	Err error
}

// MakeDeleteManagerEndpoint Impl.
func MakeDeleteManagerEndpoint(service domain.ManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteManagerRequest)

		err = service.DeleteManager(
			ctx,
			req.ManagerID,
			helpers.CallerID(ctx),
		)
		return deleteManagerResponse{
			Err: err,
		}, nil
	}
}

type deleteManagerRequest struct {
	ManagerID domain.ManagerID
}

type deleteManagerResponse struct {
	Err error
}
