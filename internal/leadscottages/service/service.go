package service

import (
	"context"
	"github.com/go-kit/log"
	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"
	"gitlab.com/zharzhanov/mercury/internal/helpers"
	"time"
)

type service struct {
	repository domain.LeadCottageRepository
}

// NewService - create a new service
func NewService(
	repository domain.LeadCottageRepository,
	logger log.Logger,
) domain.LeadCottageService {
	var service domain.LeadCottageService
	{
		service = newBasicService(repository)
		service = loggingServiceMiddleware(logger)(service)
	}
	return service
}

func newBasicService(
	repository domain.LeadCottageRepository,
) domain.LeadCottageService {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateLeadCottage(
	ctx context.Context,
	lead *domain.LeadCottage,
	caller domain.CallerID,
) (int64, error) {

	// Validate inputs
	//
	err := s.validateCreateLeadCottage(lead)
	if err != nil {
		return 0, err
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
		lead.CottageID,
		lead.IssuedAt,
		lead.ExpiresAt,
	)
	if err != nil {
		return 0, err
	}
	if isExist {
		return 0, errors.NewErrInvalidArgument("lead exists between these dates")
	}

	return s.repository.CreateLeadCottage(ctx, lead)
}

func (s *service) UpdateLeadCottage(
	ctx context.Context,
	leadID int64,
	lead *domain.LeadCottage,
	caller domain.CallerID,
) (*domain.LeadCottage, error) {
	// Validate inputs
	//
	err := s.validateUpdateLeadCottage(lead)
	if err != nil {
		return nil, err
	}

	// Assign id
	//
	{
		lead.ID = leadID
	}

	// Check if other lead exists
	//
	isExist, err := s.repository.IsOtherLeadExist(
		ctx,
		domain.LeadID(lead.ID),
		lead.CottageID,
		lead.IssuedAt,
		lead.ExpiresAt,
	)
	if err != nil {
		return nil, err
	}
	if isExist {
		return nil, errors.NewErrInvalidArgument("lead exists between these dates")
	}

	return s.repository.UpdateLeadCottage(ctx, leadID, lead)
}

func (s *service) DeleteLeadCottage(
	ctx context.Context,
	ID int64,
	caller domain.CallerID,
) error {
	if ID <= 0 {
		return errors.NewErrInvalidArgument("lead cottage id is invalid")
	}
	return s.repository.DeleteLeadCottage(ctx, ID)
}

func (s *service) GetLeadCottage(
	ctx context.Context,
	ID int64,
	caller domain.CallerID,
) (*domain.LeadCottage, error) {
	if ID <= 0 {
		return nil, errors.NewErrInvalidArgument("lead cottage id is invalid")
	}
	return s.repository.GetLeadCottage(ctx, ID)

}

func (s *service) ListLeadCottage(
	ctx context.Context,
	criteria domain.LeadCottageSearchCriteria,
	caller domain.CallerID,
) ([]*domain.LeadCottage, domain.Total, error) {
	return s.repository.ListLeadCottage(ctx, criteria)
}

func (s *service) RevokeLeadCottage(
	ctx context.Context,
	leadID domain.LeadID,
	callerID domain.CallerID,
) error {
	if leadID <= 0 {
		return errors.NewErrInvalidArgument("lead ID is needed")
	}

	return s.repository.RevokeLeadCottage(ctx, leadID)
}

func (s *service) validateCreateLeadCottage(lead *domain.LeadCottage) error {
	if lead == nil {
		return errors.NewErrInvalidArgument("cottage lead is nil")
	}
	if lead.CottageID <= 0 {
		return errors.NewErrInvalidArgument("cottage lead  id is invalid")
	}
	if lead.IssuedAt <= 0 {
		return errors.NewErrInvalidArgument("date start required")
	}
	if lead.ExpiresAt <= 0 {
		return errors.NewErrInvalidArgument("date end required")
	}
	if lead.IssuedAt > lead.ExpiresAt {
		return errors.NewErrInvalidArgument("cottage lead  issue and expire date invalid")
	}
	if lead.IssuedAt < helpers.StartOfTheDay(time.Now().Unix()) {
		return errors.NewErrInvalidArgument("cottage lead issue date invalid")
	}
	if lead.ExpiresAt <= helpers.EndOfTheDay(time.Now().Unix()) {
		return errors.NewErrInvalidArgument("cottage lead expire date invalid")
	}
	return nil
}
func (s *service) validateUpdateLeadCottage(lead *domain.LeadCottage) error {
	if lead == nil {
		return errors.NewErrInvalidArgument("cottage lead is nil")
	}
	if lead.CottageID <= 0 {
		return errors.NewErrInvalidArgument("cottage lead  id is required")
	}
	if lead.IssuedAt <= 0 {
		return errors.NewErrInvalidArgument("cottage lead issue date invalid")
	}
	if lead.ExpiresAt < 0 {
		return errors.NewErrInvalidArgument("cottage lead expire date is required")
	}
	if lead.IssuedAt > lead.ExpiresAt {
		return errors.NewErrInvalidArgument("cottage lead  issue and expire date invalid")
	}
	if lead.ExpiresAt < helpers.StartOfTheDay(time.Now().Unix()) {
		return errors.NewErrInvalidArgument("cottage lead expire date invalid")
	}

	return nil
}
