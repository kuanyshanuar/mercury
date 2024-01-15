package service

import (
	"context"
	"time"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"
	"gitlab.com/zharzhanov/mercury/internal/helpers"

	"github.com/go-kit/log"
)

type service struct {
	repository domain.LeadBuilderRepository
}

// NewService - creates a new service.
func NewService(
	repository domain.LeadBuilderRepository,
	logger log.Logger,
) domain.LeadBuilderService {
	var service domain.LeadBuilderService
	{
		service = newBasicService(repository)
		service = loggingServiceMiddleware(logger)(service)
	}
	return service
}

func newBasicService(repository domain.LeadBuilderRepository) domain.LeadBuilderService {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateLeadBuilder(
	ctx context.Context,
	lead *domain.LeadBuilder,
	_ domain.CallerID,
) error {

	// Validate input
	//
	if err := s.validateCreateLeadBuilder(ctx, lead); err != nil {
		return err
	}

	// Check if lead exists between date issued time and expiration time
	//
	isExist, err := s.repository.IsLeadExistByDateRange(
		ctx,
		lead.BuilderID,
		lead.IssuedAt,
		lead.ExpiresAt,
	)
	if err != nil {
		return err
	}
	if isExist {
		return errors.NewErrInvalidArgument("lead exists between these dates")
	}

	// Assign status
	//
	{
		lead.StatusID = domain.StatusActive
	}

	return s.repository.CreateLeadBuilder(ctx, lead)
}

func (s *service) UpdateLeadBuilder(
	ctx context.Context,
	leadID domain.LeadID,
	lead *domain.LeadBuilder,
	callerID domain.CallerID,
) error {

	// Validate input
	//
	if leadID <= 0 {
		return errors.NewErrInvalidArgument("id reuired")
	}
	if err := s.validateUpdateLeadBuilder(ctx, lead); err != nil {
		return err
	}

	// Check if other lead exists
	//
	isExist, err := s.repository.IsOtherLeadExist(
		ctx,
		lead.ID,
		lead.BuilderID,
		lead.IssuedAt,
		lead.ExpiresAt,
	)
	if err != nil {
		return err
	}
	if isExist {
		return errors.NewErrInvalidArgument("lead exists between these dates")
	}

	return s.repository.UpdateLeadBuilder(ctx, leadID, lead)
}

func (s *service) GetLeadBuilder(
	ctx context.Context,
	leadID domain.LeadID,
	callerID domain.CallerID,
) (*domain.LeadBuilder, error) {
	// Validate input
	//
	if leadID <= 0 {
		return nil, errors.NewErrInvalidArgument("id reuired")
	}

	return s.repository.GetLeadBuilder(ctx, leadID)
}

func (s *service) ListLeadBuilders(
	ctx context.Context,
	criteria domain.LeadBuilderSearchCriteria,
	callerID domain.CallerID,
) ([]*domain.LeadBuilder, domain.Total, error) {
	return s.repository.ListLeadBuilders(ctx, criteria)
}

func (s *service) DeleteLeadBuilder(
	ctx context.Context,
	leadID domain.LeadID,
	callerID domain.CallerID,
) error {

	// Validate input
	//
	if leadID <= 0 {
		return errors.NewErrInvalidArgument("id reuired")
	}

	return s.repository.DeleteLeadBuilder(ctx, leadID)
}

func (s *service) RevokeLeadBuilder(
	ctx context.Context,
	leadID domain.LeadID,
	callerID domain.CallerID,
) error {
	// Validate inputs
	//
	if leadID <= 0 {
		return errors.NewErrInvalidArgument("lead id required")
	}

	return s.repository.RevokeLeadBuilder(ctx, leadID)
}

func (s *service) validateCreateLeadBuilder(ctx context.Context, lead *domain.LeadBuilder) error {
	if lead == nil {
		return errors.NewErrInvalidArgument("lead required")
	}
	if lead.BuilderID <= 0 {
		return errors.NewErrInvalidArgument("builder id required")
	}
	if lead.IssuedAt <= 0 {
		return errors.NewErrInvalidArgument("date start required")
	}
	if lead.ExpiresAt <= 0 {
		return errors.NewErrInvalidArgument("date end required")
	}
	if lead.ExpiresAt < lead.IssuedAt {
		return errors.NewErrInvalidArgument("date end cannot be lower than date start")
	}
	if lead.IssuedAt < helpers.StartOfTheDay(time.Now().Unix()) {
		return errors.NewErrInvalidArgument("invalid date start")
	}
	if lead.ExpiresAt < helpers.StartOfTheDay(time.Now().Unix()) {
		return errors.NewErrInvalidArgument("invalid date end")
	}

	return nil
}

func (s *service) validateUpdateLeadBuilder(ctx context.Context, lead *domain.LeadBuilder) error {
	if lead == nil {
		return errors.NewErrInvalidArgument("lead required")
	}
	if lead.BuilderID <= 0 {
		return errors.NewErrInvalidArgument("builder id required")
	}
	if lead.IssuedAt <= 0 {
		return errors.NewErrInvalidArgument("date start required")
	}
	if lead.ExpiresAt <= 0 {
		return errors.NewErrInvalidArgument("date end required")
	}
	if lead.ExpiresAt < lead.IssuedAt {
		return errors.NewErrInvalidArgument("date end cannot be lower than date start")
	}
	if lead.ExpiresAt < helpers.StartOfTheDay(time.Now().Unix()) {
		return errors.NewErrInvalidArgument("invalid date end")
	}

	return nil
}
