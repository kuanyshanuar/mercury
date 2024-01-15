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
	AllowPermissionEndpoint endpoint.Endpoint
}

// NewEndpoint returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func NewEndpoint(service domain.PermissionsService, serviceSecretKey domain.ServiceSecretKey, logger log.Logger) Endpoints {
	factory := func(creator func(domain.PermissionsService) endpoint.Endpoint, logKey string) endpoint.Endpoint {
		return helpers.SetupEndpoint(creator(service), serviceSecretKey, logger, "Permissions", logKey)
	}

	return Endpoints{
		AllowPermissionEndpoint: factory(MakeAllowEndpoint, "Allow"),
	}
}

// MakeAllowEndpoint Impl.
func MakeAllowEndpoint(service domain.PermissionsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(allowPermissionRequest)

		resp, err := service.Allow(ctx, req.PermissionKey, req.UserID, req.RoleID, helpers.CallerID(ctx))
		return allowPermissionResponse{
			isAllowed: resp,
			Err:       err,
		}, nil
	}
}

type allowPermissionRequest struct {
	PermissionKey string
	UserID        domain.UserID
	RoleID        domain.RoleID
}

type allowPermissionResponse struct {
	isAllowed bool
	Err       error
}
