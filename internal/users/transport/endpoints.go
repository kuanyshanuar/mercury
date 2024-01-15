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
	ListUsersEndpoint  endpoint.Endpoint
	UpdateUserEndpoint endpoint.Endpoint
}

// NewEndpoint returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func NewEndpoint(
	service domain.UserService,
	serviceSecretKey domain.ServiceSecretKey,
	logger log.Logger,
) Endpoints {
	factory := func(creator func(domain.UserService) endpoint.Endpoint, logKey string) endpoint.Endpoint {
		return helpers.SetupEndpoint(creator(service), serviceSecretKey, logger, "Users", logKey)
	}

	return Endpoints{
		ListUsersEndpoint:  factory(MakeListUsersEndpoint, "ListUsers"),
		UpdateUserEndpoint: factory(MakeUpdateUserEndpoint, "UpdateUser"),
	}
}

// MakeListUsersEndpoint Impl.
func MakeListUsersEndpoint(service domain.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(listUsersRequest)

		resp, total, err := service.ListUsers(ctx, req.Criteria, helpers.CallerID(ctx))
		return listUsersResponse{
			Users: resp,
			Total: total,
			Err:   err,
		}, nil
	}
}

type listUsersRequest struct {
	Criteria domain.UserSearchCriteria
}

type listUsersResponse struct {
	Users []*domain.User
	Total domain.Total
	Err   error
}

// MakeUpdateUserEndpoint Impl.
func MakeUpdateUserEndpoint(service domain.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(updateUserRequest)

		err = service.UpdateUser(ctx, req.UserID, req.User, helpers.CallerID(ctx))
		return updateUserResponse{
			Err: err,
		}, nil
	}
}

type updateUserRequest struct {
	UserID domain.UserID
	User   *domain.User
}

type updateUserResponse struct {
	Err error
}
