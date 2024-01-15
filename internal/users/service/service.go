package service

import (
	"context"
	errors "gitlab.com/zharzhanov/mercury/internal/error"

	"gitlab.com/zharzhanov/mercury/internal/domain"

	"github.com/go-kit/log"
)

type service struct {
	repository domain.UserRepository
}

// NewService - creates a new service
func NewService(
	repository domain.UserRepository,
	logger log.Logger,
) domain.UserService {
	var service domain.UserService
	{
		service = newBasicService(repository)
		service = loggingServiceMiddleware(logger)(service)
	}
	return service
}

func newBasicService(
	repository domain.UserRepository,
) domain.UserService {
	return &service{
		repository: repository,
	}
}

func (s *service) ListUsers(
	ctx context.Context,
	criteria domain.UserSearchCriteria,
	_ domain.CallerID,
) ([]*domain.User, domain.Total, error) {
	return s.repository.ListUsers(ctx, criteria)
}

func (s *service) UpdateUser(
	ctx context.Context,
	userID domain.UserID,
	user *domain.User,
	_ domain.CallerID,
) error {

	// Validate inputs
	//
	if userID <= 0 {
		return errors.NewErrInvalidArgument("id required")
	}
	if user == nil {
		return errors.NewErrInvalidArgument("user required")
	}

	return s.repository.UpdateUser(ctx, userID, user)
}
