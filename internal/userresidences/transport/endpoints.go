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
	AddResidenceToFavouritesEndpoint      endpoint.Endpoint
	DeleteResidenceFromFavouritesEndpoint endpoint.Endpoint
	ListFavouriteResidencesEndpoint       endpoint.Endpoint
}

// NewEndpoint returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func NewEndpoint(service domain.UserResidenceService, serviceSecretKey domain.ServiceSecretKey, logger log.Logger) Endpoints {
	factory := func(creator func(domain.UserResidenceService) endpoint.Endpoint, logKey string) endpoint.Endpoint {
		return helpers.SetupEndpoint(creator(service), serviceSecretKey, logger, "UserResidences", logKey)
	}

	return Endpoints{
		AddResidenceToFavouritesEndpoint:      factory(MakeAddResidenceToFavouriteEndpoint, "AddResidenceToFavourite"),
		DeleteResidenceFromFavouritesEndpoint: factory(MakeDeleteResidenceFromFavouritesEndpoint, "DeleteResidenceFromFavourites"),
		ListFavouriteResidencesEndpoint:       factory(MakeListFavouriteResidencesEndpoint, "ListFavouriteResidences"),
	}
}

// MakeAddResidenceToFavouriteEndpoint - impl.
func MakeAddResidenceToFavouriteEndpoint(service domain.UserResidenceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(addResidenceToFavouritesRequest)

		err = service.AddResidenceToFavourites(
			ctx,
			req.UserID,
			req.ResidenceID,
			helpers.CallerID(ctx),
		)
		return addResidenceToFavouritesResponse{
			Err: err,
		}, nil
	}
}

type addResidenceToFavouritesRequest struct {
	UserID      domain.UserID
	ResidenceID domain.ResidenceID
}

type addResidenceToFavouritesResponse struct {
	Err error
}

// MakeDeleteResidenceFromFavouritesEndpoint - Impl.
func MakeDeleteResidenceFromFavouritesEndpoint(service domain.UserResidenceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(deleteResidenceFromFavouritesRequest)

		err = service.DeleteResidenceFromFavourites(
			ctx,
			req.UserID,
			req.ResidenceID,
			helpers.CallerID(ctx),
		)
		return deleteResidenceFromFavouritesResponse{
			Err: err,
		}, nil
	}
}

type deleteResidenceFromFavouritesRequest struct {
	UserID      domain.UserID
	ResidenceID domain.ResidenceID
}

type deleteResidenceFromFavouritesResponse struct {
	Err error
}

// MakeListFavouriteResidencesEndpoint Impl.
func MakeListFavouriteResidencesEndpoint(service domain.UserResidenceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(listFavouriteResidencesRequest)

		resp, total, err := service.ListFavouriteResidences(ctx, req.UserID, req.Criteria, helpers.CallerID(ctx))
		return listFavouriteResidencesResponse{
			ResidenceIDs: resp,
			Total:        total,
			Err:          err,
		}, nil
	}
}

type listFavouriteResidencesRequest struct {
	UserID   domain.UserID
	Criteria domain.FavouriteResidencesSearchCriteria
}

type listFavouriteResidencesResponse struct {
	ResidenceIDs []int64
	Total        domain.Total
	Err          error
}
