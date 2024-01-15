package mail

import (
	"net/smtp"

	"gitlab.com/zharzhanov/mercury/internal/domain"
)

// NewSMTPAuth - creates an authentication
func NewSMTPAuth(cfg domain.EmailConfig) smtp.Auth {
	auth := smtp.PlainAuth(
		"",
		cfg.Email,
		cfg.Password,
		cfg.Host)

	return auth
}
