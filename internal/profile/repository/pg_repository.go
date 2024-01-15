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
func NewRepository(db *gorm.DB) domain.ProfileRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetProfile(
	ctx context.Context,
	userID domain.UserID,
) (*domain.Profile, error) {
	var user *domain.Profile
	err := r.db.
		WithContext(ctx).
		Table(domain.UsersTableName).
		Take(&user, userID).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) UpdateProfile(
	ctx context.Context,
	userID domain.UserID,
	profile *domain.Profile,
) error {
	err := r.db.WithContext(ctx).
		Table(domain.UsersTableName).
		Where("id = ?", userID).
		Updates(profile).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) IsEmailExist(
	ctx context.Context,
	email string,
) (bool, error) {
	var isExist bool

	err := r.db.
		WithContext(ctx).
		Table(domain.UsersTableName).
		Select("count(id) > 0").
		Where("email = ?", email).
		Find(&isExist).Error
	if err != nil {
		return false, err
	}

	return isExist, nil
}

func (r *repository) IsPhoneExist(
	ctx context.Context,
	phone string,
) (bool, error) {
	var isExist bool

	err := r.db.
		WithContext(ctx).
		Table(domain.UsersTableName).
		Select("count(id) > 0").
		Where("phone = ?", phone).
		Find(&isExist).Error
	if err != nil {
		return false, err
	}

	return isExist, nil
}
