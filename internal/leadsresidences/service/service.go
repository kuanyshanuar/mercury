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
	repository domain.LeadResidenceRepository
}

// NewService - create a new service
func NewService(
	repository domain.LeadResidenceRepository,
	logger log.Logger,
) domain.LeadResidenceService {
	var service domain.LeadResidenceService
	{
		service = newBasicService(repository)
		service = loggingServiceMiddleware(logger)(service)
	}
	return service
}

func newBasicService(
	repository domain.LeadResidenceRepository,
) domain.LeadResidenceService {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateLeadResidence(
	ctx context.Context,
	lead *domain.LeadResidence,
	_ domain.CallerID,
) error {

	// Validate inputs
	//
	err := s.validateCreateLeadResidence(ctx, lead)
	if err != nil {
		return err
	}

	// Assign status
	//
	{
		lead.StatusID = domain.StatusActive
	}

	// Check if lead exists between date issued time and expiration time
	//
	isExist, err := s.repository.IsLeadExistByDate(
		ctx,
		lead.ResidenceID,
		lead.IssuedAt,
		lead.ExpiresAt,
	)
	if err != nil {
		return err
	}
	if isExist {
		return errors.NewErrInvalidArgument("lead exists between these dates")
	}

	return s.repository.CreateLeadResidence(ctx, lead)
}

func (s *service) ListLeadResidences(
	ctx context.Context,
	criteria domain.LeadResidenceSearchCriteria,
	_ domain.CallerID,
) ([]*domain.LeadResidence, domain.Total, error) {
	leads, total, err := s.repository.ListLeadResidences(ctx, criteria)
	if err != nil {
		return nil, 0, err
	}

	return leads, total, nil
}

func (s *service) UpdateLeadResidence(
	ctx context.Context,
	leadID domain.LeadID,
	lead *domain.LeadResidence,
	_ domain.CallerID,
) error {

	// Validate inputs
	//
	if leadID <= 0 {
		return errors.NewErrInvalidArgument("id reuired")
	}
	err := s.validateUpdateLeadResidence(ctx, lead)
	if err != nil {
		return err
	}

	// Assign id
	//
	{
		lead.ID = leadID
	}

	// Check if other lead exists with same residence id
	//
	isExist, err := s.repository.IsOtherLeadExist(
		ctx,
		lead.ID,
		lead.ResidenceID,
		lead.IssuedAt,
		lead.ExpiresAt,
	)
	if err != nil {
		return err
	}
	if isExist {
		return errors.NewErrInvalidArgument("lead exists between these dates")
	}

	return s.repository.UpdateLeadResidence(ctx, lead)
}

func (s *service) DeleteLeadResidence(
	ctx context.Context,
	leadID domain.LeadID,
	_ domain.CallerID,
) error {

	// Validate inputs
	//
	if leadID <= 0 {
		return errors.NewErrInvalidArgument("lead id required")
	}

	// Delete data from storage
	//
	err := s.repository.DeleteLeadResidence(ctx, leadID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) RevokeLeadResidence(
	ctx context.Context,
	leadID domain.LeadID,
	_ domain.CallerID,
) error {

	// Validate inputs
	//
	if leadID <= 0 {
		return errors.NewErrInvalidArgument("lead id required")
	}

	return s.repository.RevokeLeadResidence(ctx, leadID)
}

func (s *service) GetLeadResidence(
	ctx context.Context,
	leadID domain.LeadID,
	callerID domain.CallerID,
) (*domain.LeadResidence, error) {
	// Validate input
	//
	if leadID <= 0 {
		return nil, errors.NewErrInvalidArgument("lead id required")
	}

	return s.repository.GetLeadResidence(ctx, leadID)
}

func (s *service) validateCreateLeadResidence(ctx context.Context, lead *domain.LeadResidence) error {
	if lead == nil {
		return errors.NewErrInvalidArgument("lead required")
	}
	if lead.ResidenceID <= 0 {
		return errors.NewErrInvalidArgument("residence id required")
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

func (s *service) validateUpdateLeadResidence(ctx context.Context, lead *domain.LeadResidence) error {
	if lead == nil {
		return errors.NewErrInvalidArgument("lead required")
	}
	if lead.ResidenceID <= 0 {
		return errors.NewErrInvalidArgument("residence id required")
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
