package domain

import "context"

// MailSenderService - provides access to a business logic
type MailSenderService interface {
	// SendMail - sends mail
	//
	SendMail(
		ctx context.Context,
		receivers []string,
		content interface{},
		template string,
	) error
}
