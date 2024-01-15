package transport

import (
	"context"

	userresidencesapiv1 "gitlab.com/zharzhanov/mercury/api/generated/gitlab.com.mercury/residenceserviceapi/v1"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type grpcServer struct {
	addResidenceToFavourite      grpctransport.Handler
	deleteResidenceFromFavourite grpctransport.Handler
	listFavouriteResidences      grpctransport.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer.
func NewGRPCServer(endpoints Endpoints, logger log.Logger) userresidencesapiv1.UserResidenceServiceServer {
	options := helpers.SetupServerOptions(logger)

	return &grpcServer{
		addResidenceToFavourite: grpctransport.NewServer(
			endpoints.AddResidenceToFavouritesEndpoint,
			decodeAddResidenceToFavouriteRequest,
			encodeAddResidenceToFavouriteResponse,
			options...,
		),
		deleteResidenceFromFavourite: grpctransport.NewServer(
			endpoints.DeleteResidenceFromFavouritesEndpoint,
			decodeDeleteResidenceFromFavouriteRequest,
			encodeDeleteResidenceFromFavouriteResponse,
			options...,
		),
		listFavouriteResidences: grpctransport.NewServer(
			endpoints.ListFavouriteResidencesEndpoint,
			decodeListFavouriteResidenceRequest,
			encodeListFavouriteResidencesResponse,
			options...,
		),
	}
}

func decodeAddResidenceToFavouriteRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*userresidencesapiv1.AddResidenceToFavouritesRequest)

	return addResidenceToFavouritesRequest{
		UserID:      domain.UserID(req.UserId),
		ResidenceID: domain.ResidenceID(req.ResidenceId),
	}, nil
}

func encodeAddResidenceToFavouriteResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(addResidenceToFavouritesResponse)
	if resp.Err != nil {
		return &userresidencesapiv1.AddResidenceToFavouritesResponse{}, resp.Err
	}
	return &userresidencesapiv1.AddResidenceToFavouritesResponse{}, nil
}

func (s *grpcServer) AddResidenceToFavourites(ctx context.Context, request *userresidencesapiv1.AddResidenceToFavouritesRequest) (*userresidencesapiv1.AddResidenceToFavouritesResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.addResidenceToFavourite)
	if err != nil {
		return nil, err
	}
	return rep.(*userresidencesapiv1.AddResidenceToFavouritesResponse), nil
}

func (s *grpcServer) DeleteResidenceFromFavourites(ctx context.Context, request *userresidencesapiv1.DeleteResidenceFromFavouritesRequest) (*userresidencesapiv1.DeleteResidenceFromFavouritesResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.deleteResidenceFromFavourite)
	if err != nil {
		return nil, err
	}
	return rep.(*userresidencesapiv1.DeleteResidenceFromFavouritesResponse), nil
}

func decodeDeleteResidenceFromFavouriteRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*userresidencesapiv1.DeleteResidenceFromFavouritesRequest)

	return deleteResidenceFromFavouritesRequest{
		UserID:      domain.UserID(req.UserId),
		ResidenceID: domain.ResidenceID(req.ResidenceId),
	}, nil
}

func encodeDeleteResidenceFromFavouriteResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(deleteResidenceFromFavouritesResponse)
	if resp.Err != nil {
		return &userresidencesapiv1.DeleteResidenceFromFavouritesResponse{}, resp.Err
	}
	return &userresidencesapiv1.DeleteResidenceFromFavouritesResponse{}, nil
}

func (s *grpcServer) ListFavouriteResidences(ctx context.Context, request *userresidencesapiv1.ListFavouriteResidencesRequest) (*userresidencesapiv1.ListFavouriteResidencesResponse, error) {
	rep, err := helpers.ServeGrpc(ctx, request, s.listFavouriteResidences)
	if err != nil {
		return nil, err
	}
	return rep.(*userresidencesapiv1.ListFavouriteResidencesResponse), nil
}

func decodeListFavouriteResidenceRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*userresidencesapiv1.ListFavouriteResidencesRequest)

	return listFavouriteResidencesRequest{
		UserID:   domain.UserID(req.UserId),
		Criteria: decodeFavouriteResidenceSearchCriteria(req.Criteria),
	}, nil
}

func encodeListFavouriteResidencesResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(listFavouriteResidencesResponse)
	if resp.Err != nil {
		return &userresidencesapiv1.ListFavouriteResidencesResponse{}, resp.Err
	}
	return &userresidencesapiv1.ListFavouriteResidencesResponse{
		ResidenceIds: resp.ResidenceIDs,
		Total:        int64(resp.Total),
	}, nil
}
