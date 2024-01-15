package transport

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

// Endpoints - collection of endpoints for user cottage service
type Endpoints struct {
	addUserCottage    endpoint.Endpoint
	deleteUserCottage endpoint.Endpoint
	listUserCottage   endpoint.Endpoint
}

// NewEndpoints - returns endpoints for user cottages
func NewEndpoints(service domain.UserCottageService, serviceSecretKey domain.ServiceSecretKey, logger log.Logger) Endpoints {

	factory := func(creator func(domain.UserCottageService) endpoint.Endpoint, logKey string) endpoint.Endpoint {
		return helpers.SetupEndpoint(creator(service), serviceSecretKey, logger, "UserCottage", logKey)
	}

	return Endpoints{
		addUserCottage:    factory(MakeAddUserCottageEndpoint, "AddUserCottage"),
		deleteUserCottage: factory(MakeDeleteUserCottageEndpoint, "DeleteUserCottage"),
		listUserCottage:   factory(MakeListUserCottageEndpoint, "ListUserCottage"),
	}

}

type addUserCottageRequest struct {
	UserID    int64
	CottageID int64
}

type addUserCottageResponse struct {
	Error error
}

// MakeAddUserCottageEndpoint - impl.
func MakeAddUserCottageEndpoint(service domain.UserCottageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addUserCottageRequest)
		err := service.AddCottageToFavourites(ctx, req.UserID, req.CottageID, helpers.CallerID(ctx))
		if err != nil {
			return addUserCottageResponse{
				Error: err,
			}, nil
		}
		return addUserCottageResponse{
			Error: nil,
		}, nil
	}
}

type deleteUserCottageRequest struct {
	UserID    int64
	CottageID int64
}

type deleteUserCottageResponse struct {
	Error error
}

// MakeDeleteUserCottageEndpoint - impl.
func MakeDeleteUserCottageEndpoint(service domain.UserCottageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteUserCottageRequest)
		err := service.DeleteCottageFromFavourites(ctx, req.UserID, req.CottageID, helpers.CallerID(ctx))
		if err != nil {
			return deleteUserCottageResponse{
				Error: err,
			}, nil
		}
		return deleteUserCottageResponse{
			Error: nil,
		}, nil
	}
}

type listUserCottageRequest struct {
	UserID   int64
	Criteria domain.FavouriteCottagesSearchCriteria
}

type listUserCottageResponse struct {
	Result []int64
	Total  domain.Total
	Error  error
}

// MakeListUserCottageEndpoint - impl.
func MakeListUserCottageEndpoint(service domain.UserCottageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(listUserCottageRequest)
		list, total, err := service.ListFavouriteCottages(ctx, req.UserID, req.Criteria, helpers.CallerID(ctx))
		return listUserCottageResponse{
			Result: list,
			Total:  total,
			Error:  err,
		}, nil
	}
}
