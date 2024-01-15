package service

import (
	"context"
	"github.com/go-kit/log"
	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"
	"time"
)

type service struct {
	mailService domain.MailService
	repository  domain.ContactDetailsRepository
}

// NewService - creates a new service
func NewService(
	mailService domain.MailService,
	repository domain.ContactDetailsRepository,
	logger log.Logger,
) domain.ContactDetailsService {
	var service domain.ContactDetailsService
	{
		service = newBasicService(
			mailService,
			repository,
		)
		service = loggingServiceMiddleware(logger)(service)
	}
	return service
}

func newBasicService(
	mailService domain.MailService,
	repository domain.ContactDetailsRepository,
) domain.ContactDetailsService {
	return &service{
		mailService: mailService,
		repository:  repository,
	}
}

func (s *service) CreateContactDetails(
	ctx context.Context,
	contactDetails *domain.ContactDetails,
	_ domain.CallerID,
) error {
	// Validate inputs
	//
	err := s.validateInternalContactDetails(contactDetails)
	if err != nil {
		return err
	}

	// Send mail
	//
	err = s.mailService.SendContactDetailEmail(
		ctx,
		"shtabkvartir@gmail.com",
		domain.ContactDetailContent{
			FullName: contactDetails.FullName,
			Phone:    contactDetails.Phone,
			Message:  contactDetails.Message,
		},
	)
	if err != nil {
		return err
	}

	return s.repository.CreateContactDetails(ctx, contactDetails)
}

func (s *service) CreateResidenceContactDetails(
	ctx context.Context,
	contactDetails *domain.ResidenceContactDetails,
	callerID domain.CallerID,
) error {

	// Validate inputs
	//
	err := s.validateInternalResidenceContactDetails(contactDetails)
	if err != nil {
		return err
	}

	// Check if the contact by the provided residence id exists past day
	if contactDetails.ResidenceID > 0 {

		var (
			fromTime = time.Now().Truncate(24 * time.Hour)
			toTime   = fromTime.Add(24 * time.Hour)
		)

		contacts, _, err := s.repository.ListResidenceContactDetails(ctx, domain.ResidenceContactDetailsSearchCriteria{
			Page: domain.PageRequest{
				Offset: 0,
				Size:   10,
			},
			ResidenceID: contactDetails.ResidenceID,
			Phone:       contactDetails.Phone,
			FromTime:    fromTime.Unix(),
			ToTime:      toTime.Unix(),
		})
		if err != nil {
			return err
		}
		if len(contacts) > 0 {
			return nil
		}
	}

	return s.repository.CreateResidenceContactDetails(ctx, contactDetails)
}

func (s *service) ListContactDetails(
	ctx context.Context,
	_ domain.CallerID,
) ([]*domain.ContactDetails, error) {
	return s.repository.ListContactDetails(ctx)
}

func (s *service) ListResidenceContactDetails(
	ctx context.Context,
	criteria domain.ResidenceContactDetailsSearchCriteria,
	_ domain.CallerID,
) ([]*domain.ResidenceContactDetails, domain.Total, error) {
	return s.repository.ListResidenceContactDetails(ctx, criteria)
}

func (s *service) MarkAsDelivered(
	ctx context.Context,
	contactID int64,
	callerID domain.CallerID,
) error {
	if contactID <= 0 {
		return errors.NewErrInvalidArgument("id required")
	}

	return s.repository.MarkAsDelivered(ctx, contactID)
}

func (s *service) validateInternalContactDetails(details *domain.ContactDetails) error {
	if details == nil {
		return errors.NewErrInvalidArgument("contact details required")
	}
	if len(details.FullName) == 0 {
		return errors.NewErrInvalidArgument("full name required")
	}
	if len(details.Phone) == 0 {
		return errors.NewErrInvalidArgument("phone required")
	}

	return nil
}

func (s *service) validateInternalResidenceContactDetails(details *domain.ResidenceContactDetails) error {
	if details == nil {
		return errors.NewErrInvalidArgument("contact details required")
	}
	if len(details.FullName) == 0 {
		return errors.NewErrInvalidArgument("full name required")
	}
	if len(details.Phone) == 0 {
		return errors.NewErrInvalidArgument("phone required")
	}

	return nil
}
