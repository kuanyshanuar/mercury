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
func NewRepository(db *gorm.DB) domain.ManagerRepository {
	return &repository{db: db}
}

func (r *repository) Create(
	ctx context.Context,
	manager *domain.Manager,
) (domain.ManagerID, error) {
	err := r.db.
		WithContext(ctx).
		Table(domain.UsersTableName).
		Create(&manager).Error
	if err != nil {
		return 0, err
	}

	return manager.ID, err
}

func (r *repository) Update(
	ctx context.Context,
	managerID domain.ManagerID,
	manager *domain.Manager,
) error {
	err := r.db.WithContext(ctx).
		Table(domain.UsersTableName).
		Where("id = ?", managerID).
		Updates(&manager).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Get(
	ctx context.Context,
	managerID domain.ManagerID,
) (*domain.Manager, error) {
	var manager domain.Manager
	err := r.db.WithContext(ctx).
		Table(domain.UsersTableName).
		Where("id = ?", managerID).
		First(&manager).Error
	if err != nil {
		return nil, err
	}

	return &manager, nil
}

func (r *repository) List(
	ctx context.Context,
	criteria domain.ManagerSearchCriteria,
) ([]*domain.Manager, domain.Total, error) {
	var (
		db         = r.db
		managers   []*domain.Manager
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

	totalCount, err := r.getTotal(*db, domain.UsersTableName)
	if err != nil {
		return nil, 0, err
	}

	err = db.
		WithContext(ctx).
		Scopes(helpers.Paginate(criteria.Page)).
		Table(domain.UsersTableName).
		Where("role_id = ?", domain.RoleManager).
		Find(&managers).Error
	if err != nil {
		return nil, 0, err
	}

	return managers, domain.Total(totalCount), nil
}

func (r *repository) Delete(
	ctx context.Context,
	managerID domain.ManagerID,
) error {
	user := new(domain.Manager)
	err := r.db.
		WithContext(ctx).
		Table(domain.UsersTableName).
		Omit("IsFavourite").
		Where("id = ?", managerID).
		Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) getTotal(db gorm.DB, tableName string) (int64, error) {
	var totalCount int64
	err := db.
		Table(tableName).
		Where("role_id = ? AND deleted_at = 0", domain.RoleManager).
		Count(&totalCount).Error
	if err != nil {
		return 0, err
	}

	return totalCount, nil
}
