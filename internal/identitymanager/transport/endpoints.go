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
	serviceName = "Identity"
)

// Endpoints collects all of the endpoints that compose an add service. It's meant to
// be used as a helper struct, to collect all of the endpoints into a single
// parameter.
type Endpoints struct {
	CreateUserEndpoint             endpoint.Endpoint
	LoginUserEndpoint              endpoint.Endpoint
	ValidateCodeEndpoint           endpoint.Endpoint
	SendResetPasswordTokenEndpoint endpoint.Endpoint
	ResetPasswordEndpoint          endpoint.Endpoint
}

// NewEndpoint returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func NewEndpoint(
	service domain.IdentityManagerService,
	serviceSecretKey domain.ServiceSecretKey,
	logger log.Logger,
) Endpoints {
	factory := func(creator func(domain.IdentityManagerService) endpoint.Endpoint, logKey string) endpoint.Endpoint {
		return helpers.SetupEndpoint(creator(service), serviceSecretKey, logger, serviceName, logKey)
	}

	return Endpoints{
		CreateUserEndpoint:             factory(MakeRegisterEndpoint, "CreateUser"),
		LoginUserEndpoint:              factory(MakeLoginUserEndpoint, "LoginUser"),
		ValidateCodeEndpoint:           factory(MakeValidateCodeEndpoint, "ValidateCode"),
		SendResetPasswordTokenEndpoint: factory(MakeSendResetPasswordTokenEndpoint, "SendResetPasswordToken"),
		ResetPasswordEndpoint:          factory(MakeResetPasswordEndpoint, "ResetPassword"),
	}
}

// MakeRegisterEndpoint - Impl.
func MakeRegisterEndpoint(service domain.IdentityManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createUserRequest)

		err = service.CreateUser(ctx, req.User)
		return createUserResponse{
			Err: err,
		}, nil
	}
}

type createUserRequest struct {
	User *domain.User
}

type createUserResponse struct {
	Err error
}

// MakeLoginUserEndpoint - Impl.
func MakeLoginUserEndpoint(service domain.IdentityManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(loginUserRequest)

		resp, err := service.ValidateUser(ctx, req.Email, req.Password)
		return loginUserResponse{
			User: resp,
			Err:  err,
		}, nil
	}
}

type loginUserRequest struct {
	Email    string
	Password string
}

type loginUserResponse struct {
	User *domain.User
	Err  error
}

// MakeValidateCodeEndpoint - Impl.
func MakeValidateCodeEndpoint(service domain.IdentityManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(validateCodeRequest)

		userID, roleID, err := service.ValidateCode(ctx, req.Code)
		return validateCodeResponse{
			UserID: userID,
			RoleID: roleID,
			Err:    err,
		}, nil
	}
}

type validateCodeRequest struct {
	Code string
}

type validateCodeResponse struct {
	UserID domain.UserID
	RoleID domain.RoleID
	Err    error
}

// MakeSendResetPasswordTokenEndpoint Impl.
func MakeSendResetPasswordTokenEndpoint(service domain.IdentityManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(sendResetPasswordTokenRequest)

		err = service.SendResetPasswordToken(ctx, req.Email)
		return sendResetPasswordTokenResponse{
			Err: err,
		}, nil
	}
}

type sendResetPasswordTokenRequest struct {
	Email string
}

type sendResetPasswordTokenResponse struct {
	Err error
}

// MakeResetPasswordEndpoint Impl.
func MakeResetPasswordEndpoint(service domain.IdentityManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(resetPasswordRequest)

		err = service.ResetPassword(ctx, req.Token, req.NewPassword)
		return resetPasswordResponse{
			Err: err,
		}, nil
	}
}

type resetPasswordRequest struct {
	Token       string
	NewPassword string
}

type resetPasswordResponse struct {
	Err error
}
