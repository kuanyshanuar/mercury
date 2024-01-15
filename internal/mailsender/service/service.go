package service

import (
	"bytes"
	"context"
	"fmt"
	"net/smtp"
	"text/template"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"

	"github.com/go-kit/log"
)

type service struct {
	auth  smtp.Auth
	host  string
	email string
}

// NewService - creates a new service.
func NewService(
	client smtp.Auth,
	host string,
	email string,
	logger log.Logger,
) domain.MailSenderService {
	var service domain.MailSenderService
	{
		service = newBasicService(
			client,
			host,
			email,
		)
		service = loggingServiceMiddleware(logger)(service)
	}
	return service
}

func newBasicService(
	auth smtp.Auth,
	host string,
	email string,
) domain.MailSenderService {
	return &service{
		auth:  auth,
		host:  host,
		email: email,
	}
}

func (s *service) SendMail(
	_ context.Context,
	receivers []string,
	content interface{},
	template string,
) error {

	// Validate inputs
	//
	if len(receivers) == 0 {
		return errors.NewErrInvalidArgument("email required")
	}

	// Get template
	//
	mailTemplate, err := s.assertGetTemplate(template)
	if err != nil {
		return err
	}

	// Get mail content
	//
	mailContent, err := s.asserGetMailContent(mailTemplate, content)
	if err != nil {
		return err
	}

	// Send email
	//
	err = smtp.SendMail(s.host, s.auth, s.email, receivers, mailContent)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) assertGetTemplate(templateName string) (*template.Template, error) {

	// Parse file
	//
	mailTemplate, err := template.ParseFiles(templateName)
	if err != nil {
		return nil, err
	}

	return mailTemplate, nil
}

func (s *service) asserGetMailContent(template *template.Template, content interface{}) ([]byte, error) {

	var (
		body        bytes.Buffer
		mimeHeaders = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	)

	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	err := template.Execute(&body, content)
	if err != nil {
		return nil, err
	}

	return body.Bytes(), nil
}
