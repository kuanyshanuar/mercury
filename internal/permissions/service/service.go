package service

import (
	"context"
	"fmt"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	errors "gitlab.com/zharzhanov/mercury/internal/error"

	"github.com/go-kit/log"
)

type service struct {
	repository domain.PermissionsRepository
}

func (s *service) CreatePermission(ctx context.Context, permission *domain.Permission, callerID domain.CallerID) (domain.PermissionID, error) {

	// Validate input
	if err := s.validatePermissionInternal(ctx, permission); err != nil {
		return 0, err
	}

	return s.repository.CreatePermission(ctx, permission)
}

// NewService - creates a new service
func NewService(
	repository domain.PermissionsRepository,
	logger log.Logger,
) domain.PermissionsService {
	var service domain.PermissionsService
	{
		service = newBasicService(repository)
		service = loggingServiceMiddleware(logger)(service)
	}
	return service
}

func newBasicService(
	repository domain.PermissionsRepository,
) domain.PermissionsService {
	return &service{
		repository: repository,
	}
}

func (s *service) List(
	ctx context.Context,
	callerID domain.CallerID,
) ([]*domain.Permission, error) {
	return nil, nil
}

func (s *service) Allow(
	ctx context.Context,
	permissionKey string,
	userID domain.UserID,
	roleID domain.RoleID,
	_ domain.CallerID,
) (bool, error) {

	// Validate if user exists and not disabled
	if userID <= 0 {
		return false, errors.NewErrInvalidArgument("user id required")
	}
	if roleID <= 0 {
		return false, errors.NewErrInvalidArgument("role id required")
	}

	// Get permission by endpoint
	permission, err := s.repository.GetPermissionByEndpoint(ctx, permissionKey)
	if err != nil {
		return false, err
	}

	if permission == nil {
		return false, errors.NewErrNotFound(
			fmt.Sprintf("permission %s not found", permissionKey),
		)
	}

	// Check if permission exists by role id and permission id
	isExist, err := s.repository.IsPermissionAllowed(ctx, domain.RoleID(roleID), permission.ID)
	if err != nil {
		return false, err
	}

	return isExist, nil
}

func (s *service) validatePermissionInternal(ctx context.Context, permission *domain.Permission) error {

	if len(permission.EndpointName) == 0 {
		return errors.NewErrInvalidArgument("endpoint name required")
	}
	if len(permission.Action) == 0 {
		return errors.NewErrInvalidArgument("action required")
	}
	if len(permission.CrudType) == 0 {
		return errors.NewErrInvalidArgument("crud type required")
	}

	return nil
}
