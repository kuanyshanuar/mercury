package service

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"

	"github.com/go-kit/log"
)

type service struct {
	senderService domain.MailSenderService
	emailTemplate domain.EmailTemplate
}

// NewService - creates a new service.
func NewService(
	senderService domain.MailSenderService,
	emailTemplate domain.EmailTemplate,
	logger log.Logger,
) domain.MailService {
	var service domain.MailService
	{
		service = newBasicService(
			senderService,
			emailTemplate,
		)
		service = loggingServiceMiddleware(logger)(service)
	}
	return service
}

func newBasicService(
	senderService domain.MailSenderService,
	emailTemplate domain.EmailTemplate,
) domain.MailService {
	return &service{
		senderService: senderService,
		emailTemplate: emailTemplate,
	}
}

func (s *service) SendResetPasswordEmail(
	ctx context.Context,
	email string,
	content domain.ResetMailPasswordEmailContent,
) error {

	// Validate inputs
	//
	if len(email) == 0 {
		return errors.NewErrInvalidArgument("email required")
	}

	// Sends email
	//
	err := s.senderService.SendMail(
		ctx,
		[]string{
			email,
		},
		content,
		s.emailTemplate.ResetPasswordTemplate,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) SendResidenceContactDetailEmail(
	ctx context.Context,
	email string,
	content domain.ResidenceContactDetailContent,
) error {

	// Validate inputs
	//
	if len(email) == 0 {
		return errors.NewErrInvalidArgument("email required")
	}

	// Sends email
	//
	err := s.senderService.SendMail(
		ctx,
		[]string{
			email,
		},
		content,
		s.emailTemplate.ResidenceContactDetailTemplate,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) SendContactDetailEmail(
	ctx context.Context,
	email string,
	content domain.ContactDetailContent,
) error {

	// Validate inputs
	//
	if len(email) == 0 {
		return errors.NewErrInvalidArgument("email required")
	}

	// Sends email
	//
	err := s.senderService.SendMail(
		ctx,
		[]string{
			email,
		},
		content,
		s.emailTemplate.ContactDetailTemplate,
	)
	if err != nil {
		return err
	}

	return nil
}
