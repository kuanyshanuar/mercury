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

// NewRepository _ creates a new repository
func NewRepository(db *gorm.DB) domain.LeadBuilderRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateLeadBuilder(
	ctx context.Context,
	lead *domain.LeadBuilder,
) error {
	err := r.db.
		WithContext(ctx).
		Table(domain.LeadBuildersTableName).
		Create(&lead).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetLeadBuilder(
	ctx context.Context,
	leadID domain.LeadID,
) (*domain.LeadBuilder, error) {
	var lead domain.LeadBuilder
	err := r.db.
		WithContext(ctx).
		Table(domain.LeadBuildersTableName).
		Take(&lead, leadID).Error
	if err != nil {
		return nil, err
	}

	return &lead, nil
}

func (r *repository) UpdateLeadBuilder(
	ctx context.Context,
	leadID domain.LeadID,
	lead *domain.LeadBuilder,
) error {
	err := r.db.
		WithContext(ctx).
		Table(domain.LeadBuildersTableName).
		Where("id = ?", leadID).
		Updates(&lead).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) ListLeadBuilders(
	ctx context.Context,
	criteria domain.LeadBuilderSearchCriteria,
) ([]*domain.LeadBuilder, domain.Total, error) {
	var (
		db         = r.db
		leads      []*domain.LeadBuilder
		totalCount int64
	)

	if criteria.StatusID > 0 {
		db = db.Where("builder_leads.status_id = ?", criteria.StatusID)
	}
	if len(criteria.Name) > 0 {
		db = db.Joins(`JOIN users AS r ON r.id = builder_leads.builder_id AND UPPER(r.first_name) LIKE UPPER(?)`, "%"+criteria.Name+"%")
	}

	countResult := db.Table(domain.LeadBuildersTableName).Where("builder_leads.deleted_at = 0").Count(&totalCount)
	if countResult.Error != nil {
		return nil, 0, countResult.Error
	}

	err := db.
		WithContext(ctx).
		Scopes(helpers.Paginate(criteria.Page)).
		Table(domain.LeadBuildersTableName).
		Preload(domain.LeadBuildersAssociation). // @TODO
		Preload(domain.LeadStatusAssociation).
		Order("status_id ASC, id DESC").
		Find(&leads).Error
	if err != nil {
		return nil, 0, err
	}

	return leads, domain.Total(totalCount), nil
}

func (r *repository) DeleteLeadBuilder(
	ctx context.Context,
	leadID domain.LeadID,
) error {
	lead := new(domain.LeadBuilder)
	err := r.db.
		WithContext(ctx).
		Table(domain.LeadBuildersTableName).
		Where("id = ?", leadID).
		Delete(&lead).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) RevokeLeadBuilder(
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

func (r *repository) IsLeadExistByDateRange(
	ctx context.Context,
	builderID domain.BuilderID,
	issuedAt int64,
	expiresAt int64,
) (bool, error) {
	var isExist bool
	err := r.db.
		WithContext(ctx).
		Table(domain.LeadBuildersTableName).
		Select("count(id) > 0").
		Where("builder_id = ? AND ? <= expires_at AND  ? >= issued_at AND status_id = 1", builderID, issuedAt, expiresAt).
		Find(&isExist).Error
	if err != nil {
		return false, err
	}

	return isExist, nil
}

func (r *repository) IsOtherLeadExist(
	ctx context.Context,
	leadID domain.LeadID,
	builderID domain.BuilderID,
	issuedAt int64,
	expiresAt int64,
) (bool, error) {
	var isExist bool
	err := r.db.
		WithContext(ctx).
		Table(domain.LeadBuildersTableName).
		Select("count(id) > 0").
		Where("residence_id = ? AND ? < expires_at AND ? > issued_at AND status_id = 1 AND id != ?", builderID, issuedAt, expiresAt, leadID).
		Find(&isExist).Error
	if err != nil {
		return false, err
	}

	return isExist, nil
}
