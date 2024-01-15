package transport

import (
	"context"

	apiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/identityserviceapi/v1"

	"gitlab.com/zharzhanov/mercury/internal/helpers"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type grpcServer struct {
	sendCode grpctransport.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer.
func NewGRPCServer(endpoints Endpoints, logger log.Logger) apiv1.CodeServiceServer {
	options := helpers.SetupServerOptions(logger)

	return &grpcServer{
		sendCode: grpctransport.NewServer(
			endpoints.SendCodeEndpoint,
			decodeSendCodeRequestV1,
			encodeSendCodeResponseV1,
			options...,
		),
	}
}

func (s *grpcServer) SendCode(ctx context.Context, request *apiv1.SendCodeRequest) (*apiv1.SendCodeResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.sendCode)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.SendCodeResponse), nil
}

func decodeSendCodeRequestV1(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*apiv1.SendCodeRequest)

	return sendCodeRequest{
		Phone: req.Phone,
	}, nil
}

func encodeSendCodeResponseV1(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(sendCodeResponse)
	if resp.Err != nil {
		return &apiv1.SendCodeResponse{}, resp.Err
	}

	return &apiv1.SendCodeResponse{}, nil
}
