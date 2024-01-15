package repository

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repository struct {
	db *gorm.DB
}

// NewRepository - creates UserCottageRepository
func NewRepository(db *gorm.DB) domain.UserCottageRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) AddFavouriteCottage(
	ctx context.Context,
	userID int64,
	cottageID int64,
) error {
	err := r.db.WithContext(ctx).Table(domain.FavouriteCottageTable).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "cottage_id"}},
		DoNothing: true,
	}).Create(&domain.UserCottage{
		UserID:    userID,
		CottageID: cottageID,
	}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) IsCottageExists(
	ctx context.Context,
	userID int64,
	cottageID int64,
) (bool, error) {
	var total int64
	err := r.db.WithContext(ctx).
		Table(domain.UserCottageTableName).
		Select("WHERE user_id = ? AND cottage_id = ?", userID, cottageID).
		Count(&total).Error
	if err != nil {
		return false, err
	}
	if total > 0 {
		return true, nil
	}
	return false, nil

}

func (r *repository) DeleteFavouriteCottage(
	ctx context.Context,
	userID int64,
	cottageID int64,
) error {
	var userCottage domain.UserCottage
	err := r.db.WithContext(ctx).
		Table(domain.UserCottageTableName).
		Where("user_id = ? AND cottage_id = ?", userID, cottageID).
		Delete(&userCottage).
		Error
	return err
}

func (r *repository) ListFavouriteCottages(
	ctx context.Context,
	userID int64,
	criteria domain.FavouriteCottagesSearchCriteria,
) ([]int64, domain.Total, error) {
	var (
		ids        []int64
		totalCount int64
		db         = r.db
	)

	order := helpers.PrepareOrder(criteria.Sorts)
	if len(order) == 0 {
		order = `"created_at" DESC`
	}
	db = db.Order(order)

	db = db.
		WithContext(ctx).
		Select("cottage_id").
		Table(domain.FavouriteCottageTable).
		Where("user_id = ?", userID)

	totalCount, err := r.getTotal(*db, domain.FavouriteCottageTable)
	if err != nil {
		return nil, 0, err
	}

	err = db.
		Scopes(helpers.Paginate(criteria.Page)).
		Find(&ids).Error
	return ids, domain.Total(totalCount), err
}

func (r *repository) getTotal(db gorm.DB, tableName string) (int64, error) {
	var totalCount int64
	err := db.
		Table(tableName).
		Count(&totalCount).Error
	if err != nil {
		return 0, err
	}

	return totalCount, nil
}
