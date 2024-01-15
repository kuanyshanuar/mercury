package service

import (
	"context"
	"fmt"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"

	"github.com/go-kit/log"
)

type service struct {
	repository     domain.UserCottageRepository
	cottageService domain.CottageService
}

// NewService - creates CottageService
func NewService(repository domain.UserCottageRepository, cottageService domain.CottageService, logger log.Logger) domain.UserCottageService {
	var service domain.UserCottageService
	{
		service = newBasicService(repository, cottageService)
		service = loggingServiceMiddleware(logger)(service)
	}

	return service
}

func newBasicService(
	repository domain.UserCottageRepository,
	cottageService domain.CottageService,

) domain.UserCottageService {
	return &service{
		repository:     repository,
		cottageService: cottageService,
	}
}

func (s *service) AddCottageToFavourites(
	ctx context.Context,
	userID int64,
	CottageID int64,
	callerID domain.CallerID,
) error {
	err := s.validateUserCottageInput(ctx, userID, CottageID, callerID)
	if err != nil {
		return err
	}

	return s.repository.AddFavouriteCottage(ctx, userID, CottageID)

}

func (s *service) DeleteCottageFromFavourites(
	ctx context.Context,
	userID int64,
	CottageID int64,
	callerID domain.CallerID,
) error {
	if userID <= 0 {
		return errors.NewErrInvalidArgument("user id is required")
	}
	if CottageID <= 0 {
		return errors.NewErrInvalidArgument("cottage id is required")
	}
	return s.repository.DeleteFavouriteCottage(ctx, userID, CottageID)

}

func (s *service) ListFavouriteCottages(
	ctx context.Context,
	userID int64,
	criteria domain.FavouriteCottagesSearchCriteria,
	callerID domain.CallerID,
) ([]int64, domain.Total, error) {
	if userID <= 0 {
		return nil, 0, errors.NewErrInvalidArgument("user id is required")
	}
	return s.repository.ListFavouriteCottages(ctx, userID, criteria)
}

func (s *service) validateUserCottageInput(
	ctx context.Context,
	userID int64,
	cottageID int64,
	callerID domain.CallerID,
) error {
	// Validate user id.
	//
	if userID < 0 {
		return errors.NewErrInvalidArgument("user id required")
	}

	// Validate cottage id.
	//
	if cottageID < 0 {
		return errors.NewErrInvalidArgument("cottage id required")
	}

	exist, err := s.cottageService.IsCottageExist(ctx, cottageID, callerID)
	if err != nil {
		return err
	}
	if !exist {
		return errors.NewErrNotFound(
			fmt.Sprintf("cottage not found by id: %v", cottageID),
		)
	}

	return nil
}
