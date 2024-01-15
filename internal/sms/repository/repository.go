package repository

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// NewRepository - creates a new repository
func NewRepository(db *gorm.DB) domain.SmsRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, sms *domain.Sms) error {
	err := r.db.WithContext(ctx).
		Table(domain.SmsMessagesTableName).
		Create(&sms).Error
	if err != nil {
		return err
	}

	return nil
}
