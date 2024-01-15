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
func NewRepository(db *gorm.DB) domain.ContactDetailsRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateContactDetails(ctx context.Context, contactDetails *domain.ContactDetails) error {
	err := r.db.
		WithContext(ctx).
		Table(domain.ContactDetailsTableName).
		Create(&contactDetails).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) ListContactDetails(ctx context.Context) ([]*domain.ContactDetails, error) {
	var list []*domain.ContactDetails
	err := r.db.
		WithContext(ctx).
		Table(domain.ContactDetailsTableName).
		Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *repository) CreateResidenceContactDetails(ctx context.Context, contactDetails *domain.ResidenceContactDetails) error {
	err := r.db.
		WithContext(ctx).
		Table(domain.ResidenceContactDetailsTableName).
		Create(&contactDetails).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) ListResidenceContactDetails(
	ctx context.Context,
	criteria domain.ResidenceContactDetailsSearchCriteria,
) ([]*domain.ResidenceContactDetails, domain.Total, error) {
	var (
		db         = r.db
		list       []*domain.ResidenceContactDetails
		totalCount int64
	)

	if criteria.IsDelivered == 1 {
		db = db.Where("is_delivered = true")
	}
	if criteria.IsDelivered == 2 {
		db = db.Where("is_delivered = false")
	}
	if criteria.ResidenceID > 0 {
		db = db.Where("residence_id = ?", criteria.ResidenceID)
	}
	if len(criteria.Phone) > 0 {
		db = db.Where("phone = ?", criteria.Phone)
	}
	if criteria.FromTime > 0 {
		db = db.Where("created_at >= ?", criteria.FromTime)
	}
	if criteria.ToTime > 0 {
		db = db.Where("created_at <= ?", criteria.ToTime)
	}

	err := db.
		Table(domain.ResidenceContactDetailsTableName).
		Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.
		Debug().
		WithContext(ctx).
		Table(domain.ResidenceContactDetailsTableName).
		Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	return list, domain.Total(totalCount), nil
}

func (r *repository) MarkAsDelivered(
	ctx context.Context,
	contactID int64,
) error {
	err := r.db.
		WithContext(ctx).
		Table(domain.ResidenceContactDetailsTableName).
		Where("id = ?", contactID).
		Update("is_delivered", true).Error
	if err != nil {
		return err
	}

	return nil
}
