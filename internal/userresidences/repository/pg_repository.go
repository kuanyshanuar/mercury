package repository

import (
	"context"
	"gitlab.com/zharzhanov/mercury/internal/helpers"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repository struct {
	db *gorm.DB
}

// NewRepository - creates a new repository
func NewRepository(db *gorm.DB) domain.UserResidenceRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) AddResidenceToFavourites(
	ctx context.Context,
	userID domain.UserID,
	residenceID domain.ResidenceID,
) error {
	err := r.db.WithContext(ctx).Table(domain.FavouriteResidenceTable).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "residence_id"}},
		DoNothing: true,
	}).Create(&domain.UserResidence{
		UserID:      int64(userID),
		ResidenceID: int64(residenceID),
	}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) IsResidenceExist(ctx context.Context, userID domain.UserID, residenceID domain.ResidenceID) (bool, error) {
	return false, nil
}

func (r *repository) DeleteResidenceFromFavourites(ctx context.Context, userID domain.UserID, residenceID domain.ResidenceID) error {
	err := r.db.WithContext(ctx).Table(domain.FavouriteResidenceTable).
		Where("user_id = ? AND residence_id = ?", userID, residenceID).
		Delete(&domain.UserResidence{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) ListFavouriteResidences(
	ctx context.Context,
	userID domain.UserID,
	criteria domain.FavouriteResidencesSearchCriteria,
) ([]int64, domain.Total, error) {
	var (
		db         = r.db
		ids        []int64
		totalCount int64
	)

	order := helpers.PrepareOrder(criteria.Sorts)
	if len(order) == 0 {
		order = `"created_at" DESC`
	}
	db = db.Order(order)

	db = db.
		WithContext(ctx).
		Select("residence_id").
		Table(domain.FavouriteResidenceTable).
		Where("user_id = ?", userID)

	totalCount, err := r.getTotal(*db, domain.FavouriteResidenceTable)
	if err != nil {
		return nil, 0, err
	}

	err = db.
		Scopes(helpers.Paginate(criteria.Page)).
		Find(&ids).Error
	if err != nil {
		return nil, 0, err
	}

	return ids, domain.Total(totalCount), nil
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
