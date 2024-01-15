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
func NewRepository(db *gorm.DB) domain.PermissionsRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreatePermission(ctx context.Context, permission *domain.Permission) (domain.PermissionID, error) {
	err := r.db.WithContext(ctx).Table(domain.PermissionsTable).Create(&permission).Error
	if err != nil {
		return 0, err
	}

	return permission.ID, nil
}

func (r *repository) ListPermissionsByRole(ctx context.Context, roleID int64) ([]*domain.Permission, error) {
	var (
		db            = r.db
		permissions   []*domain.Permission
		permissionIDs int64
	)

	err := db.WithContext(ctx).Table(domain.RolesToPermissionsTable).
		Where("role_id = ?", roleID).
		Find(&permissionIDs).Error
	if err != nil {
		return nil, err
	}

	err = r.db.WithContext(ctx).Table(domain.PermissionsTable).
		Where("permission_id IN ?", permissionIDs).
		Find(&permissions).Error
	if err != nil {
		return nil, err
	}

	return permissions, nil
}

func (r *repository) GetPermissionByEndpoint(ctx context.Context, endpointName string) (*domain.Permission, error) {
	var permission *domain.Permission
	err := r.db.WithContext(ctx).Table(domain.PermissionsTable).
		Where("endpoint_name = ? AND is_active = true", endpointName).
		First(&permission).Error
	if err != nil {
		return nil, err
	}

	return permission, nil
}

func (r *repository) IsPermissionAllowed(ctx context.Context, roleID domain.RoleID, permissionID domain.PermissionID) (bool, error) {
	var isExist bool
	err := r.db.WithContext(ctx).Table(domain.RolesToPermissionsTable).
		Where("role_id = ? AND permission_id = ? AND is_active = true", roleID, permissionID).
		Select("count(permission_id) > 0 ").Find(&isExist).Error
	if err != nil {
		return false, err
	}

	return isExist, nil
}
