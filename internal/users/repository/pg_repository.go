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
func NewRepository(db *gorm.DB) domain.UserRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateUser(
	ctx context.Context,
	user *domain.User,
) (*domain.User, error) {
	err := r.db.
		WithContext(ctx).
		Table(domain.UsersTableName).
		Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) GetUser(
	ctx context.Context,
	userID domain.UserID,
) (*domain.User, error) {
	var user domain.User
	err := r.db.WithContext(ctx).
		Table(domain.UsersTableName).
		Where("id = ?", userID).
		First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) GetUserByEmailPassword(
	ctx context.Context,
	email string,
	password string,
) (*domain.User, error) {
	var user domain.User
	err := r.db.WithContext(ctx).
		Table(domain.UsersTableName).
		Where("email = ? AND password = ? AND is_banned = false", email, password).
		First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) UpdateUser(
	ctx context.Context,
	userID domain.UserID,
	user *domain.User,
) error {
	err := r.db.
		WithContext(ctx).
		Table(domain.UsersTableName).
		Where("id = ?", userID).
		Updates(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) MakeUserVerified(
	ctx context.Context,
	userID domain.UserID,
) error {
	err := r.db.
		WithContext(ctx).
		Table(domain.UsersTableName).
		Where("id = ?", userID).
		Update("is_verified", true).Error

	return err
}

func (r *repository) ListUsers(
	ctx context.Context,
	criteria domain.UserSearchCriteria,
) ([]*domain.User, domain.Total, error) {

	var (
		db         = r.db
		users      []*domain.User
		totalCount int64
	)

	if criteria.ID > 0 {
		db = db.Where("id = ?", criteria.ID)
	}
	if len(criteria.Name) > 0 {
		db = db.Where(`UPPER(CONCAT(first_name, ' ', last_name)) LIKE UPPER(?)`, "%"+criteria.Name+"%")
	}
	if len(criteria.Email) > 0 {
		db = db.Where(`UPPER(email) LIKE UPPER(?)`, "%"+criteria.Email+"%")
	}
	if len(criteria.Phone) > 0 {
		db = db.Where(`UPPER(phone) LIKE UPPER(?)`, "%"+criteria.Phone+"%")
	}

	err := db.
		Table(domain.UsersTableName).
		Where("role_id = ? AND deleted_at = 0", domain.RoleClient).
		Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.
		WithContext(ctx).
		Scopes(helpers.Paginate(criteria.Page)).
		Table(domain.UsersTableName).
		Where("role_id = ?", domain.RoleClient).
		Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, domain.Total(totalCount), nil
}

func (r *repository) DeleteUser(
	ctx context.Context,
	userID domain.UserID,
) error {
	user := new(domain.User)
	err := r.db.
		WithContext(ctx).
		Table(domain.UsersTableName).
		Where("id = ?", userID).
		Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) BanUser(
	ctx context.Context,
	userID domain.UserID,
) error {
	err := r.db.
		WithContext(ctx).
		Table(domain.UsersTableName).
		Where("id = ?", userID).
		Update("is_banned = ?", true).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) ResetPassword(
	ctx context.Context,
	userID domain.UserID,
	password string,
) error {
	err := r.db.
		WithContext(ctx).
		Table(domain.UsersTableName).
		Where("id = ?", userID).
		Update("password", password).Error
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

func (r *repository) IsPhoneNumberExist(
	ctx context.Context,
	phoneNumber string,
) (bool, error) {
	var isExist bool

	err := r.db.
		WithContext(ctx).
		Table(domain.UsersTableName).
		Select("count(id) > 0").
		Where("phone = ?", phoneNumber).
		Find(&isExist).Error
	if err != nil {
		return false, err
	}

	return isExist, nil
}

func (r *repository) GetUserByPhoneNumber(
	ctx context.Context,
	phone string,
) (*domain.User, error) {
	var user domain.User
	err := r.db.WithContext(ctx).
		Table(domain.UsersTableName).
		Where("phone = ?", phone).
		First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) GetUserByEmail(
	ctx context.Context,
	email string,
) (*domain.User, error) {
	var user domain.User
	err := r.db.WithContext(ctx).
		Table(domain.UsersTableName).
		Where("email = ?", email).
		First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) CreateResetPasswordToken(
	ctx context.Context,
	resetPasswordToken *domain.ResetPasswordToken,
) error {
	err := r.db.
		WithContext(ctx).
		Table(domain.ResetPasswordTokenTableName).
		Create(&resetPasswordToken).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetResetPasswordToken(
	ctx context.Context,
	token string,
) (*domain.ResetPasswordToken, error) {
	var resetToken domain.ResetPasswordToken
	err := r.db.
		WithContext(ctx).
		Table(domain.ResetPasswordTokenTableName).
		Where("token = ?", token).
		First(&resetToken).Error
	if err != nil {
		return nil, err
	}

	return &resetToken, nil
}

func (r *repository) UpdateResetPasswordToken(
	ctx context.Context,
	resetPasswordToken *domain.ResetPasswordToken,
) error {
	err := r.db.
		WithContext(ctx).
		Table(domain.ResetPasswordTokenTableName).
		Where("user_id = ? AND token = ?", resetPasswordToken.UserID, resetPasswordToken.Token).
		Updates(&resetPasswordToken).Error
	if err != nil {
		return err
	}

	return nil
}
