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
	CreateLeadCottageEndpoint endpoint.Endpoint
	ListLeadCottagesEndpoint  endpoint.Endpoint
	UpdateLeadCottageEndpoint endpoint.Endpoint
	DeleteLeadCottageEndpoint endpoint.Endpoint
	GetLeadCottageEndpoint    endpoint.Endpoint
}

// NewEndpoint returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func NewEndpoint(service domain.LeadCottageService, serviceSecretKey domain.ServiceSecretKey, logger log.Logger) Endpoints {
	factory := func(creator func(cottageService domain.LeadCottageService) endpoint.Endpoint, logKey string) endpoint.Endpoint {
		return helpers.SetupEndpoint(creator(service), serviceSecretKey, logger, "LeadCottages", logKey)
	}

	return Endpoints{
		CreateLeadCottageEndpoint: factory(MakeCreateLeadCottageEndpoint, "CreateLeadCottage"),
		ListLeadCottagesEndpoint:  factory(MakeListLeadCottagesEndpoint, "ListLeadCottages"),
		UpdateLeadCottageEndpoint: factory(MakeUpdateLeadCottageEndpoint, "UpdateLeadCottage"),
		DeleteLeadCottageEndpoint: factory(MakeDeleteLeadCottageEndpoint, "DeleteLeadCottage"),
		GetLeadCottageEndpoint:    factory(MakeGetLeadCottageEndpoint, "GetLeadCottage"),
	}
}

// MakeCreateLeadCottageEndpoint Impl.
func MakeCreateLeadCottageEndpoint(service domain.LeadCottageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createLeadCottagesRequest)

		id, err := service.CreateLeadCottage(ctx, req.Lead, helpers.CallerID(ctx))

		return creteLeadCottagesResponse{
			ID:  id,
			Err: err,
		}, nil
	}
}

type createLeadCottagesRequest struct {
	Lead *domain.LeadCottage
}

type creteLeadCottagesResponse struct {
	ID  int64
	Err error
}

// MakeListLeadCottagesEndpoint Impl.
func MakeListLeadCottagesEndpoint(service domain.LeadCottageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(listLeadCottagesRequest)

		resp, total, err := service.ListLeadCottage(ctx, req.Criteria, helpers.CallerID(ctx))
		return listLeadCottagesResponse{
			Leads: resp,
			Total: total,
			Err:   err,
		}, nil
	}
}

type listLeadCottagesRequest struct {
	Criteria domain.LeadCottageSearchCriteria
}

type listLeadCottagesResponse struct {
	Leads []*domain.LeadCottage
	Total domain.Total
	Err   error
}

// MakeUpdateLeadCottageEndpoint Impl.
func MakeUpdateLeadCottageEndpoint(service domain.LeadCottageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(updateLeadCottageRequest)

		resp, err := service.UpdateLeadCottage(ctx, req.LeadID, req.Lead, helpers.CallerID(ctx))
		return updateLeadCottageResponse{
			LeadCottage: resp,
			Err:         err,
		}, nil
	}
}

type updateLeadCottageRequest struct {
	LeadID int64
	Lead   *domain.LeadCottage
}

type updateLeadCottageResponse struct {
	LeadCottage *domain.LeadCottage
	Err         error
}

// MakeDeleteLeadCottageEndpoint Impl.
func MakeDeleteLeadCottageEndpoint(service domain.LeadCottageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteLeadCottageRequest)

		err = service.DeleteLeadCottage(ctx, req.LeadID, helpers.CallerID(ctx))

		return deleteLeadCottageResponse{
			Err: err,
		}, nil
	}
}

type deleteLeadCottageRequest struct {
	LeadID int64
}

type deleteLeadCottageResponse struct {
	Err error
}

// MakeGetLeadCottageEndpoint Impl.
func MakeGetLeadCottageEndpoint(service domain.LeadCottageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getLeadCottageRequest)

		lead, err := service.GetLeadCottage(ctx, req.LeadID, helpers.CallerID(ctx))

		return getLeadCottageResponse{
			Lead: lead,
			Err:  err,
		}, nil
	}
}

type getLeadCottageRequest struct {
	LeadID int64
}

type getLeadCottageResponse struct {
	Lead *domain.LeadCottage
	Err  error
}
