package service

import (
	"context"
	"fmt"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"

	"github.com/go-kit/log"
)

type service struct {
	residenceService domain.ResidencesService
	repository       domain.UserResidenceRepository
}

// NewService - creates a new service with necessary dependencies.
func NewService(
	residenceService domain.ResidencesService,
	repository domain.UserResidenceRepository,
	logger log.Logger,
) domain.UserResidenceService {
	var service domain.UserResidenceService
	{
		service = newBasicService(
			residenceService,
			repository,
		)
		service = loggingServiceMiddleware(logger)(service)
	}
	return service
}

func newBasicService(
	residenceService domain.ResidencesService,
	repository domain.UserResidenceRepository,
) domain.UserResidenceService {
	return &service{
		residenceService: residenceService,
		repository:       repository,
	}
}

func (s *service) AddResidenceToFavourites(
	ctx context.Context,
	userID domain.UserID,
	residenceID domain.ResidenceID,
	callerID domain.CallerID,
) error {

	// Validate input.
	//
	if err := s.validateInputInternal(
		ctx,
		userID,
		residenceID,
		callerID,
	); err != nil {
		return err
	}

	return s.repository.AddResidenceToFavourites(ctx, userID, residenceID)
}

func (s *service) DeleteResidenceFromFavourites(
	ctx context.Context,
	userID domain.UserID,
	residenceID domain.ResidenceID,
	callerID domain.CallerID,
) error {

	// Validate user id
	//
	if userID <= 0 {
		return errors.NewErrInvalidArgument("invalid user id")
	}

	// Validate residence id
	//
	if residenceID <= 0 {
		return errors.NewErrInvalidArgument("invalid residence id")
	}

	return s.repository.DeleteResidenceFromFavourites(ctx, userID, residenceID)
}

func (s *service) ListFavouriteResidences(
	ctx context.Context,
	userID domain.UserID,
	criteria domain.FavouriteResidencesSearchCriteria,
	_ domain.CallerID,
) ([]int64, domain.Total, error) {

	// Validate user id
	//
	if userID <= 0 {
		return nil, 0, errors.NewErrInvalidArgument("user id required")
	}

	return s.repository.ListFavouriteResidences(ctx, userID, criteria)
}

func (s *service) validateInputInternal(
	ctx context.Context,
	userID domain.UserID,
	residenceID domain.ResidenceID,
	callerID domain.CallerID,
) error {

	// Validate user id.
	//
	if userID < 0 {
		return errors.NewErrInvalidArgument("user id required")
	}

	// Validate residence id.
	//
	if residenceID < 0 {
		return errors.NewErrInvalidArgument("residence id required")
	}
	exist, err := s.residenceService.IsResidenceExist(ctx, residenceID, callerID)
	if err != nil {
		return err
	}
	if !exist {
		return errors.NewErrNotFound(
			fmt.Sprintf("residence not found by id: %v", residenceID),
		)
	}

	return nil
}
