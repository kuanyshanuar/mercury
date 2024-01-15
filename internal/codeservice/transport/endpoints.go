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
	SendCodeEndpoint endpoint.Endpoint
}

// NewEndpoint returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func NewEndpoint(
	service domain.CodeService,
	serviceSecretKey domain.ServiceSecretKey,
	logger log.Logger,
) Endpoints {
	factory := func(creator func(service domain.CodeService) endpoint.Endpoint, logKey string) endpoint.Endpoint {
		return helpers.SetupEndpoint(creator(service), serviceSecretKey, logger, "CodeService", logKey)
	}

	return Endpoints{
		SendCodeEndpoint: factory(MakeSendCodeEndpoint, "SendCode"),
	}
}

// MakeSendCodeEndpoint Impl.
func MakeSendCodeEndpoint(service domain.CodeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(sendCodeRequest)

		err = service.SendCode(ctx, req.Phone)
		return sendCodeResponse{
			Err: err,
		}, nil
	}
}

type sendCodeRequest struct {
	Phone string
}

type sendCodeResponse struct {
	Err error
}
