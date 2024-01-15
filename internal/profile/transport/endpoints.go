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
	GetProfileEndpoint    endpoint.Endpoint
	UpdateProfileEndpoint endpoint.Endpoint
	ValidatePhoneEndpoint endpoint.Endpoint
}

// NewEndpoint returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func NewEndpoint(service domain.ProfileService, serviceSecretKey domain.ServiceSecretKey, logger log.Logger) Endpoints {
	factory := func(creator func(domain.ProfileService) endpoint.Endpoint, logKey string) endpoint.Endpoint {
		return helpers.SetupEndpoint(creator(service), serviceSecretKey, logger, "Identity", logKey)
	}

	return Endpoints{
		GetProfileEndpoint:    factory(MakeGetProfileEndpoint, "GetProfile"),
		UpdateProfileEndpoint: factory(MakeUpdateProfileEndpoint, "UpdateProfile"),
		ValidatePhoneEndpoint: factory(MakeValidatePhoneEndpoint, "ValidatePhone"),
	}
}

// MakeGetProfileEndpoint Impl.
func MakeGetProfileEndpoint(service domain.ProfileService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getProfileRequest)

		resp, err := service.GetProfile(ctx, req.UserID, helpers.CallerID(ctx))
		return getProfileResponse{
			Profile: resp,
			Err:     err,
		}, nil
	}
}

type getProfileRequest struct {
	UserID domain.UserID
}

type getProfileResponse struct {
	Profile *domain.Profile
	Err     error
}

// MakeUpdateProfileEndpoint Impl.
func MakeUpdateProfileEndpoint(service domain.ProfileService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(updateProfileRequest)

		err = service.UpdateProfile(ctx, req.UserID, req.Profile, helpers.CallerID(ctx))
		return updateProfileResponse{
			Err: err,
		}, nil
	}
}

type updateProfileRequest struct {
	UserID  domain.UserID
	Profile *domain.Profile
}

type updateProfileResponse struct {
	Err error
}

// MakeValidatePhoneEndpoint Impl.
func MakeValidatePhoneEndpoint(service domain.ProfileService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(validatePhoneRequest)

		err = service.ValidatePhone(ctx, req.Code, helpers.CallerID(ctx))
		return validatePhoneResponse{
			Err: err,
		}, nil
	}
}

type validatePhoneRequest struct {
	Code string
}

type validatePhoneResponse struct {
	Err error
}
