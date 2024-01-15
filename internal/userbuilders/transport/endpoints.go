package transport

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

const (
	serviceName = "UserBuilderService"
)

// Endpoints collects all of the endpoints that compose an add service. It's meant to
// be used as a helper struct, to collect all of the endpoints into a single
// parameter.
type Endpoints struct {
	SubscribeEndpoint    endpoint.Endpoint
	UnsubscribeEndpoint  endpoint.Endpoint
	ListBuildersEndpoint endpoint.Endpoint
}

// NewEndpoint returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func NewEndpoint(
	service domain.UserBuilderService,
	serviceSecretKey domain.ServiceSecretKey,
	logger log.Logger,
) Endpoints {
	factory := func(creator func(domain.UserBuilderService) endpoint.Endpoint, logKey string) endpoint.Endpoint {
		return helpers.SetupEndpoint(creator(service), serviceSecretKey, logger, serviceName, logKey)
	}

	return Endpoints{
		SubscribeEndpoint:    factory(MakeSubscribeEndpoint, "Subscribe"),
		UnsubscribeEndpoint:  factory(MakeUnsubscribeEndpoint, "Unsubscribe"),
		ListBuildersEndpoint: factory(MakeListBuildersEndpoint, "ListBuilders"),
	}
}

// MakeSubscribeEndpoint Impl.
func MakeSubscribeEndpoint(service domain.UserBuilderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(subscribeRequest)

		err = service.Subscribe(ctx, req.SubscriberID, req.BuilderID, helpers.CallerID(ctx))

		return subscribeResponse{
			Err: err,
		}, nil
	}
}

type subscribeRequest struct {
	SubscriberID int64
	BuilderID    int64
}

type subscribeResponse struct {
	Err error
}

// MakeUnsubscribeEndpoint Impl.
func MakeUnsubscribeEndpoint(service domain.UserBuilderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(unsubscribeRequest)

		err = service.Unsubscribe(ctx, req.SubscriberID, req.BuilderID, helpers.CallerID(ctx))

		return unsubscribeResponse{
			Err: err,
		}, nil
	}
}

type unsubscribeRequest struct {
	SubscriberID int64
	BuilderID    int64
}

type unsubscribeResponse struct {
	Err error
}

// MakeListBuildersEndpoint Impl.
func MakeListBuildersEndpoint(service domain.UserBuilderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(listBuildersRequest)

		resp, total, err := service.ListBuilders(ctx, req.SubscriberID, req.Criteria, helpers.CallerID(ctx))

		return listBuildersResponse{
			Builders: resp,
			Total:    total,
			Err:      err,
		}, nil
	}
}

type listBuildersRequest struct {
	Criteria     domain.UserBuilderSearchCriteria
	SubscriberID int64
}

type listBuildersResponse struct {
	Builders []*domain.Builder
	Total    domain.Total
	Err      error
}
