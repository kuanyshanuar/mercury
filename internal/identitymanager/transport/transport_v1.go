package transport

import (
	"context"

	identityapiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/identityserviceapi/v1"

	"gitlab.com/zharzhanov/mercury/internal/helpers"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type grpcServer struct {
	create        grpctransport.Handler
	login         grpctransport.Handler
	validate      grpctransport.Handler
	sendToken     grpctransport.Handler
	resetPassword grpctransport.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer.
func NewGRPCServer(endpoints Endpoints, logger log.Logger) identityapiv1.AuthenticationServiceServer {
	options := helpers.SetupServerOptions(logger)

	return &grpcServer{
		create: grpctransport.NewServer(
			endpoints.CreateUserEndpoint,
			decodeCreateUserRequestV1,
			encodeCreateUserResponseV1,
			options...,
		),
		login: grpctransport.NewServer(
			endpoints.LoginUserEndpoint,
			decodeLoginUserRequestV1,
			encodeLoginUserResponseV1,
			options...,
		),
		validate: grpctransport.NewServer(
			endpoints.ValidateCodeEndpoint,
			decodeValidateUserRequestV2,
			encodeValidateUserRequestV2,
			options...,
		),
		sendToken: grpctransport.NewServer(
			endpoints.SendResetPasswordTokenEndpoint,
			decodeSendTokenRequestV1,
			encodeSendTokenRequestV1,
			options...,
		),
		resetPassword: grpctransport.NewServer(
			endpoints.ResetPasswordEndpoint,
			decodeResetPasswordRequest,
			encodeResetPasswordResponse,
			options...,
		),
	}
}

func (s *grpcServer) CreateUser(ctx context.Context, request *identityapiv1.CreateUserRequest) (*identityapiv1.CreateUserResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.create)
	if err != nil {
		return nil, err
	}
	return rep.(*identityapiv1.CreateUserResponse), nil
}

func decodeCreateUserRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*identityapiv1.CreateUserRequest)

	return createUserRequest{
		User: decodeUserV1(req.User),
	}, nil
}

func encodeCreateUserResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(createUserResponse)
	if resp.Err != nil {
		return &identityapiv1.CreateUserResponse{}, resp.Err
	}
	return &identityapiv1.CreateUserResponse{}, nil
}

func (s *grpcServer) Login(ctx context.Context, request *identityapiv1.LoginRequest) (*identityapiv1.LoginResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.login)
	if err != nil {
		return nil, err
	}
	return rep.(*identityapiv1.LoginResponse), nil
}

func decodeLoginUserRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*identityapiv1.LoginRequest)

	return loginUserRequest{
		Email:    req.Email,
		Password: req.Password,
	}, nil
}

func encodeLoginUserResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(loginUserResponse)
	if resp.Err != nil {
		return &identityapiv1.LoginResponse{}, resp.Err
	}
	return &identityapiv1.LoginResponse{
		User: encodeUserV1(resp.User),
	}, nil
}

func (s *grpcServer) ValidateCode(ctx context.Context, request *identityapiv1.ValidateCodeRequest) (*identityapiv1.ValidateCodeResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.validate)
	if err != nil {
		return nil, err
	}
	return rep.(*identityapiv1.ValidateCodeResponse), nil
}

func decodeValidateUserRequestV2(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*identityapiv1.ValidateCodeRequest)

	return validateCodeRequest{
		Code: req.Code,
	}, nil
}

func encodeValidateUserRequestV2(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(validateCodeResponse)
	if resp.Err != nil {
		return &identityapiv1.ValidateCodeResponse{}, resp.Err
	}
	return &identityapiv1.ValidateCodeResponse{
		UserId: int64(resp.UserID),
		RoleId: int64(resp.RoleID),
	}, nil
}

func (s *grpcServer) SendResetPasswordToken(ctx context.Context, request *identityapiv1.SendResetPasswordTokenRequest) (*identityapiv1.SendResetPasswordTokenResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.sendToken)
	if err != nil {
		return nil, err
	}
	return rep.(*identityapiv1.SendResetPasswordTokenResponse), nil
}

func decodeSendTokenRequestV1(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*identityapiv1.SendResetPasswordTokenRequest)

	return sendResetPasswordTokenRequest{
		Email: req.Email,
	}, nil
}

func encodeSendTokenRequestV1(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(sendResetPasswordTokenResponse)
	if resp.Err != nil {
		return &identityapiv1.SendResetPasswordTokenResponse{}, resp.Err
	}
	return &identityapiv1.SendResetPasswordTokenResponse{}, nil
}

func (s *grpcServer) ResetPasswordToken(ctx context.Context, request *identityapiv1.ResetPasswordTokenRequest) (*identityapiv1.ResetPasswordTokenResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.resetPassword)
	if err != nil {
		return nil, err
	}
	return rep.(*identityapiv1.ResetPasswordTokenResponse), nil
}

func decodeResetPasswordRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*identityapiv1.ResetPasswordTokenRequest)

	return resetPasswordRequest{
		Token:       req.Token,
		NewPassword: req.NewPassword,
	}, nil
}

func encodeResetPasswordResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(resetPasswordResponse)
	if resp.Err != nil {
		return &identityapiv1.ResetPasswordTokenResponse{}, resp.Err
	}
	return &identityapiv1.ResetPasswordTokenResponse{}, nil
}
