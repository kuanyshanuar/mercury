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

// NewRepository - creates a new repository.
func NewRepository(db *gorm.DB) domain.LeadResidenceRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateLeadResidence(
	ctx context.Context,
	lead *domain.LeadResidence,
) error {
	err := r.db.
		WithContext(ctx).
		Table(domain.LeadResidenceTableName).
		Create(&lead).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) ListLeadResidences(
	ctx context.Context,
	criteria domain.LeadResidenceSearchCriteria,
) ([]*domain.LeadResidence, domain.Total, error) {
	var (
		db         = r.db
		leads      []*domain.LeadResidence
		totalCount int64
	)

	if criteria.StatusID > 0 {
		db = db.Where("residence_leads.status_id = ?", criteria.StatusID)
	}
	if len(criteria.Name) > 0 {
		db = db.Joins(`JOIN residences AS r ON r.id = residence_leads.residence_id AND UPPER(r.title) LIKE UPPER(?)`, "%"+criteria.Name+"%")
	}

	countResult := db.Table(domain.LeadResidenceTableName).Where("residence_leads.deleted_at = 0").Count(&totalCount)
	if countResult.Error != nil {
		return nil, 0, countResult.Error
	}

	err := db.
		WithContext(ctx).
		Scopes(helpers.Paginate(criteria.Page)).
		Table(domain.LeadResidenceTableName).
		Preload(domain.LeadResidenceAssociation).
		Preload(domain.LeadStatusAssociation).
		Order("status_id ASC, id DESC").
		Find(&leads).Error
	if err != nil {
		return nil, 0, err
	}

	return leads, domain.Total(totalCount), nil
}

func (r *repository) UpdateLeadResidence(
	ctx context.Context,
	lead *domain.LeadResidence,
) error {
	err := r.db.
		WithContext(ctx).
		Table(domain.LeadResidenceTableName).
		Where("id = ?", lead.ID).
		Updates(lead).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteLeadResidence(
	ctx context.Context,
	leadID domain.LeadID,
) error {
	leadResidence := new(domain.LeadResidence)
	err := r.db.
		WithContext(ctx).
		Table(domain.LeadResidenceTableName).
		Where("id = ?", leadID).
		Delete(&leadResidence).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) RevokeLeadResidence(
	ctx context.Context,
	leadID domain.LeadID,
) error {
	err := r.db.
		WithContext(ctx).
		Table(domain.LeadResidenceTableName).
		Where("id = ?", leadID).
		Update("status_id", domain.StatusInactive).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetLeadResidence(
	ctx context.Context,
	leadID domain.LeadID,
) (*domain.LeadResidence, error) {
	var leadResidence *domain.LeadResidence

	err := r.db.
		Table(domain.LeadResidenceTableName).
		Take(&leadResidence, leadID).Error
	if err != nil {
		return nil, err
	}

	return leadResidence, nil
}

func (r *repository) IsLeadExistByDate(
	ctx context.Context,
	residenceID domain.ResidenceID,
	issuedAt int64,
	expiresAt int64,
) (bool, error) {
	var isExist bool
	err := r.db.
		WithContext(ctx).
		Table(domain.LeadResidenceTableName).
		Select("count(id) > 0").
		Where("residence_id = ? AND ? < expires_at AND  ? > issued_at AND status_id = 1", residenceID, issuedAt, expiresAt).
		Find(&isExist).Error
	if err != nil {
		return false, err
	}

	return isExist, nil
}

func (r *repository) IsOtherLeadExist(
	ctx context.Context,
	leadID domain.LeadID,
	residenceID domain.ResidenceID,
	issuedAt int64,
	expiresAt int64,
) (bool, error) {
	var isExist bool
	err := r.db.
		WithContext(ctx).
		Table(domain.LeadResidenceTableName).
		Select("count(id) > 0").
		Where("residence_id = ? AND ? < expires_at AND ? > issued_at AND status_id = 1 AND id != ?", residenceID, issuedAt, expiresAt, leadID).
		Find(&isExist).Error
	if err != nil {
		return false, err
	}

	return isExist, nil
}
