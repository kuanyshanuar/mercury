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
	CreateLeadBuilderEndpoint endpoint.Endpoint
	ListLeadBuildersEndpoint  endpoint.Endpoint
	UpdateLeadBuilderEndpoint endpoint.Endpoint
	DeleteLeadBuilderEndpoint endpoint.Endpoint
}

// NewEndpoint returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func NewEndpoint(service domain.LeadBuilderService, serviceSecretKey domain.ServiceSecretKey, logger log.Logger) Endpoints {
	factory := func(creator func(domain.LeadBuilderService) endpoint.Endpoint, logKey string) endpoint.Endpoint {
		return helpers.SetupEndpoint(creator(service), serviceSecretKey, logger, "LeadBuilders", logKey)
	}

	return Endpoints{
		CreateLeadBuilderEndpoint: factory(MakeCreateLeadBuilderEndpoint, "CreateLeadBuilder"),
		ListLeadBuildersEndpoint:  factory(MakeListLeadBuildersEndpoint, "ListLeadBuilders"),
		UpdateLeadBuilderEndpoint: factory(MakeUpdateLeadBuilderEndpoint, "UpdateLeadBuilder"),
		DeleteLeadBuilderEndpoint: factory(MakeDeleteLeadBuilderEndpoint, "DeleteLeadBuilder"),
	}
}

// MakeCreateLeadBuilderEndpoint Impl.
func MakeCreateLeadBuilderEndpoint(service domain.LeadBuilderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(creteLeadBuildersRequest)

		err = service.CreateLeadBuilder(ctx, req.Lead, helpers.CallerID(ctx))

		return creteLeadBuildersResponse{
			Err: err,
		}, nil
	}
}

type creteLeadBuildersRequest struct {
	Lead *domain.LeadBuilder
}

type creteLeadBuildersResponse struct {
	Err error
}

// MakeListLeadBuildersEndpoint Impl.
func MakeListLeadBuildersEndpoint(service domain.LeadBuilderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(listLeadBuildersRequest)

		resp, total, err := service.ListLeadBuilders(ctx, req.Criteria, helpers.CallerID(ctx))
		return listLeadBuildersResponse{
			Leads: resp,
			Total: total,
			Err:   err,
		}, nil
	}
}

type listLeadBuildersRequest struct {
	Criteria domain.LeadBuilderSearchCriteria
}

type listLeadBuildersResponse struct {
	Leads []*domain.LeadBuilder
	Total domain.Total
	Err   error
}

// MakeUpdateLeadBuilderEndpoint Impl.
func MakeUpdateLeadBuilderEndpoint(service domain.LeadBuilderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(updateLeadBuilderRequest)

		err = service.UpdateLeadBuilder(ctx, req.LeadID, req.Lead, helpers.CallerID(ctx))
		return updateLeadBuilderResponse{
			Err: err,
		}, nil
	}
}

type updateLeadBuilderRequest struct {
	LeadID domain.LeadID
	Lead   *domain.LeadBuilder
}

type updateLeadBuilderResponse struct {
	LeadBuilder *domain.LeadBuilder
	Err         error
}

// MakeDeleteLeadBuilderEndpoint Impl.
func MakeDeleteLeadBuilderEndpoint(service domain.LeadBuilderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteLeadBuilderRequest)

		err = service.DeleteLeadBuilder(ctx, req.LeadID, helpers.CallerID(ctx))

		return deleteLeadBuilderResponse{
			Err: err,
		}, nil
	}
}

type deleteLeadBuilderRequest struct {
	LeadID domain.LeadID
}

type deleteLeadBuilderResponse struct {
	Err error
}
