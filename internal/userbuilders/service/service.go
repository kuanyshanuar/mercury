package service

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"

	"github.com/go-kit/log"
)

type service struct {
	repository domain.UserBuilderRepository
}

// NewService - creates a new service
func NewService(
	repository domain.UserBuilderRepository,
	logger log.Logger,
) domain.UserBuilderService {
	var service domain.UserBuilderService
	{
		service = newBasicService(repository)
		service = loggingServiceMiddleware(logger)(service)
	}
	return service
}

func newBasicService(
	repository domain.UserBuilderRepository,
) domain.UserBuilderService {
	return &service{
		repository: repository,
	}
}

func (s *service) Subscribe(
	ctx context.Context,
	subscriberID int64,
	builderID int64,
	_ domain.CallerID,
) (err error) {

	// Validate inputs
	//
	err = s.validateInternal(subscriberID, builderID)
	if err != nil {
		return err
	}

	return s.repository.Subscribe(ctx, subscriberID, builderID)
}

func (s *service) Unsubscribe(
	ctx context.Context,
	subscriberID int64,
	builderID int64,
	_ domain.CallerID,
) (err error) {

	// Validate inputs
	//
	err = s.validateInternal(subscriberID, builderID)
	if err != nil {
		return err
	}

	return s.repository.Unsubscribe(ctx, subscriberID, builderID)
}

func (s *service) ListBuilders(
	ctx context.Context,
	subscriberID int64,
	criteria domain.UserBuilderSearchCriteria,
	_ domain.CallerID,
) (builders []*domain.Builder, total domain.Total, err error) {

	// Validate inputs
	//
	if subscriberID <= 0 {
		return nil, 0, errors.NewErrInvalidArgument("subscriber id required")
	}

	return s.repository.ListBuilders(ctx, subscriberID, criteria)
}

func (s *service) validateInternal(subscriberID int64, builderID int64) error {
	if subscriberID <= 0 {
		return errors.NewErrInvalidArgument("subscriber id required")
	}
	if builderID <= 0 {
		return errors.NewErrInvalidArgument("builder id required")
	}

	return nil
}
