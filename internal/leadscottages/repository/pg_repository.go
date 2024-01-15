package repository

import (
	"context"
	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// NewRepository exports LeadCottageRepository to be later used for the service layer
func NewRepository(db *gorm.DB) domain.LeadCottageRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateLeadCottage(
	ctx context.Context,
	lead *domain.LeadCottage,
) (int64, error) {
	err := r.db.
		WithContext(ctx).
		Table(domain.LeadCottageTableName).
		Create(&lead).Error
	if err != nil {
		return 0, err
	}

	return lead.ID, nil
}

func (r *repository) UpdateLeadCottage(
	ctx context.Context,
	id int64,
	lead *domain.LeadCottage,
) (*domain.LeadCottage, error) {
	err := r.db.
		WithContext(ctx).
		Table(domain.LeadCottageTableName).
		Where("id = ?", lead.ID).
		Updates(lead).Error
	if err != nil {
		return nil, err
	}

	return lead, nil
}

func (r *repository) DeleteLeadCottage(
	ctx context.Context,
	leadID int64,
) error {
	leadCottage := new(domain.LeadCottage)
	err := r.db.
		WithContext(ctx).
		Table(domain.LeadCottageTableName).
		Where("id = ?", leadID).
		Delete(&leadCottage).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetLeadCottage(
	ctx context.Context,
	leadID int64,
) (*domain.LeadCottage, error) {
	var leadCottage *domain.LeadCottage

	err := r.db.
		WithContext(ctx).
		Table(domain.LeadCottageTableName).
		Preload(domain.LeadCottageAssociation).
		Preload(domain.LeadStatusAssociation).
		Take(&leadCottage, leadID).Error
	if err != nil {
		return nil, err
	}

	return leadCottage, nil
}

func (r *repository) ListLeadCottage(
	ctx context.Context,
	criteria domain.LeadCottageSearchCriteria,
) ([]*domain.LeadCottage, domain.Total, error) {
	var (
		db         = r.db
		lead       []*domain.LeadCottage
		totalCount int64
	)

	if criteria.StatusID > 0 {
		db = db.Where("cottage_leads.status_id = ?", criteria.StatusID)
	}
	if len(criteria.Name) > 0 {
		db = db.Joins(`JOIN cottages AS c ON c.id = cottage_leads.cottage_id AND UPPER(c.title) LIKE UPPER(?)`, "%"+criteria.Name+"%")
	}

	countResult := db.Table(domain.LeadCottageTableName).Where("cottage_leads.deleted_at = 0").Count(&totalCount)
	if countResult.Error != nil {
		return nil, 0, countResult.Error
	}

	err := db.
		WithContext(ctx).
		Scopes(helpers.Paginate(criteria.Page)).
		Table(domain.LeadCottageTableName).
		Preload(domain.LeadCottageAssociation).
		Preload(domain.LeadCottageStatusAssociation).
		Order("status_id ASC, id DESC").
		Find(&lead).Error
	if err != nil {
		return nil, 0, err
	}

	return lead, domain.Total(totalCount), nil
}

func (r *repository) RevokeLeadCottage(
	ctx context.Context,
	leadID domain.LeadID,
) error {
	err := r.db.
		WithContext(ctx).
		Table(domain.LeadCottageTableName).
		Where("id = ?", leadID).
		Update("status_id", domain.StatusInactive).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) IsLeadExistByDate(
	ctx context.Context,
	cottageID int64,
	issuedAt int64,
	expiresAt int64,
) (bool, error) {
	var isExist bool
	err := r.db.
		WithContext(ctx).
		Table(domain.LeadCottageTableName).
		Select("count(id) > 0").
		Where("cottage_id = ? AND ? < expires_at AND  ? > issued_at AND status_id = 1", cottageID, issuedAt, expiresAt).
		Find(&isExist).Error
	if err != nil {
		return false, err
	}

	return isExist, nil
}

func (r *repository) IsOtherLeadExist(
	ctx context.Context,
	leadID domain.LeadID,
	cottageID int64,
	issuedAt int64,
	expiresAt int64,
) (bool, error) {
	var isExist bool
	err := r.db.
		WithContext(ctx).
		Table(domain.LeadCottageTableName).
		Select("count(id) > 0").
		Where("cottage_id = ? AND ? < expires_at AND ? > issued_at AND status_id = 1 AND id != ?", cottageID, issuedAt, expiresAt, leadID).
		Find(&isExist).Error
	if err != nil {
		return false, err
	}

	return isExist, nil
}
