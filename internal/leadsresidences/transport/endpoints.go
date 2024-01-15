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
	CreateLeadResidenceEndpoint endpoint.Endpoint
	ListLeadResidencesEndpoint  endpoint.Endpoint
	UpdateLeadResidenceEndpoint endpoint.Endpoint
	DeleteLeadResidenceEndpoint endpoint.Endpoint
}

// NewEndpoint returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func NewEndpoint(service domain.LeadResidenceService, serviceSecretKey domain.ServiceSecretKey, logger log.Logger) Endpoints {
	factory := func(creator func(domain.LeadResidenceService) endpoint.Endpoint, logKey string) endpoint.Endpoint {
		return helpers.SetupEndpoint(creator(service), serviceSecretKey, logger, "LeadResidences", logKey)
	}

	return Endpoints{
		CreateLeadResidenceEndpoint: factory(MakeCreateLeadResidenceEndpoint, "CreateLeadResidence"),
		ListLeadResidencesEndpoint:  factory(MakeListLeadResidencesEndpoint, "ListLeadResidences"),
		UpdateLeadResidenceEndpoint: factory(MakeUpdateLeadResidenceEndpoint, "UpdateLeadResidence"),
		DeleteLeadResidenceEndpoint: factory(MakeDeleteLeadResidenceEndpoint, "DeleteLeadResidence"),
	}
}

// MakeCreateLeadResidenceEndpoint Impl.
func MakeCreateLeadResidenceEndpoint(service domain.LeadResidenceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(creteLeadResidencesRequest)

		err = service.CreateLeadResidence(ctx, req.Lead, helpers.CallerID(ctx))

		return creteLeadResidencesResponse{
			Err: err,
		}, nil
	}
}

type creteLeadResidencesRequest struct {
	Lead *domain.LeadResidence
}

type creteLeadResidencesResponse struct {
	Err error
}

// MakeListLeadResidencesEndpoint Impl.
func MakeListLeadResidencesEndpoint(service domain.LeadResidenceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(listLeadResidencesRequest)

		resp, total, err := service.ListLeadResidences(ctx, req.Criteria, helpers.CallerID(ctx))
		return listLeadResidencesResponse{
			Leads: resp,
			Total: total,
			Err:   err,
		}, nil
	}
}

type listLeadResidencesRequest struct {
	Criteria domain.LeadResidenceSearchCriteria
}

type listLeadResidencesResponse struct {
	Leads []*domain.LeadResidence
	Total domain.Total
	Err   error
}

// MakeUpdateLeadResidenceEndpoint Impl.
func MakeUpdateLeadResidenceEndpoint(service domain.LeadResidenceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(updateLeadResidenceRequest)

		err = service.UpdateLeadResidence(ctx, req.LeadID, req.Lead, helpers.CallerID(ctx))
		return updateLeadResidenceResponse{
			Err: err,
		}, nil
	}
}

type updateLeadResidenceRequest struct {
	LeadID domain.LeadID
	Lead   *domain.LeadResidence
}

type updateLeadResidenceResponse struct {
	LeadResidence *domain.LeadResidence
	Err           error
}

// MakeDeleteLeadResidenceEndpoint Impl.
func MakeDeleteLeadResidenceEndpoint(service domain.LeadResidenceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteLeadResidenceRequest)

		err = service.DeleteLeadResidence(ctx, req.LeadID, helpers.CallerID(ctx))

		return deleteLeadResidenceResponse{
			Err: err,
		}, nil
	}
}

type deleteLeadResidenceRequest struct {
	LeadID domain.LeadID
}

type deleteLeadResidenceResponse struct {
	Err error
}
