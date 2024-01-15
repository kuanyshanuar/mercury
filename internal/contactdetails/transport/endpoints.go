package transport

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

// Service name
const (
	serviceName = "ContactDetails"
)

// Endpoints collects all of the endpoints that compose an add service. It's meant to
// be used as a helper struct, to collect all of the endpoints into a single
// parameter.
type Endpoints struct {
	CreateContactDetailsEndpoint          endpoint.Endpoint
	CreateResidenceContactDetailsEndpoint endpoint.Endpoint
	ListContactDetailsEndpoint            endpoint.Endpoint
	ListResidenceContactDetailsEndpoint   endpoint.Endpoint
}

// NewEndpoint returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func NewEndpoint(service domain.ContactDetailsService, serviceSecretKey domain.ServiceSecretKey, logger log.Logger) Endpoints {

	factory := func(creator func(service domain.ContactDetailsService) endpoint.Endpoint, logKey string) endpoint.Endpoint {
		return helpers.SetupEndpoint(creator(service), serviceSecretKey, logger, serviceName, logKey)
	}

	return Endpoints{
		CreateContactDetailsEndpoint:          factory(MakeCreateContactDetails, "CreateContactDetails"),
		CreateResidenceContactDetailsEndpoint: factory(MakeCreateResidenceContactDetails, "CreateResidenceContactDetails"),
		ListContactDetailsEndpoint:            factory(MakeListContactDetails, "ListContactDetails"),
		ListResidenceContactDetailsEndpoint:   factory(MakeListResidenceContactDetailsEndpoint, "ListResidenceContactDetails"),
	}
}

// MakeCreateContactDetails Impl.
func MakeCreateContactDetails(service domain.ContactDetailsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createContactDetailsRequest)

		err = service.CreateContactDetails(ctx, req.ContactDetails, helpers.CallerID(ctx))
		return createContactDetailsResponse{
			Err: err,
		}, nil
	}
}

type createContactDetailsRequest struct {
	ContactDetails *domain.ContactDetails
}

type createContactDetailsResponse struct {
	Err error
}

// MakeCreateResidenceContactDetails Impl.
func MakeCreateResidenceContactDetails(service domain.ContactDetailsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createResidenceContactDetailsRequest)

		err = service.CreateResidenceContactDetails(ctx, req.ContactDetails, helpers.CallerID(ctx))
		return createResidenceContactDetailsResponse{
			Err: err,
		}, nil
	}
}

type createResidenceContactDetailsRequest struct {
	ContactDetails *domain.ResidenceContactDetails
}

type createResidenceContactDetailsResponse struct {
	Err error
}

// MakeListContactDetails Impl.
func MakeListContactDetails(service domain.ContactDetailsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(listContactDetailsRequest)
		_ = req
		result, err := service.ListContactDetails(ctx, helpers.CallerID(ctx))
		return listContactDetailsResponse{
			ContactDetails: result,
			Err:            err,
		}, nil
	}
}

type listContactDetailsRequest struct {
}

type listContactDetailsResponse struct {
	ContactDetails []*domain.ContactDetails
	Err            error
}

// MakeListResidenceContactDetailsEndpoint Impl.
func MakeListResidenceContactDetailsEndpoint(service domain.ContactDetailsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(listResidenceContactDetailsRequest)

		result, total, err := service.ListResidenceContactDetails(ctx, req.Criteria, helpers.CallerID(ctx))
		return listResidenceContactDetailsResponse{
			ContactDetails: result,
			Total:          total,
			Err:            err,
		}, nil
	}
}

type listResidenceContactDetailsRequest struct {
	Criteria domain.ResidenceContactDetailsSearchCriteria
}

type listResidenceContactDetailsResponse struct {
	ContactDetails []*domain.ResidenceContactDetails
	Total          domain.Total
	Err            error
}
