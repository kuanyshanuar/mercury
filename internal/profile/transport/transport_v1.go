package transport

import (
	"context"

	identityapiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/identityserviceapi/v1"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type grpcServer struct {
	getProfile    grpctransport.Handler
	updateProfile grpctransport.Handler
	validatePhone grpctransport.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer.
func NewGRPCServer(endpoints Endpoints, logger log.Logger) identityapiv1.ProfileServiceServer {
	options := helpers.SetupServerOptions(logger)

	return &grpcServer{
		getProfile: grpctransport.NewServer(
			endpoints.GetProfileEndpoint,
			decodeGetProfileRequestV1,
			encodeGetProfileResponseV1,
			options...,
		),
		updateProfile: grpctransport.NewServer(
			endpoints.UpdateProfileEndpoint,
			decodeUpdateProfileRequestV1,
			encodeUpdateProfileResponseV1,
			options...,
		),
		validatePhone: grpctransport.NewServer(
			endpoints.ValidatePhoneEndpoint,
			decodeValidatePhoneRequestV1,
			encodeValidatePhoneResponseV1,
			options...,
		),
	}
}

func (s *grpcServer) GetProfile(ctx context.Context, request *identityapiv1.GetProfileRequest) (*identityapiv1.GetProfileResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.getProfile)
	if err != nil {
		return nil, err
	}
	return rep.(*identityapiv1.GetProfileResponse), nil
}

func decodeGetProfileRequestV1(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*identityapiv1.GetProfileRequest)

	return getProfileRequest{
		UserID: domain.UserID(req.UserId),
	}, nil
}

func encodeGetProfileResponseV1(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(getProfileResponse)
	if resp.Err != nil {
		return &identityapiv1.GetProfileResponse{}, resp.Err
	}
	return &identityapiv1.GetProfileResponse{
		Profile: encodeProfileV1(resp.Profile),
	}, nil
}

func (s *grpcServer) UpdateProfile(ctx context.Context, request *identityapiv1.UpdateProfileRequest) (*identityapiv1.UpdateProfileResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.updateProfile)
	if err != nil {
		return nil, err
	}
	return rep.(*identityapiv1.UpdateProfileResponse), nil
}

func decodeUpdateProfileRequestV1(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*identityapiv1.UpdateProfileRequest)

	return updateProfileRequest{
		UserID:  domain.UserID(req.UserId),
		Profile: decodeProfileV1(req.Profile),
	}, nil
}

func encodeUpdateProfileResponseV1(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(updateProfileResponse)
	if resp.Err != nil {
		return &identityapiv1.UpdateProfileResponse{}, resp.Err
	}
	return &identityapiv1.UpdateProfileResponse{}, nil
}

func (s *grpcServer) ValidatePhone(ctx context.Context, request *identityapiv1.ValidatePhoneRequest) (*identityapiv1.ValidatePhoneResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.validatePhone)
	if err != nil {
		return nil, err
	}
	return rep.(*identityapiv1.ValidatePhoneResponse), nil
}

func decodeValidatePhoneRequestV1(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*identityapiv1.ValidatePhoneRequest)

	return validatePhoneRequest{
		Code: req.Code,
	}, nil
}

func encodeValidatePhoneResponseV1(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(validatePhoneResponse)
	if resp.Err != nil {
		return &identityapiv1.ValidatePhoneResponse{}, resp.Err
	}
	return &identityapiv1.ValidatePhoneResponse{}, nil
}
