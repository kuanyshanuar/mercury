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
func NewRepository(db *gorm.DB) domain.UserBuilderRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Subscribe(
	ctx context.Context,
	subscriberID int64,
	builderID int64,
) (err error) {

	err = r.db.
		WithContext(ctx).
		Table(domain.UserBuildersTableName).
		Create(&domain.UserBuilder{
			BuilderID:    builderID,
			SubscriberID: subscriberID,
		}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Unsubscribe(
	ctx context.Context,
	subscriberID int64,
	builderID int64,
) (err error) {
	userBuilder := new(domain.UserBuilder)
	err = r.db.
		WithContext(ctx).
		Table(domain.UserBuildersTableName).
		Where("subscriber_id = ? AND builder_id = ?", subscriberID, builderID).
		Delete(&userBuilder).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) ListBuilders(
	ctx context.Context,
	subscriberID int64,
	criteria domain.UserBuilderSearchCriteria,
) ([]*domain.Builder, domain.Total, error) {

	var (
		db         = r.db
		builderIDs []int64
		builders   []*domain.Builder
		totalCount int64
	)

	db = db.
		WithContext(ctx).
		Select("builder_id").
		Table(domain.UserBuildersTableName).
		Where("subscriber_id = ?", subscriberID)

	err := db.
		Table(domain.UserBuildersTableName).
		Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.
		Scopes(helpers.Paginate(criteria.Page)).
		Find(&builderIDs).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.
		WithContext(ctx).
		Table(domain.UsersTableName).
		Select("*, EXISTS(SELECT * FROM subscribers sub WHERE sub.builder_id = users.id and sub.subscriber_id = ?) as is_favourite", subscriberID).
		Where("id IN (?)", builderIDs).
		Find(&builders).Error
	if err != nil {
		return nil, 0, err
	}

	return builders, domain.Total(totalCount), nil
}
