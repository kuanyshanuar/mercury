package service

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"

	"github.com/go-kit/log"
)

type service struct {
	repository domain.BuilderRepository
}

// NewService - creates a new service.
func NewService(
	repository domain.BuilderRepository,
	logger log.Logger,
) domain.BuilderService {
	var service domain.BuilderService
	{
		service = newBasicService(repository)
		service = loggingServiceMiddleware(logger)(service)
	}
	return service
}

func newBasicService(
	repository domain.BuilderRepository,
) domain.BuilderService {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateBuilder(
	ctx context.Context,
	builder *domain.Builder,
	_ domain.CallerID,
) (domain.BuilderID, error) {

	// Assign role
	//
	{
		builder.RoleID = domain.RoleBuilder
		builder.IsVerified = true
	}

	// Validate inputs
	//
	if err := s.validateCreateBuilder(builder); err != nil {
		return 0, err
	}

	// Create builder
	//
	return s.repository.Create(ctx, builder)
}

func (s *service) UpdateBuilder(
	ctx context.Context,
	builderID domain.BuilderID,
	builder *domain.Builder,
	_ domain.CallerID,
) error {

	// Validate inputs
	//
	if builderID <= 0 {
		return errors.NewErrInvalidArgument("builder id required")
	}
	if builder == nil {
		return errors.NewErrInvalidArgument("builder required")
	}

	return s.repository.Update(ctx, builderID, builder)
}

func (s *service) DeleteBuilder(
	ctx context.Context,
	builderID domain.BuilderID,
	_ domain.CallerID,
) error {

	// Validate inputs
	//
	if builderID <= 0 {
		return errors.NewErrInvalidArgument("builder id required")
	}

	return s.repository.Delete(ctx, builderID)
}

func (s *service) GetBuilder(
	ctx context.Context,
	builderID domain.BuilderID,
	callerID domain.CallerID,
) (*domain.Builder, error) {

	// Validate inputs
	//
	if builderID <= 0 {
		return nil, errors.NewErrInvalidArgument("builder id required")
	}

	// Get builder
	//
	builder, err := s.repository.Get(ctx, builderID)
	if err != nil {
		return nil, err
	}

	// Get is favourite
	//
	if callerID.UserID > 0 {
		isFavourite, err := s.repository.IsFavouriteBuilder(ctx, builderID, domain.UserID(callerID.UserID))
		if err != nil {
			return nil, err
		}
		builder.IsFavourite = isFavourite
	}

	return builder, nil
}

func (s *service) ListBuilders(
	ctx context.Context,
	criteria domain.BuilderSearchCriteria,
	_ domain.CallerID,
) ([]*domain.Builder, domain.Total, error) {
	return s.repository.List(ctx, criteria)
}

func (s *service) validateCreateBuilder(builder *domain.Builder) error {

	if builder == nil {
		return errors.NewErrInvalidArgument("builder required")
	}
	if builder.RoleID <= 0 {
		return errors.NewErrInvalidArgument("role id required")
	}
	if builder.RoleID != domain.RoleBuilder {
		return errors.NewErrInvalidArgument("role id does not match")
	}
	if len(builder.FirstName) == 0 {
		return errors.NewErrInvalidArgument("first name required")
	}
	if len(builder.LastName) == 0 {
		return errors.NewErrInvalidArgument("last name required")
	}
	if len(builder.Email) == 0 {
		return errors.NewErrInvalidArgument("email required")
	}
	if len(builder.Phone) == 0 {
		return errors.NewErrInvalidArgument("phone required")
	}
	if len(builder.ConsultationPhoneNumber) == 0 {
		return errors.NewErrInvalidArgument("consultation phone number required")
	}
	if len(builder.Password) == 0 {
		return errors.NewErrInvalidArgument("password required")
	}
	if len(builder.City) == 0 {
		return errors.NewErrInvalidArgument("city required")
	}

	return nil
}
