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

// NewRepository - creates a new repository
func NewRepository(db *gorm.DB) domain.BuilderRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(
	ctx context.Context,
	builder *domain.Builder,
) (domain.BuilderID, error) {
	err := r.db.
		WithContext(ctx).
		Table(domain.UsersTableName).
		Omit("IsFavourite").
		Create(&builder).Error
	if err != nil {
		return 0, err
	}

	return builder.ID, nil
}

func (r *repository) Update(
	ctx context.Context,
	builderID domain.BuilderID,
	builder *domain.Builder,
) error {
	err := r.db.WithContext(ctx).
		Table(domain.UsersTableName).
		Where("id = ?", builderID).
		Omit("IsFavourite").
		Updates(&builder).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Get(
	ctx context.Context,
	builderID domain.BuilderID,
) (*domain.Builder, error) {
	var builder domain.Builder
	err := r.db.WithContext(ctx).
		Table(domain.UsersTableName).
		Where("id = ?", builderID).
		First(&builder).Error
	if err != nil {
		return nil, err
	}

	return &builder, nil
}

func (r *repository) List(
	ctx context.Context,
	criteria domain.BuilderSearchCriteria,
) ([]*domain.Builder, domain.Total, error) {
	var (
		db         = r.db
		builders   []*domain.Builder
		totalCount int64
	)

	if criteria.ID > 0 {
		db = db.Where("id = ?", criteria.ID)
	}
	if len(criteria.Name) > 0 {
		db = db.Where(`UPPER(CONCAT (first_name, ' ', last_name)) LIKE UPPER(?)`, "%"+criteria.Name+"%")
	}
	if len(criteria.Email) > 0 {
		db = db.Where(`UPPER(email) LIKE UPPER(?)`, "%"+criteria.Email+"%")
	}
	if len(criteria.Phone) > 0 {
		db = db.Where(`UPPER(phone) LIKE UPPER(?)`, "%"+criteria.Phone+"%")
	}

	totalCount, err := r.getTotal(*db, domain.UsersTableName)
	if err != nil {
		return nil, 0, err
	}

	err = db.
		WithContext(ctx).
		Scopes(helpers.Paginate(criteria.Page)).
		Table(domain.UsersTableName).
		Where("role_id = ?", domain.RoleBuilder).
		Select("*, EXISTS(SELECT * FROM subscribers sub WHERE sub.builder_id = users.id and sub.subscriber_id = ?) as is_favourite", criteria.UserID).
		Order("id DESC").
		Find(&builders).Error
	if err != nil {
		return nil, 0, err
	}

	return builders, domain.Total(totalCount), nil
}

func (r *repository) Delete(
	ctx context.Context,
	builderID domain.BuilderID,
) error {
	user := new(domain.Builder)
	err := r.db.
		WithContext(ctx).
		Table(domain.UsersTableName).
		Omit("IsFavourite").
		Where("id = ?", builderID).
		Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) IsFavouriteBuilder(
	ctx context.Context,
	builderID domain.BuilderID,
	userID domain.UserID,
) (bool, error) {

	var isFavourite bool
	err := r.db.
		WithContext(ctx).
		Table(domain.UserBuildersTableName).
		Select("count(*) > 0").
		Where("builder_id = ? AND subscriber_id = ?", builderID, userID).
		Find(&isFavourite).Error
	if err != nil {
		return false, err
	}

	return isFavourite, nil
}

func (r *repository) getTotal(db gorm.DB, tableName string) (int64, error) {
	var totalCount int64
	err := db.
		Table(tableName).
		Where("role_id = ? AND deleted_at = 0", domain.RoleBuilder).
		Count(&totalCount).Error
	if err != nil {
		return 0, err
	}

	return totalCount, nil
}
