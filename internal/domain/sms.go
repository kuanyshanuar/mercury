package domain

import (
	"context"
	"fmt"
)

// SmsMessagesTableName - name of table
const (
	SmsMessagesTableName = "sms_messages"
)

// Sms - represents sms.
type Sms struct {
	ID        int    `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Phone     string `json:"phone" gorm:"column:phone"`
	Message   string `json:"message" gorm:"column:message"`
	ExpiresAt int64  `json:"expires_at" gorm:"column:expires_at"`
	CreatedAt int64  `json:"created_at" gorm:"not null;autoCreateTime;column:created_at"`
}

// GetMessage - returns message string
func (s *Sms) GetMessage() string {
	return fmt.Sprintf("shtab-kvartir.kz: %s", s.Message)
}

// SmsRepository - provides access to storage.
type SmsRepository interface {

	// Create - stores sms in storage.
	//
	Create(ctx context.Context, sms *Sms) error
}

// SmsService - provides access to business logic.
type SmsService interface {

	// SendSms - sends sms to user
	//
	SendSms(ctx context.Context, sms *Sms) error
}
